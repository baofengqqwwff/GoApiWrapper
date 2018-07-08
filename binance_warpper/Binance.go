package binance_warpper

import (
	"encoding/json"
	"github.com/nntaoli-project/GoEx/binance"
	"net/http"
	"github.com/nntaoli-project/GoEx"
	. "github.com/baofengqqwwff/GoApiWrapper"
)

type BinanceWarpper struct {
	*binance.Binance
}

func New(client *http.Client, api_key, secret_key string) *BinanceWarpper {
	binanceWarpper := &BinanceWarpper{}
	binanceWarpper.Binance = binance.New(client, api_key, secret_key)
	return binanceWarpper
}

func (bn *BinanceWarpper) GetExchangeName() string {
	return bn.Binance.GetExchangeName()
}

func (bn *BinanceWarpper) GetTicker(currencyPair string) (*Ticker, error) {

	goexTicker, err := bn.Binance.GetTicker(goex.NewCurrencyPair2(currencyPair))
	if err != nil {
		return nil, err
	}
	goexjson, _ := json.Marshal(goexTicker)
	ticker := &Ticker{}
	err = json.Unmarshal(goexjson, ticker)
	if err != nil {
		return nil, err
	}
	return ticker, nil
}

func (bn *BinanceWarpper) GetDepth(size int, currencyPair string) (*Depth, error) {
	goexDepth, err := bn.Binance.GetDepth(size, goex.NewCurrencyPair2(currencyPair))
	if err != nil {
		return nil, err
	}
	goexjson, _ := json.Marshal(goexDepth)
	depth := &Depth{}
	err = json.Unmarshal(goexjson, depth)
	if err != nil {
		return nil, err
	}
	return depth, nil
}

func (bn *BinanceWarpper) GetAccount() (*Account, error) {
	goexAccount, err := bn.Binance.GetAccount()
	if err != nil {
		return nil, err
	}
	goexjson, _ := json.Marshal(goexAccount)
	account := &Account{}
	err = json.Unmarshal(goexjson, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

//
//func (bn *Binance) LimitBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
//	return bn.placeOrder(amount, price, currencyPair, "LIMIT", "BUY")
//}
//
//func (bn *Binance) LimitSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
//	return bn.placeOrder(amount, price, currencyPair, "LIMIT", "SELL")
//}
//
//func (bn *Binance) MarketBuy(amount, price string, currencyPair CurrencyPair) (*Order, error) {
//	return bn.placeOrder(amount, price, currencyPair, "MARKET", "BUY")
//}
//
//func (bn *Binance) MarketSell(amount, price string, currencyPair CurrencyPair) (*Order, error) {
//	return bn.placeOrder(amount, price, currencyPair, "MARKET", "SELL")
//}
//
//func (bn *Binance) CancelOrder(orderId string, currencyPair CurrencyPair) (bool, error) {
//	currencyPair = bn.adaptCurrencyPair(currencyPair)
//	path := API_V3 + ORDER_URI
//	params := url.Values{}
//	params.Set("symbol", currencyPair.ToSymbol(""))
//	params.Set("orderId", orderId)
//
//	bn.buildParamsSigned(&params)
//
//	resp, err := HttpDeleteForm(bn.httpClient, path, params, map[string]string{"X-MBX-APIKEY": bn.accessKey})
//
//	//log.Println("resp:", string(resp), "err:", err)
//	if err != nil {
//		return false, err
//	}
//
//	respmap := make(map[string]interface{})
//	err = json.Unmarshal(resp, &respmap)
//	if err != nil {
//		log.Println(string(resp))
//		return false, err
//	}
//
//	orderIdCanceled := ToInt(respmap["orderId"])
//	if orderIdCanceled <= 0 {
//		return false, errors.New(string(resp))
//	}
//
//	return true, nil
//}
//
//func (bn *Binance) GetOneOrder(orderId string, currencyPair CurrencyPair) (*Order, error) {
//	params := url.Values{}
//	currencyPair = bn.adaptCurrencyPair(currencyPair)
//	params.Set("symbol", currencyPair.ToSymbol(""))
//	if orderId != "" {
//		params.Set("orderId", orderId)
//	}
//	params.Set("orderId", orderId)
//
//	bn.buildParamsSigned(&params)
//	path := API_V3 + ORDER_URI + params.Encode()
//
//	respmap, err := HttpGet2(bn.httpClient, path, map[string]string{"X-MBX-APIKEY": bn.accessKey})
//	//log.Println(respmap)
//	if err != nil {
//		return nil, err
//	}
//	status := respmap["status"].(string)
//	side := respmap["side"].(string)
//
//	ord := Order{}
//	ord.Currency = currencyPair
//	ord.OrderID = ToInt(orderId)
//	ord.OrderID2 = orderId
//
//	if side == "SELL" {
//		ord.Side = SELL
//	} else {
//		ord.Side = BUY
//	}
//
//	switch status {
//	case "FILLED":
//		ord.Status = ORDER_FINISH
//	case "PARTIALLY_FILLED":
//		ord.Status = ORDER_PART_FINISH
//	case "CANCELED":
//		ord.Status = ORDER_CANCEL
//	case "PENDING_CANCEL":
//		ord.Status = ORDER_CANCEL_ING
//	case "REJECTED":
//		ord.Status = ORDER_REJECT
//	}
//
//	ord.Amount = ToFloat64(respmap["origQty"].(string))
//	ord.Price = ToFloat64(respmap["price"].(string))
//	ord.DealAmount = ToFloat64(respmap["executedQty"])
//	ord.AvgPrice = ord.Price // response no avg price ， fill price
//
//	return &ord, nil
//}
//
//func (bn *Binance) GetUnfinishOrders(currencyPair CurrencyPair) ([]Order, error) {
//	params := url.Values{}
//	currencyPair = bn.adaptCurrencyPair(currencyPair)
//	params.Set("symbol", currencyPair.ToSymbol(""))
//
//	bn.buildParamsSigned(&params)
//	path := API_V3 + UNFINISHED_ORDERS_INFO + params.Encode()
//
//	respmap, err := HttpGet3(bn.httpClient, path, map[string]string{"X-MBX-APIKEY": bn.accessKey})
//	//log.Println("respmap", respmap, "err", err)
//	if err != nil {
//		return nil, err
//	}
//
//	orders := make([]Order, 0)
//	for _, v := range respmap {
//		ord := v.(map[string]interface{})
//		side := ord["side"].(string)
//		orderSide := SELL
//		if side == "BUY" {
//			orderSide = BUY
//		}
//
//		orders = append(orders, Order{
//			OrderID:   ToInt(ord["orderId"]),
//			OrderID2:  fmt.Sprint(ToInt(ord["id"])),
//			Currency:  currencyPair,
//			Price:     ToFloat64(ord["price"]),
//			Amount:    ToFloat64(ord["origQty"]),
//			Side:      TradeSide(orderSide),
//			Status:    ORDER_UNFINISH,
//			OrderTime: ToInt(ord["time"])})
//	}
//	return orders, nil
//}
//
//func (bn *Binance) GetKlineRecords(currency CurrencyPair, period, size, since int) ([]Kline, error) {
//	panic("not implements")
//}
//
////非个人，整个交易所的交易记录
//func (bn *Binance) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
//	panic("not implements")
//}
//
//func (bn *Binance) GetOrderHistorys(currency CurrencyPair, currentPage, pageSize int) ([]Order, error) {
//	panic("not implements")
//}
//
//func (ba *Binance) adaptCurrencyPair(pair CurrencyPair) CurrencyPair {
//	return pair.AdaptBchToBcc().AdaptUsdToUsdt()
//}