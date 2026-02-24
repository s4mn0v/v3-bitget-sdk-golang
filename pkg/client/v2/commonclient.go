package v2

import (
	"github.com/s4mn0v/bitget/internal/common"
)

type CommonClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *CommonClient) Init() *CommonClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

// Get Server Time (Doc 14)

func (p *CommonClient) ServerTime() (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/public/time", nil)
}

// Query Announcements (Doc 12)

func (p *CommonClient) Announcements(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/public/annoucements", params)
}

// Get Trade Rate for specific symbol (Doc 15)

func (p *CommonClient) TradeRate(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/common/trade-rate", params)
}

// Get All Trade Rates for a business line (Doc 16)

func (p *CommonClient) AllTradeRates(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/common/all-trade-rate", params)
}
