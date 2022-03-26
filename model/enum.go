package model

type EndpointSecurityType = string

var (
	EndpointSecurityTypeNone       = "NONE"
	EndpointSecurityTypeTrade      = "TRADE"
	EndpointSecurityTypeMargin     = "MARGIN"
	EndpointSecurityTypeUserData   = "USER_DATA"
	EndpointSecurityTypeUserStream = "USER_STREAM"
	EndpointSecurityTypeMarketData = "MARKET_DATA"
)

// Enum definitions
// https://binance-docs.github.io/apidocs/spot/en/#public-api-definitions

type SymbolStatus = string

var (
	SymbolStatusPreTrading   = SymbolStatus("PRE_TRADING")
	SymbolStatusTrading      = SymbolStatus("TRADING")
	SymbolStatusPostTrading  = SymbolStatus("POST_TRADING")
	SymbolStatusEndOfDay     = SymbolStatus("END_OF_DAY")
	SymbolStatusHalt         = SymbolStatus("HALT")
	SymbolStatusAuctionMatch = SymbolStatus("AUCTION_MATCH")
	SymbolStatusBreak        = SymbolStatus("BREAK")
)

type SymbolType = string

var (
	SymbolTypeSpot = SymbolType("SPOT")
)

type OrderStatus = string

var (
	OrderStatusNew             = OrderStatus("NEW")
	OrderStatusPartiallyFilled = OrderStatus("PARTIALLY_FILLED")
	OrderStatusFilled          = OrderStatus("FILLED")
	OrderStatusCanceled        = OrderStatus("CANCELED")
	OrderStatusPendingCancel   = OrderStatus("PENDING_CANCEL")
	OrderStatusRejected        = OrderStatus("REJECTED")
	OrderStatusExpired         = OrderStatus("EXPIRED")
)

type OCOStatus = string

var (
	OCOStatusResponse    = OCOStatus("RESPONSE")
	OCOStatusExecStarted = OCOStatus("EXEC_STARTED")
	OCOStatusAllDone     = OCOStatus("ALL_DONE")
)

type OCOOrderStatus = string

var (
	OCOOrderStatusExecuting = OCOOrderStatus("EXECUTING")
	OCOOrderStatusAllDone   = OCOOrderStatus("ALL_DONE")
	OCOOrderStatusReject    = OCOOrderStatus("REJECT")
)

type ContingencyType = string

var (
	ContingencyTypeOCO = ContingencyType("OCO")
)

type OrderType = string

var (
	OrderTypeLimit           = OrderType("LIMIT")
	OrderTypeMarket          = OrderType("MARKET")
	OrderTypeStopLoss        = OrderType("STOP_LOSS")
	OrderTypeStopLossLimit   = OrderType("STOP_LOSS_LIMIT")
	OrderTypeTakeProfit      = OrderType("TAKE_PROFIT")
	OrderTypeTakeProfitLimit = OrderType("TAKE_PROFIT_LIMIT")
	OrderTypeLimitMaker      = OrderType("LIMIT_MAKER")
)

type OrderResponseType = string

var (
	OrderResponseTypeAck    = OrderResponseType("ACK")
	OrderResponseTypeResult = OrderResponseType("RESULT")
	OrderResponseTypeFull   = OrderResponseType("FULL")
)

type OrderSide = string

var (
	OrderSideBuy  = OrderSide("BUY")
	OrderSideSell = OrderSide("SELL")
)

type TimeInForce = string

var (
	// TimeInForceGTG
	// Good Til Canceled
	//
	// An order will be on the book unless the order is canceled.
	TimeInForceGTG = TimeInForce("GTC")
	// TimeInForceIOC
	// Immediate Or Cancel
	//
	// An order will try to fill the order as much as it can before the order expires.
	TimeInForceIOC = TimeInForce("IOC")
	// TimeInForceFOK
	// Fill or Kill
	//
	// An order will expire if the full order cannot be filled upon execution.
	TimeInForceFOK = TimeInForce("FOK")
)

type Interval = string

var (
	Interval1Minute  = Interval("1m")
	Interval3Minute  = Interval("3m")
	Interval5Minute  = Interval("5m")
	Interval15Minute = Interval("15m")
	Interval30Minute = Interval("30m")
	Interval1Hour    = Interval("1h")
	Interval2Hour    = Interval("2h")
	Interval4Hour    = Interval("4h")
	Interval6Hour    = Interval("6h")
	Interval8Hour    = Interval("8h")
	Interval12Hour   = Interval("12h")
	Interval1Day     = Interval("1d")
	Interval3Day     = Interval("3d")
	Interval1Week    = Interval("1w")
	Interval1Month   = Interval("1M")
)

type RateLimiter = string

var (
	RateLimiterWeight      = RateLimiter("REQUEST_WEIGHT")
	RateLimiterOrders      = RateLimiter("ORDERS")
	RateLimiterRawRequests = RateLimiter("RAW_REQUESTS")
)

type RateLimitInterval = string

var (
	RateLimitIntervalSecond = RateLimitInterval("SECOND")
	RateLimitIntervalMinute = RateLimitInterval("MINUTE")
	RateLimitIntervalDay    = RateLimitInterval("DAY")
)

type StreamDeptLevel = string

var (
	DeptLevel5  = StreamDeptLevel("5")
	DeptLevel10 = StreamDeptLevel("10")
	DeptLevel20 = StreamDeptLevel("20")
)
