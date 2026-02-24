package v2

import (
	"github.com/s4mn0v/bitget/internal/common"
)

type TaxClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *TaxClient) Init() *TaxClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *TaxClient) SpotRecord(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/tax/spot-record", params)
}

func (p *TaxClient) FutureRecord(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/tax/future-record", params)
}

func (p *TaxClient) MarginRecord(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/tax/margin-record", params)
}

func (p *TaxClient) P2pRecord(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/tax/p2p-record", params)
}
