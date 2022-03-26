package websocket

import (
	"github.com/NattapornTee22816/binance-connector-golang/model"
	"strings"
)

type StreamType = string

var (
	AggregateTradeStreamType        = StreamType("<symbol>@aggTrade")
	TradeStreamType                 = StreamType("<symbol>@trade")
	KlineStreamType                 = StreamType("<symbol>@kline_<interval>")
	IndividualMiniTickerStreamType  = StreamType("<symbol>@miniTicker")
	AllMarketMiniTickersStreamType  = StreamType("!miniTicker@arr")
	IndividualTickerStreamType      = StreamType("<symbol>@ticker")
	AllMarketTickersStreamType      = StreamType("!ticker@arr")
	IndividualBookTickerStreamType  = StreamType("<symbol>@bookTicker")
	AllBookTickersStreamType        = StreamType("!bookTicker")
	PartialBookDepthStreamType      = StreamType("<symbol>@depth<levels>")
	PartialBookDepth100msStreamType = StreamType("<symbol>@depth<levels>@100ms")
	DiffDepthStreamType             = StreamType("<symbol>@depth")
	DiffDepth100msStreamType        = StreamType("<symbol>@depth@100ms")
)

// NewAggregateTradeStreamType
//  - create stream name '<symbol>@aggTrade'
//  - The Aggregate Trade Streams push trade information that is aggregated for a single taker order.
//  - https://binance-docs.github.io/apidocs/spot/en/#aggregate-trade-streams
func NewAggregateTradeStreamType(symbol string) (string, error) {
	return newStreamSymbol(AggregateTradeStreamType, symbol, "", "")
}

func (s *Stream) SubscribeAggregateTradeStreams(streams []string, handler ...AggTradeStreamHandler) error {
	if len(handler) == 0 && len(s.aggTradeStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@aggTrade"); err != nil {
		return err
	}

	err := s.subscribe(streams)
	if err != nil {
		return err
	}

	s.appendStreams(AggregateTradeStreamType, streams)
	s.aggTradeStreamHandler = append(s.aggTradeStreamHandler, handler...)
	return nil
}

// NewTradeStreamType
//  - create stream name '<symbol>@trade'
//  - The Trade Streams push raw trade information; each trade has a unique buyer and seller.
//  - https://binance-docs.github.io/apidocs/spot/en/#trade-streams
func NewTradeStreamType(symbol string) (string, error) {
	return newStreamSymbol(TradeStreamType, symbol, "", "")
}

func (s *Stream) SubscribeTradeStreams(streams []string, handler ...TradeStreamHandler) error {
	if len(handler) == 0 && len(s.aggTradeStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@trade"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(TradeStreamType, streams)
	s.tradeStreamHandler = append(s.tradeStreamHandler, handler...)
	return nil
}

// NewKlineStreamType
//  - create stream name '<symbol>@kline_<interval>'
//  - The Kline/Candlestick Stream push updates to the current klines/candlestick every second.
//  - https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-streams
func NewKlineStreamType(symbol string, interval model.Interval) (string, error) {
	return newStreamSymbol(KlineStreamType, symbol, interval, "")
}

func (s *Stream) SubscribeKlineStreams(streams []string, handler ...KlineStreamHandler) error {
	if len(handler) == 0 && len(s.klineStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@kline_(1m|3m|5m|15m|30m|1h|2h|4h|6h|8h|12h|1d|3d|1w|1M)"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(KlineStreamType, streams)
	s.klineStreamHandler = append(s.klineStreamHandler, handler...)
	return nil
}

// NewIndividualMiniTickerStreamType (Individual Symbol Mini Ticker Stream)
//  - create stream name '<symbol>@miniTicker'
//  - 24hr rolling window mini-ticker statistics. These are NOT the statistics of the UTC day,
//    but a 24hr rolling window for the previous 24hrs.
//  - https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-mini-ticker-stream
func NewIndividualMiniTickerStreamType(symbol string) (string, error) {
	return newStreamSymbol(IndividualMiniTickerStreamType, symbol, "", "")
}

func (s *Stream) SubscribeIndividualMiniTickerStreams(streams []string, handler ...IndividualMiniTickerStreamHandler) error {
	if len(handler) == 0 && len(s.aggTradeStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@miniTicker"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(IndividualMiniTickerStreamType, streams)
	s.individualMiniTickerStreamHandler = append(s.individualMiniTickerStreamHandler, handler...)
	return nil
}

// SubscribeAllMarketMiniTickersStreams
//  - 24hr rolling window mini-ticker statistics for all symbols that changed in an array.
//    These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
//  - Note that only tickers that have changed will be present in the array.
//  - https://binance-docs.github.io/apidocs/spot/en/#all-market-mini-tickers-stream
func (s *Stream) SubscribeAllMarketMiniTickersStreams(handler ...AllMarketMiniTickerStreamHandler) error {
	if len(handler) == 0 && len(s.allMarketMiniTickerStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	streams := make([]string, 0)
	streams = append(streams, "!miniTicker@arr")
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(AllMarketMiniTickersStreamType, streams)
	s.allMarketMiniTickerStreamHandler = append(s.allMarketMiniTickerStreamHandler, handler...)
	return nil
}

// NewIndividualTickerStreamType
//  - create stream name '<symbol>@ticker'
//  - 24hr rolling window ticker statistics for a single symbol.
//    These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
//  - https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-ticker-streams
func NewIndividualTickerStreamType(symbol string) (string, error) {
	return newStreamSymbol(IndividualTickerStreamType, symbol, "", "")
}

func (s *Stream) SubscribeIndividualTickerStream(streams []string, handler ...IndividualTickerStreamHandler) error {
	if len(handler) == 0 && len(s.individualTickerStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@ticker"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(IndividualTickerStreamType, streams)
	s.individualTickerStreamHandler = append(s.individualTickerStreamHandler, handler...)
	return nil
}

// SubscribeAllMarketTickersStream
//  - 24hr rolling window ticker statistics for all symbols that changed in an array.
//    These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
//    Note that only tickers that have changed will be present in the array.
//  - https://binance-docs.github.io/apidocs/spot/en/#all-market-tickers-stream
func (s *Stream) SubscribeAllMarketTickersStream(handler ...AllMarketTickersStreamHandler) error {
	if len(handler) == 0 && len(s.allMarketTickersStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	streams := make([]string, 0)
	streams = append(streams, "!ticker@arr")
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(AllMarketTickersStreamType, streams)
	s.allMarketTickersStreamHandler = append(s.allMarketTickersStreamHandler, handler...)
	return nil
}

// NewIndividualBookTickerStreamType
//
// create stream name '<symbol>@bookTicker'
func NewIndividualBookTickerStreamType(symbol string) (string, error) {
	return newStreamSymbol(IndividualBookTickerStreamType, symbol, "", "")
}

func (s *Stream) SubscribeIndividualBookTickerStream(streams []string, handler ...IndividualBookTickerStreamHandler) error {
	if len(handler) == 0 && len(s.individualBookTickerStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@bookTicker"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(IndividualBookTickerStreamType, streams)
	s.individualBookTickerStreamHandler = append(s.individualBookTickerStreamHandler, handler...)
	return nil
}

// SubscribeAllBookTickerStream
// - Pushes any update to the best bid or ask's price or quantity in real-time for all symbols.
// - https://binance-docs.github.io/apidocs/spot/en/#all-book-tickers-stream
func (s *Stream) SubscribeAllBookTickerStream(handler ...AllBookTickerStreamHandler) error {
	if len(handler) == 0 && len(s.allBookTickerStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	streams := make([]string, 0)
	streams = append(streams, "!bookTicker")
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(AllBookTickersStreamType, streams)
	s.allBookTickerStreamHandler = append(s.allBookTickerStreamHandler, handler...)
	return nil
}

// NewPartialBookDepthStreamType
//  - create stream name '<symbol>@depth<levels>'
//  - use SubscribePartialBookDepthStream for subscribe
func NewPartialBookDepthStreamType(symbol string, levels model.StreamDeptLevel) (string, error) {
	return newStreamSymbol(PartialBookDepthStreamType, symbol, "", levels)
}

// NewPartialBookDepth100msStreamType
//  - create stream name '<symbol>@depth<levels>@100ms'
//  - use SubscribePartialBookDepthStream for subscribe
func NewPartialBookDepth100msStreamType(symbol string, levels model.StreamDeptLevel) (string, error) {
	return newStreamSymbol(PartialBookDepth100msStreamType, symbol, "", levels)
}

// SubscribePartialBookDepthStream
//  - Top bids and asks, Valid are 5, 10, or 20.
//  - https://binance-docs.github.io/apidocs/spot/en/#partial-book-depth-streams
func (s *Stream) SubscribePartialBookDepthStream(streams []string, handler ...PartialBookDepthStreamHandler) error {
	if len(handler) == 0 && len(s.partialBookDepthStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@depth(5|10|20)(@100ms)?"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(PartialBookDepthStreamType, streams)
	s.partialBookDepthStreamHandler = append(s.partialBookDepthStreamHandler, handler...)
	return nil
}

// NewDiffDepthStreamType
//  - create stream name '<symbol>@depth'
//  - use SubscribeDiffDepthStream for subscribe
func NewDiffDepthStreamType(symbol string) (string, error) {
	return newStreamSymbol(DiffDepthStreamType, symbol, "", "")
}

// NewDiffDepth100msStreamType
//  - create stream name '<symbol>@depth@100ms'
//  - use SubscribeDiffDepthStream for subscribe
func NewDiffDepth100msStreamType(symbol string) (string, error) {
	return newStreamSymbol(DiffDepth100msStreamType, symbol, "", "")
}

// SubscribeDiffDepthStream
//  - Order book price and quantity depth updates used to locally manage an order book.
//  - https://binance-docs.github.io/apidocs/spot/en/#diff-depth-stream
func (s *Stream) SubscribeDiffDepthStream(streams []string, handler ...DiffDepthStreamHandler) error {
	if len(handler) == 0 && len(s.diffDepthStreamHandler) == 0 {
		return ErrNoStreamHandler
	}

	if err := s.validateStreams(streams, "[a-z0-9]+@depth(@100ms)?"); err != nil {
		return err
	}
	if err := s.subscribe(streams); err != nil {
		return err
	}

	s.appendStreams(DiffDepthStreamType, streams)
	s.diffDepthStreamHandler = append(s.diffDepthStreamHandler, handler...)
	return nil
}

// newStreamSymbol
//
// create stream name with pattern
// 	- <symbol>
//  - <interval>
//  - <levels>
func newStreamSymbol(streamType StreamType, symbol string, interval model.Interval, levels model.StreamDeptLevel) (string, error) {
	if len(streamType) == 0 {
		return "", &ErrStreamParameterRequired{
			Message: "parameter 'streamType' is required",
		}
	}
	if strings.Contains(streamType, "<symbol>") && len(symbol) == 0 {
		return "", &ErrStreamParameterRequired{
			Message: "parameter 'symbol' is required",
		}
	} else {
		streamType = strings.Replace(streamType, "<symbol>", strings.ToLower(symbol), 1)
	}
	if strings.Contains(streamType, "<interval>") && len(symbol) == 0 {
		return "", &ErrStreamParameterRequired{
			Message: "parameter 'interval' is required",
		}
	} else {
		streamType = strings.Replace(streamType, "<interval>", interval, 1)
	}
	if strings.Contains(streamType, "<levels>") && len(symbol) == 0 {
		return "", &ErrStreamParameterRequired{
			Message: "parameter 'levels' is required",
		}
	} else {
		streamType = strings.Replace(streamType, "<levels>", levels, 1)
	}

	return streamType, nil
}
