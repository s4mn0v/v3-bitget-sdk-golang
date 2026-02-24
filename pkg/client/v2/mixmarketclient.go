// Package v2
package v2

import (
	"github.com/s4mn0v/bitget/internal/common"
)

type MixMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixMarketClient) Init() *MixMarketClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *MixMarketClient) Contracts(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/contracts", params)
	return resp, err
}

func (p *MixMarketClient) Orderbook(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/orderbook", params)
	return resp, err
}

func (p *MixMarketClient) Ticker(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/ticker", params)
	return resp, err
}

func (p *MixMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/fills", params)
	return resp, err
}

func (p *MixMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/candles", params)
	return resp, err
}

// NEW FUNCTIONS

func (p *MixMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/tickers", params)
	return resp, err
}

func (p *MixMarketClient) CurrentFundRate(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/current-fund-rate", params)
	return resp, err
}

func (p *MixMarketClient) HistoryFundRate(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/history-fund-rate", params)
	return resp, err
}

func (p *MixMarketClient) OpenInterest(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/open-interest", params)
	return resp, err
}

func (p *MixMarketClient) SymbolPrice(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/symbol-price", params)
	return resp, err
}

func (p *MixMarketClient) PositionTier(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/mix/market/query-position-lever", params)
	return resp, err
}

func (p *MixMarketClient) TakerBuySellVolume(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/mix/market/taker-buy-sell", params)
}

func (p *MixMarketClient) LongShortRatio(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/mix/market/long-short", params)
}

func (p *MixMarketClient) MarketFundFlow(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/spot/market/fund-flow", params)
}

func (p *MixMarketClient) HistoryPositionLongShort(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/mix/market/position-long-short", params)
}
