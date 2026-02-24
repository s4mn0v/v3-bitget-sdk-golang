# Bitget Golang SDK (Modified Version)

Bitget Golang SDK, fixed some errors from the original SDK and migrated to v2 of Bitget API.

For the original version, please visit the link below:
[Original version](https://github.com/BitgetLimited/v3-bitget-api-sdk/blob/master/bitget-golang-sdk-api/README_EN.md)

---

# Modifications Made

- **Enhanced Trading Operations (v2)**
  - Added support for **Modify Order** (Price/Size) for active limit orders.
  - Added **Flash Close Position** to exit positions at market price immediately.
  - Added **Reversal (Click Backhand)** functionality to flip positions.
  - Added specific endpoints for **TP/SL Plan Orders** (Place and Modify).
  > Ref: [Modify Order](https://www.bitget.com/api-doc/contract/trade/Modify-Order), [Flash Close](https://www.bitget.com/api-doc/contract/trade/Flash-Close-Position)
  > Modified Code: `pkg/client/v2/mixorderclient.go`

- **Advanced Market Data (v2)**
  - Added **Current & History Funding Rate** queries.
  - Added **Open Interest** and **Mark/Index/Market Prices** real-time data.
  - Added **Position Tier** (Leverage brackets) query.
  - Added **Historical Candles** for Index and Mark prices.
  > Ref: [Market Tickers](https://www.bitget.com/api-doc/contract/market/Get-All-Symbol-Ticker), [Funding Rate](https://www.bitget.com/api-doc/contract/market/Get-Current-Funding-Rate)
  > Modified Code: `pkg/client/v2/mixmarketclient.go`

- **Account & Sub-Account Management (v2)**
  - Added **Estimated Open Count** to calculate max affordable size before ordering.
  - Added **Account Bills** with support for the last 90 days.
  - Added **Sub-account Assets** query for master accounts.
  > Ref: [Est. Open Count](https://www.bitget.com/api-doc/contract/account/Est-Open-Count), [Sub-Account Assets](https://www.bitget.com/api-doc/contract/account/Get-Sub-Account-Contract-Assets)
  > Modified Code: `pkg/client/v2/mixaccountclient.go`

- **FUTURES History & Fills**
  - Added new model for listing FILLED historical transaction data.
  - Added func for get HISTORY POSITION of all symbols.
  > Ref: [Get Fill History](https://www.bitget.com/api-doc/contract/trade/Get-Fill-History), [Get History Position](https://www.bitget.com/api-doc/contract/position/Get-History-Position)
  > Modified Code: `internal/model/history.go`, `pkg/client/v2/mixorderclient.go`, `pkg/client/v2/mixaccountclient.go`

- **Spot Trading, Wallet, and WebSocket Extensions (v2)**
  - Added multiple new **Spot Account endpoints**:
    - Sub-account assets, deduct info, upgrade status, and account upgrade.
    - Modified Code: `pkg/client/v2/spotaccountclient.go`

  - Added new **Spot Market endpoints**:
    - VIP fee rates, merged depth, auction info, historical candles, and trade history.
    - Modified Code: `pkg/client/v2/spotmarketclient.go`

  - Added extended **Spot Trading functionality**:
    - Cancel & Replace Order
    - Batch Cancel & Replace Orders
    - Cancel all orders by symbol
    - Order info lookup
    - Plan order modification and batch cancel support
    - Modified Code: `pkg/client/v2/spotorderclient.go`

  - Added new **Spot Wallet and Sub-account operations**:
    - Modify deposit account
    - Sub-account transfers and deposit address queries
    - Transfer coin info
    - Cancel withdrawal
    - Sub-account deposit and transfer records
    - BGB deduct switch support
    - Modified Code: `pkg/client/v2/spotwalletclient.go`

  - Added **WebSocket trade operation support**
    - Implemented `SendTrade` method for sending trade operations via WebSocket.
    - Enables structured trade execution messaging.
    - Modified Code: `pkg/client/ws/bitgetwsclient.go`

  - Minor cleanup and structural improvements:
    - Removed redundant comments
    - Improved function grouping and readability
    - Modified Code: `pkg/client/v2/mixorderclient.go`

- **Public, Tax, Convert, Sub-Account & Market Intelligence Extensions (v2)**

  - Public Information (Common Client)
    - Added global **Server Time** endpoint.
    - Added **Announcements** query.
    - Added **Trade Rate (symbol-specific)** query.
    - Added **All Trade Rates (business line)** query.
    - New file: `pkg/client/v2/commonclient.go`

  - Tax Records (Tax Client)
    - Added **Spot transaction records** export.
    - Added **Futures transaction records** export.
    - Added **Margin transaction records** export.
    - Added **P2P transaction records** export.
    - New file: `pkg/client/v2/taxclient.go`

  - Convert & BGB Conversion (Convert Client)
    - Added full **Flash Convert (Swap)** suite:
      - Supported currencies
      - Quoted price
      - Execute convert trade
      - Convert history records
    - Added **BGB Dust Conversion**:
      - Supported coins list
      - Execute BGB convert
      - BGB convert records
    - New file: `pkg/client/v2/convertclient.go`

  - Virtual Sub-Account & API Key Management (User Client)
    - Added **Create virtual sub-account**
    - Added **Modify virtual sub-account**
    - Added **Virtual sub-account list**
    - Added **Create sub-account API key**
    - Added **Modify sub-account API key**
    - Added **Sub-account API key list**
    - New file: `pkg/client/v2/userclient.go`

  - Market Intelligence / Big Data Extensions
    - Added **Taker Buy/Sell Volume**
    - Added **Long/Short Ratio**
    - Added **Market Fund Flow**
    - Added **Historical Position Long/Short Ratio**
    - Extended file: `pkg/client/v2/mixmarketclient.go`
  
  Based on the updated Bitget documentation (Public, Tax, Sub-accounts, Convert, and extended Market Data), a significant number of missing endpoints were implemented and grouped into new client files to preserve the SDK’s existing architecture.

---

# Branch Structure and Module Usage

This repository provides two branches to support different import and module use cases:

- **main branch**
  - Module name: `github.com/s4mn0v/v3-bitget-sdk-golang`
  - Intended for external use as a Go module.
  - Recommended when importing the SDK in your own projects using Go modules.

- **fixed branch**
  - Module name: `bitget`
  - All internal imports use the local module name `bitget`.
  - Intended for local development, internal integration, or when embedding the SDK directly into your project without external module references.

Choose the branch depending on your integration needs. Use `main` for standard Go module usage, or `fixed` for local/internal SDK integration.

---

# Usage examples

### 1. Get Futures Fill History

```go
func (a *App) GetFuturesFillHistory(symbol string, productType string) (interface{}, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": productType,
  "limit":       "50",
 }

 resp, err := a.mixOrderClient.FillHistory(params)
 if err != nil {
  return nil, err
 }

 var result map[string]interface{}
 if err := json.Unmarshal([]byte(resp), &result); err != nil {
  return nil, err
 }

 if data, ok := result["data"].(map[string]interface{}); ok {
  if fillList, ok := data["fillList"]; ok {
   return fillList, nil
  }
@ }

 return []interface{}{}, nil
}

```

---

### 2. Flash Close a Position

```go
func (a *App) ClosePositionFast(symbol string, holdSide string) (string, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": "USDT-FUTURES",
  "holdSide":    holdSide, // "long" or "short"
 }

 return a.mixOrderClient.FlashClosePositions(params)
}
```

---

### 3. Get Estimated Openable Quantity

```go
func (a *App) GetMaxOpenSize(symbol string, price string, leverage string) (string, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": "USDT-FUTURES",
  "marginCoin":  "USDT",
  "openPrice":   price,
  "openAmount":  "1000", // Margin to use
  "leverage":    leverage,
 }

 return a.mixAccountClient.OpenCount(params)
}
```
