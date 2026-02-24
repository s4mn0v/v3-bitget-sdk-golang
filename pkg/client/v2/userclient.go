package v2

import (
	"github.com/s4mn0v/bitget/internal"
	"github.com/s4mn0v/bitget/internal/common"
)

type UserClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *UserClient) Init() *UserClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *UserClient) CreateVirtualSubaccount(params map[string]interface{}) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/user/create-virtual-subaccount", body)
}

func (p *UserClient) ModifyVirtualSubaccount(params map[string]interface{}) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/user/modify-virtual-subaccount", body)
}

func (p *UserClient) VirtualSubaccountList(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/user/virtual-subaccount-list", params)
}

func (p *UserClient) CreateSubaccountAPIKey(params map[string]interface{}) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/user/create-virtual-subaccount-apikey", body)
}

func (p *UserClient) ModifySubaccountAPIKey(params map[string]interface{}) (string, error) {
	body, _ := internal.ToJson(params)
	return p.BitgetRestClient.DoPost("/api/v2/user/modify-virtual-subaccount-apikey", body)
}

func (p *UserClient) SubaccountAPIKeyList(params map[string]string) (string, error) {
	return p.BitgetRestClient.DoGet("/api/v2/user/virtual-subaccount-apikey-list", params)
}
