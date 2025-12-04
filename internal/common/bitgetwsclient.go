package common

import (
	"bitget/config"
	"bitget/constants"
	"bitget/internal"
	"bitget/internal/model"
	"bitget/logging/applogger"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron"
	"sync"
	"time"
)

type BitgetBaseWsClient struct {
	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Listener         OnReceive
	ErrorListener    OnReceive
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSuribe        *model.Set
	Signer           *Signer
	ScribeMap        map[model.SubscribeReq]OnReceive
	isClosed          bool
}

func (p *BitgetBaseWsClient) Close() {
	p.isClosed = true

	if p.WebSocketClient != nil {
		_ = p.WebSocketClient.Close()
	}

	if p.Ticker != nil {
		p.Ticker.Stop()
	}
}

func (p *BitgetBaseWsClient) Init() *BitgetBaseWsClient {
	p.Connection = false
	p.AllSuribe = model.NewSet()
	p.Signer = new(Signer).Init(config.SecretKey)
	p.ScribeMap = make(map[model.SubscribeReq]OnReceive)
	p.SendMutex = &sync.Mutex{}
	p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	p.LastReceivedTime = time.Now()

	return p
}

func (p *BitgetBaseWsClient) SetListener(msgListener OnReceive, errorListener OnReceive) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *BitgetBaseWsClient) Connect() {

	p.tickerLoop()
	p.ExecuterPing()
}

func (p *BitgetBaseWsClient) ConnectWebSocket() {
	var err error
	applogger.Info("WebSocket connecting...")
	p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(config.WsUrl, nil)
	if err != nil {
		fmt.Printf("WebSocket connected error: %s\n", err)
		return
	}
	applogger.Info("WebSocket connected")
	p.Connection = true
}

func (p *BitgetBaseWsClient) Login() {
	timesStamp := internal.TimesStampSec()
	sign := p.Signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	if constants.RSA == config.SignType {
		sign = p.Signer.SignByRSA(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)
	}

	loginReq := model.WsLoginReq{
		ApiKey:     config.ApiKey,
		Passphrase: config.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}
	var args []interface{}
	args = append(args, loginReq)

	baseReq := model.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: args,
	}
	p.SendByType(baseReq)
}

func (p *BitgetBaseWsClient) StartReadLoop() {
	go p.ReadLoop()
}

func (p *BitgetBaseWsClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc("*/15 * * * * *", p.ping)
	c.Start()
}
func (p *BitgetBaseWsClient) ping() {
	if p.isClosed {
		return
	}
	p.Send("ping")
}

func (p *BitgetBaseWsClient) SendByType(req model.WsBaseReq) {
	json, _ := internal.ToJson(req)
	p.Send(json)
}

func (p *BitgetBaseWsClient) Send(data string) {
	if p.WebSocketClient == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}
	applogger.Info("sendMessage:%s", data)
	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *BitgetBaseWsClient) tickerLoop() {
	applogger.Info("tickerLoop started")
	for range p.Ticker.C {
		elapsedSecond := time.Since(p.LastReceivedTime).Seconds()

		if elapsedSecond > constants.ReconnectWaitSecond {
			applogger.Info("WebSocket reconnect...")
			p.disconnectWebSocket()
			p.ConnectWebSocket()
		}
	}
}

func (p *BitgetBaseWsClient) disconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	fmt.Println("WebSocket disconnecting...")
	err := p.WebSocketClient.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

func (p *BitgetBaseWsClient) ReadLoop() {

	// Evita que un panic tumbe la app
	defer func() {
		if r := recover(); r != nil {
			applogger.Error("Recovered from WS panic: %v", r)
		}
	}()

	for {
		if p.isClosed {
			applogger.Info("ReadLoop stopped: client closed")
			return
		}

		if p.WebSocketClient == nil {
			applogger.Info("ReadLoop: WebSocketClient is nil, exiting read loop")
			return
		}

		// ---- INTENTA LEER UN MENSAJE ----
		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			applogger.Info("ReadLoop: WebSocket read error: %s", err)

			if p.isClosed {
				return // termina sin spamear errores
			}

			continue
		}

		p.LastReceivedTime = time.Now()
		message := string(buf)

		applogger.Info("rev:" + message)

		// ---- KEEP ALIVE ----
		if message == "pong" {
			applogger.Info("Keep connected: pong")
			continue
		}

		jsonMap := internal.JSONToMap(message)

		// ---- ERRORES DEL SERVIDOR ----
		if v, ok := jsonMap["code"]; ok {
			if int(v.(float64)) != 0 {
				p.ErrorListener(message)
				continue
			}
		}

		// ---- LOGIN WS ----
		if v, ok := jsonMap["event"]; ok && v == "login" {
			applogger.Info("login msg:" + message)
			p.LoginStatus = true
			continue
		}

		// ---- MENSAJE CON DATA ----
		if _, ok := jsonMap["data"]; ok {
			listener := p.GetListener(jsonMap["arg"])
			listener(message)
			continue
		}

		// ---- MENSAJE GENÉRICO ----
		p.handleMessage(message)
	}
}

func (p *BitgetBaseWsClient) GetListener(argJson interface{}) OnReceive {

	mapData := argJson.(map[string]interface{})

	subscribeReq := model.SubscribeReq{
		InstType: fmt.Sprintf("%v", mapData["instType"]),
		Channel:  fmt.Sprintf("%v", mapData["channel"]),
		InstId:   fmt.Sprintf("%v", mapData["instId"]),
	}

	v, e := p.ScribeMap[subscribeReq]

	if !e {
		return p.Listener
	}
	return v
}

type OnReceive func(message string)

func (p *BitgetBaseWsClient) handleMessage(msg string) {
	fmt.Println("default:" + msg)
}
