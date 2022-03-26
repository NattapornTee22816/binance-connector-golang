package spot

import (
	"github.com/NattapornTee22816/binance-connector-golang/model"
	"net/http"
)

// NewOrderTest (Test New Order (TRADE))
// Test new order creation and signature/recvWindow long.
// Creates and validates a new order but does not send it into the matching engine.
// POST /api/v3/order/test (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#test-new-order-trade
func (r *API) NewOrderTest(param *model.OrderParam) error {
	_, err := r.sendRequest(http.MethodPost, "/api/v3/order/test", param, model.EndpointSecurityTypeTrade)
	return err
}

// NewOrder (TRADE)
// Send in a new order.
// POST /api/v3/order (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#new-order-trade
func (r *API) NewOrder(param *model.OrderParam) (*model.Order, error) {
	bytes, err := r.sendRequest(http.MethodPost, "/api/v3/order", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseOrder(bytes)
}

// CancelOrder (TRADE)
// Cancel an active order.
// DELETE /api/v3/order (HMAC SHA256)
func (r *API) CancelOrder(param *model.CancelOrderParam) (*model.CancelOrder, error) {
	bytes, err := r.sendRequest(http.MethodDelete, "/api/v3/order", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseCancelOrder(bytes)
}

// CancelOpenOrder
// Cancel all Open Orders on a Symbol (TRADE)
// Cancels all active orders on a symbol.
// This includes OCO orders.
// DELETE /api/v3/openOrders
// https://binance-docs.github.io/apidocs/spot/en/#cancel-all-open-orders-on-a-symbol-trade
func (r *API) CancelOpenOrder(param *model.CancelOrderParam) ([]*model.CancelOpenOrder, error) {
	bytes, err := r.sendRequest(http.MethodDelete, "/api/v3/openOrders", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseCancelOpenOrder(bytes)
}

// GetOrder
// Query Order (USER_DATA)
// Check an order's status.
// GET /api/v3/order (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#query-order-user_data
func (r *API) GetOrder(param *model.GetOrderParam) (*model.GetOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/order", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOrder(bytes)
}

// GetOpenOrders
// Current Open Orders (USER_DATA)
// Get all open orders on a symbol. Careful when accessing this with no symbol.
// GET /api/v3/openOrders (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#current-open-orders-user_data
func (r *API) GetOpenOrders(param *model.GetOpenOrdersParam) ([]*model.GetOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/openOrders", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOpenOrder(bytes)
}

// GetOrders
// All Orders (USER_DATA)
// Get all account orders; active, canceled, or filled.
// GET /api/v3/allOrders (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#all-orders-user_data
func (r *API) GetOrders(param *model.GetOrdersParam) ([]*model.GetOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/allOrders", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOrders(bytes)
}

// NewOcoOrder
// New OCO (TRADE)
// Send in a new OCO
// POST /api/v3/order/oco (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#new-oco-trade
func (r *API) NewOcoOrder(param *model.NewOcoOrderParam) (*model.OcoOrder, error) {
	bytes, err := r.sendRequest(http.MethodPost, "/api/v3/order/oco", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseNewOcoOrder(bytes)
}

// CancelOcoOrder
// Cancel OCO (TRADE)
// Cancel an entire Order List.
// DELETE /api/v3/orderList (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#cancel-oco-trade
func (r *API) CancelOcoOrder(param *model.CancelOcoOrderParam) (*model.CancelOcoOrder, error) {
	bytes, err := r.sendRequest(http.MethodDelete, "/api/v3/orderList", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseCancelOcoOrder(bytes)
}

// GetOcoOrder
// Query OCO (USER_DATA)
// Retrieves a specific OCO based on provided optional parameters
// GET /api/v3/orderList (HMAC SHA256)
func (r *API) GetOcoOrder(param *model.GetOcoOrderParam) (*model.GetOcoOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/orderList", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOcoOrder(bytes)
}

// GetOcoOrders
// Query all OCO (USER_DATA)
// Retrieves all OCO based on provided optional parameters
// GET /api/v3/allOrderList (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#query-all-oco-user_data
func (r *API) GetOcoOrders(param *model.GetOcoOrdersParam) ([]*model.GetOcoOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/allOrderList", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOcoOrders(bytes)
}

// GetOcoOpenOrders
// Query Open OCO (USER_DATA)
// GET /api/v3/openOrderList (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#query-open-oco-user_data
func (r *API) GetOcoOpenOrders(param *model.GetOcoOpenOrdersParam) ([]*model.GetOcoOrder, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/openOrderList", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOcoOpenOrders(bytes)
}

// Account Information (USER_DATA)
// Get current account information.
// GET /api/v3/account (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#account-information-user_data
func (r *API) Account(param *model.AccountParam) (*model.Account, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/account", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseAccount(bytes)
}

// MyTrades
// Account Trade List (USER_DATA)
// Get trades for a specific account and symbol.
// GET /api/v3/myTrades (HMAC SHA256)
// https://binance-docs.github.io/apidocs/spot/en/#account-trade-list-user_data
func (r *API) MyTrades(param model.MyTradesParam) ([]*model.MyTrade, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/myTrades", param, model.EndpointSecurityTypeUserData)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseMyTrades(bytes)
}

// GetOrderRateLimit
// Query Current Order Count Usage (TRADE)
// Displays the user's current order count usage for all intervals.
// GET /api/v3/rateLimit/order
// https://binance-docs.github.io/apidocs/spot/en/#query-current-order-count-usage-trade
func (r *API) GetOrderRateLimit(param *model.GetOrderRateLimitParam) ([]*model.GetOrderRateLimit, error) {
	bytes, err := r.sendRequest(http.MethodGet, "/api/v3/rateLimit/order", param, model.EndpointSecurityTypeTrade)
	if err != nil {
		if r.logger.CanDebug() {
			r.logger.Error(err.Error())
		}
		return nil, err
	}

	return r.parser.ParseGetOrderRateLimit(bytes)
}
