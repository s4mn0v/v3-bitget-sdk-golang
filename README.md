# Bitget Golang SDK (Modified Version)

Bitget Golang SDK, fixed some errors from the original SDK and migreated to v2 of Bitget API.


For the original version, please visit the link below:
[Original version](https://github.com/BitgetLimited/v3-bitget-api-sdk/blob/master/bitget-golang-sdk-api/README_EN.md)


--- 

# Modifications Made

- Adding new model for v2, listing historical transaction data of FUTURES operations  
  > Ref: [Get Fill History](https://www.bitget.com/api-doc/contract/trade/Get-Fill-History)  
  > Modified Code:  [new model](./internal/model/history.go)  [edited mixorderclient.go](./pkg/client/v2/mixorderclient.go)


Usage example in app.go:

```go
func (a *App) GetFuturesFillHistory(symbol string, productType string) (interface{}, error) {
	params := map[string]string{
		"symbol":      symbol,
		"productType": productType,
		"limit":       "50",
	}

	resp, err := a.mixClient.FillHistory(params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(resp), &result); err != nil {
		return nil, err
	}

	// Acceder a data -> fillList
	if data, ok := result["data"].(map[string]interface{}); ok {
		if fillList, ok := data["fillList"]; ok {
			return fillList, nil
		}
	}

	// Si no hay fills, devolver array vacío
	return []interface{}{}, nil
}
```