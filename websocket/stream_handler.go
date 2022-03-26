package websocket

import "github.com/buger/jsonparser"

type StreamData struct {
	Stream string `json:"stream"`
	Data   []byte `json:"data"`
}

func parseStreamData(b []byte) (*StreamData, error) {
	result := new(StreamData)

	if v, err := jsonparser.GetString(b, "stream"); err == nil {
		result.Stream = v
	} else {
		return nil, err
	}
	if v, _, _, err := jsonparser.Get(b, "data"); err == nil {
		result.Data = v
	} else {
		return nil, err
	}

	return result, nil
}

// AggTradeStreamHandler
//
// func(stream string, value *AggregateTradeStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type AggTradeStreamHandler = func(string, *AggregateTradeStream, error)

func (s *Stream) callAggTradeStreamHandler(stream string, data *AggregateTradeStream) {
	var err error
	for _, handler := range s.aggTradeStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// TradeStreamHandler
//
// func(stream string, value *TradeStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type TradeStreamHandler = func(string, *TradeStream, error)

func (s *Stream) callTradeStreamHandler(stream string, data *TradeStream) {
	var err error
	for _, handler := range s.tradeStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// KlineStreamHandler
//
// func(stream string, value *KlineStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type KlineStreamHandler = func(string, *KlineStream, error)

func (s *Stream) callKlineStreamHandler(stream string, data *KlineStream) {
	var err error
	for _, handler := range s.klineStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// IndividualMiniTickerStreamHandler
//
// func(stream string, value *IndividualMiniTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type IndividualMiniTickerStreamHandler = func(string, *IndividualMiniTickerStream, error)

func (s *Stream) callIndividualMiniTickerStreamHandler(stream string, data *IndividualMiniTickerStream) {
	var err error
	for _, handler := range s.individualMiniTickerStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// AllMarketMiniTickerStreamHandler
//
// func(stream string, value []*IndividualMiniTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type AllMarketMiniTickerStreamHandler = func(string, []*IndividualMiniTickerStream, error)

func (s *Stream) callAllMarketMiniTickerStreamHandler(stream string, data []*IndividualMiniTickerStream) {
	var err error
	for _, handler := range s.allMarketMiniTickerStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// IndividualTickerStreamHandler
//
// func(stream string, value *IndividualTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type IndividualTickerStreamHandler = func(string, *IndividualTickerStream, error)

func (s *Stream) callIndividualTickerStreamHandler(stream string, data *IndividualTickerStream) {
	var err error
	for _, handler := range s.individualTickerStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// AllMarketTickersStreamHandler
//
// func(stream string, value *IndividualTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type AllMarketTickersStreamHandler = func(string, []*IndividualTickerStream, error)

func (s *Stream) callAllMarketTickersStreamHandler(stream string, data []*IndividualTickerStream) {
	var err error
	for _, handler := range s.allMarketTickersStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// IndividualBookTickerStreamHandler
//
// func(stream string, value *IndividualBookTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type IndividualBookTickerStreamHandler = func(string, *IndividualBookTickerStream, error)

func (s *Stream) callIndividualBookTickerStreamHandler(stream string, data *IndividualBookTickerStream) {
	var err error
	for _, handler := range s.individualBookTickerStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// AllBookTickerStreamHandler
//
// func(stream string, value *IndividualBookTickerStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type AllBookTickerStreamHandler = func(string, *IndividualBookTickerStream, error)

func (s *Stream) callAllBookTickerStreamHandler(stream string, data *IndividualBookTickerStream) {
	var err error
	for _, handler := range s.individualBookTickerStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// PartialBookDepthStreamHandler
// PartialBookDepthStreamType & NewPartialBookDepth100msStreamType
//
// func(stream string, value *PartialBookDepthStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type PartialBookDepthStreamHandler = func(string, *PartialBookDepthStream, error)

func (s *Stream) callPartialBookDepthStreamHandler(stream string, data *PartialBookDepthStream) {
	var err error
	for _, handler := range s.partialBookDepthStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}

// DiffDepthStreamHandler
// DiffDepthStreamType & DiffDepth100msStreamType
//
// func(stream string, value *DiffDepthStream, err error) {
//   do something
//   when have error then set to err and return for stop next handler
// }
type DiffDepthStreamHandler = func(string, *DiffDepthStream, error)

func (s *Stream) callDiffDepthStreamHandler(stream string, data *DiffDepthStream) {
	var err error
	for _, handler := range s.diffDepthStreamHandler {
		handler(stream, data, err)
		if err != nil {
			break
		}
	}
}
