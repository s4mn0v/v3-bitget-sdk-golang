package v2

import (
	"github.com/s4mn0v/bitget/internal"
	"github.com/s4mn0v/bitget/internal/common"
)

type ConvertClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *ConvertClient) Init() *ConvertClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

// Flash Convert Endpoints

func (p *ConvertClient) Currencies() (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/convert/currencies", nil)
}

func (p *ConvertClient) QuotedPrice(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/convert/quoted-price", params)
}

func (p *ConvertClient) Trade(params map[string]string) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/convert/trade", body)
}

func (p *ConvertClient) Record(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/convert/convert-record", params)
}

// BGB Convert Endpoints

func (p *ConvertClient) BgbConvertCoins() (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/convert/bgb-convert-coin-list", nil)
}

func (p *ConvertClient) BgbConvert(params map[string]interface{}) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/convert/bgb-convert", body)
}

func (p *ConvertClient) BgbConvertRecords(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/convert/bgb-convert-records", params)
}
