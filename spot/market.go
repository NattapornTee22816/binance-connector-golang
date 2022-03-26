package spot

import (
	"encoding/json"
	"github.com/NattapornTee22816/binance-connector-golang/model"
	"github.com/buger/jsonparser"
	"net/http"
)

// Ping
// Test connectivity to the Rest API.
// GET /api/v3/ping
// https://binance-docs.github.io/apidocs/spot/en/#test-connectivity
func (r *API) Ping() error {
	_, err := r.sendRequest(http.MethodGet, "/api/v3/ping", nil, model.EndpointSecurityTypeNone)
	return err
}

// CheckServerTime
// Test connectivity to the Rest API and get the current server time.
// GET /api/v3/time
// https://binance-docs.github.io/apidocs/spot/en/#check-server-time
func (r *API) CheckServerTime() (int64, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/time", nil, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return 0, err
	}
	return jsonparser.GetInt(bytes, "serverTime")
}

// ExchangeInformation
// Current exchange trading rules and symbol information
// GET /api/v3/exchangeInfo
// https://binance-docs.github.io/apidocs/spot/en/#exchange-information
func (r *API) ExchangeInformation(param *model.ExchangeInformationParam) (*model.ExchangeInformation, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/exchangeInfo", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	body := new(model.ExchangeInformation)
	if err := json.Unmarshal(bytes, &body); err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return body, nil
}

// OrderBook
// GET /api/v3/depth
// https://binance-docs.github.io/apidocs/spot/en/#order-book
func (r *API) OrderBook(param *model.OrderBookParam) (*model.OrderBook, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/depth", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseOrderBook(bytes)
}

// RecentTradesList
// Get recent trades.
// GET /api/v3/trades
// https://binance-docs.github.io/apidocs/spot/en/#recent-trades-list
func (r *API) RecentTradesList(param *model.RecentTradeParam) ([]*model.RecentTrade, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/trades", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Debug(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseRecentTrade(bytes)
}

// HistoricalTrades (OldTradeLookup) (MARKET_DATA)
// Get older market trades.
// GET /api/v3/historicalTrades
// https://binance-docs.github.io/apidocs/spot/en/#old-trade-lookup-market_data
func (r *API) HistoricalTrades(param *model.OldTradeLookupParam) ([]*model.OldTradeLookup, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/historicalTrades", param, model.EndpointSecurityTypeMarketData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseOldTradeLookup(bytes)
}

// AggTrades (Compressed/Aggregate Trades List)
// Get compressed, aggregate trades. Trades that fill at the time,
// from the same order, with the same price will have the quantity aggregated.
// GET /api/v3/aggTrades
// https://binance-docs.github.io/apidocs/spot/en/#compressed-aggregate-trades-list
func (r *API) AggTrades(param *model.AggregateTradeParam) ([]*model.AggregateTrade, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/aggTrades", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseAggregateTrade(bytes)
}

// Klines (Kline/Candlestick Data)
// Kline/candlestick bars for a symbol.
// Klines are uniquely identified by their open time.
// GET /api/v3/klines
// https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-data
func (r *API) Klines(param *model.KlineParam) ([]*model.Kline, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/klines", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseKline(bytes)
}

// AveragePrice (Current Average Price)
// Current average price for a symbol.
// GET /api/v3/avgPrice
// https://binance-docs.github.io/apidocs/spot/en/#current-average-price
func (r *API) AveragePrice(param *model.AveragePriceParam) (*model.AveragePrice, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/avgPrice", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	result := new(model.AveragePrice)
	if err := json.Unmarshal(bytes, &result); err != nil {
		if r.logger.CanDebug() {
			r.logger.Debug(err.Error())
		}
		return nil, err
	}

	return result, nil
}

// Ticker24hr (24hr Ticker Price Change Statistics)
// 24 hour rolling window price change statistics. Careful when accessing this with no symbol.
// GET /api/v3/ticker/24hr
// https://binance-docs.github.io/apidocs/spot/en/#24hr-ticker-price-change-statistics
func (r *API) Ticker24hr(param *model.Ticker24hrParam) ([]*model.Ticker24hr, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/ticker/24hr", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseTicker24hr(bytes)
}

// TickerPrice (Symbol Price Ticker)
// Latest price for a symbol or symbols.
// GET /api/v3/ticker/price
// https://binance-docs.github.io/apidocs/spot/en/#symbol-price-ticker
func (r *API) TickerPrice(param *model.TickerPriceParam) ([]*model.TickerPrice, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/ticker/price", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseTickerPrice(bytes)
}

// BookTicker (Symbol Order Book Ticker)
// Best price/qty on the order book for a symbol or symbols.
// GET /api/v3/ticker/bookTicker
// https://binance-docs.github.io/apidocs/spot/en/#symbol-order-book-ticker
func (r *API) BookTicker(param *model.BookTickerParam) ([]*model.BookTicker, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/ticker/bookTicker", param, model.EndpointSecurityTypeNone)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseBookTicker(bytes)
}
