package model

type FillHistoryResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    []Fill `json:"data"`
}

type Fill struct {
	Symbol      string `json:"symbol"`
	OrderId     string `json:"orderId"`
	FillId      string `json:"fillId"`
	Side        string `json:"side"`        // buy, sell
	OrderRole   string `json:"orderRole"`   // taker, maker
	Price       string `json:"price"`
	Size        string `json:"size"`
	Fee         string `json:"fee"`
	FeeByCoin   string `json:"feeByCoin"`
	TradeScope  string `json:"tradeScope"`
	Ctime       string `json:"cTime"`       // Create time
	Profit      string `json:"profit"`      // PNL
}