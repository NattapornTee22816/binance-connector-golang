package websocket

import (
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"github.com/buger/jsonparser"
	"time"
)

type AggregateTradeStream struct {
	EventType          string    `json:"e"`
	EventTime          time.Time `json:"E"`
	AggregateTradeId   int64     `json:"a"`
	Price              string    `json:"p"`
	Quantity           string    `json:"q"`
	FirstTradeId       int64     `json:"f"`
	LastTradeId        int64     `json:"l"`
	TradeTime          time.Time `json:"T"`
	IsBuyerMarketMaker bool      `json:"m"`
}

func parseAggregateTradeStream(b []byte) (*AggregateTradeStream, error) {
	result := new(AggregateTradeStream)

	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	}
	if v, err := jsonparser.GetInt(b, "a"); err == nil {
		result.AggregateTradeId = v
	}
	if v, err := jsonparser.GetString(b, "p"); err == nil {
		result.Price = v
	}
	if v, err := jsonparser.GetString(b, "q"); err == nil {
		result.Quantity = v
	}
	if v, err := jsonparser.GetInt(b, "f"); err == nil {
		result.FirstTradeId = v
	}
	if v, err := jsonparser.GetInt(b, "l"); err == nil {
		result.LastTradeId = v
	}
	if v, err := jsonparser.GetInt(b, "T"); err == nil {
		result.TradeTime = lib.ConvertIntToTime(v, 0)
	}
	if v, err := jsonparser.GetBoolean(b, "m"); err == nil {
		result.IsBuyerMarketMaker = v
	}

	return result, nil
}

type TradeStream struct {
	EventType          string    `json:"e"`
	EventTime          time.Time `json:"E"`
	Symbol             string    `json:"s"`
	TradeId            int64     `json:"t"`
	Price              string    `json:"p"`
	Quantity           string    `json:"q"`
	BuyerOrderId       int64     `json:"b"`
	SellerOrderId      int64     `json:"a"`
	TradeTime          time.Time `json:"T"`
	IsBuyerMarketMaker bool      `json:"m"`
}

func parseTradeStream(b []byte) (*TradeStream, error) {
	result := new(TradeStream)

	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "t"); err == nil {
		result.TradeId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "p"); err == nil {
		result.Price = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "q"); err == nil {
		result.Quantity = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "b"); err == nil {
		result.BuyerOrderId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "a"); err == nil {
		result.SellerOrderId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "a"); err == nil {
		result.TradeTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetBoolean(b, "m"); err == nil {
		result.IsBuyerMarketMaker = v
	} else {
		return nil, err
	}

	return result, nil
}

type KlineStream struct {
	EventType string          `json:"e"`
	EventTime time.Time       `json:"E"`
	Symbol    string          `json:"s"`
	Info      KlineStreamInfo `json:"k"`
}

type KlineStreamInfo struct {
	KlineStartTime           time.Time `json:"t"`
	KlineCloseTime           time.Time `json:"T"`
	Symbol                   string    `json:"s"`
	Interval                 string    `json:"i"`
	FirstTradeId             int64     `json:"f"`
	LastTradeId              int64     `json:"L"`
	OpenPrice                string    `json:"o"`
	ClosePrice               string    `json:"c"`
	HighPrice                string    `json:"h"`
	LowPrice                 string    `json:"l"`
	BaseAssetVolume          string    `json:"v"`
	NumberOfTrades           int64     `json:"n"`
	IsKlineClosed            bool      `json:"x"`
	QuoteAssetVolume         string    `json:"q"`
	TakerBuyBaseAssetVolume  string    `json:"V"`
	TakerBuyQuoteAssetVolume string    `json:"Q"`
}

func parseKlineStream(b []byte) (*KlineStream, error) {
	result := new(KlineStream)
	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	result.Info = KlineStreamInfo{}
	if v, err := jsonparser.GetInt(b, "k", "t"); err == nil {
		result.Info.KlineStartTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "k", "T"); err == nil {
		result.Info.KlineCloseTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "s"); err == nil {
		result.Info.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "i"); err == nil {
		result.Info.Interval = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "k", "f"); err == nil {
		result.Info.FirstTradeId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "k", "L"); err == nil {
		result.Info.LastTradeId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "o"); err == nil {
		result.Info.OpenPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "c"); err == nil {
		result.Info.ClosePrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "h"); err == nil {
		result.Info.HighPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "l"); err == nil {
		result.Info.LowPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "v"); err == nil {
		result.Info.BaseAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "k", "n"); err == nil {
		result.Info.NumberOfTrades = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetBoolean(b, "k", "x"); err == nil {
		result.Info.IsKlineClosed = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "q"); err == nil {
		result.Info.QuoteAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "V"); err == nil {
		result.Info.TakerBuyBaseAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "k", "Q"); err == nil {
		result.Info.TakerBuyQuoteAssetVolume = v
	} else {
		return nil, err
	}
	return result, nil
}

type IndividualMiniTickerStream struct {
	EventType                  string    `json:"e"`
	EventTime                  time.Time `json:"E"`
	Symbol                     string    `json:"s"`
	ClosePrice                 string    `json:"c"`
	OpenPrice                  string    `json:"o"`
	HighPrice                  string    `json:"h"`
	LowPrice                   string    `json:"l"`
	TotalTradeBaseAssetVolume  string    `json:"v"`
	TotalTradeQuoteAssetVolume string    `json:"q"`
}

func parseIndividualMiniTickerStream(b []byte) (*IndividualMiniTickerStream, error) {
	result := new(IndividualMiniTickerStream)
	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "c"); err == nil {
		result.ClosePrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "o"); err == nil {
		result.OpenPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "h"); err == nil {
		result.HighPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "l"); err == nil {
		result.LowPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "v"); err == nil {
		result.TotalTradeBaseAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "q"); err == nil {
		result.TotalTradeQuoteAssetVolume = v
	} else {
		return nil, err
	}
	return result, nil
}

func parseAllMiniTickerStream(b []byte) ([]*IndividualMiniTickerStream, error) {
	results := make([]*IndividualMiniTickerStream, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if v, _err := parseIndividualMiniTickerStream(value); _err == nil {
			results = append(results, v)
		}
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

type IndividualTickerStream struct {
	EventType                  string    `json:"e"`
	EventTime                  time.Time `json:"E"`
	Symbol                     string    `json:"s"`
	PriceChange                string    `json:"p"`
	PricePercentChange         string    `json:"P"`
	WeightAveragePrice         string    `json:"w"`
	FirstTraderBefore24hr      string    `json:"x"`
	LastPrice                  string    `json:"c"`
	LastQuantity               string    `json:"Q"`
	BestBidPrice               string    `json:"b"`
	BestBidQuantity            string    `json:"B"`
	BestAskPrice               string    `json:"a"`
	BestAskQuantity            string    `json:"A"`
	OpenPrice                  string    `json:"o"`
	HighPrice                  string    `json:"h"`
	LowPrice                   string    `json:"l"`
	TotalTradeBaseAssetVolume  string    `json:"v"`
	TotalTradeQuoteAssetVolume string    `json:"q"`
	StatisticsOpenTime         int64     `json:"O"`
	StatisticsCloseTime        int64     `json:"C"`
	FirstTradeId               int64     `json:"F"`
	LastTradeId                int64     `json:"L"`
	TotalNumberOfTrades        int64     `json:"n"`
}

func parseIndividualTickerStream(b []byte) (*IndividualTickerStream, error) {
	result := new(IndividualTickerStream)
	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "p"); err == nil {
		result.PriceChange = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "P"); err == nil {
		result.PricePercentChange = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "w"); err == nil {
		result.WeightAveragePrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "x"); err == nil {
		result.FirstTraderBefore24hr = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "c"); err == nil {
		result.LastPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "Q"); err == nil {
		result.LastQuantity = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "b"); err == nil {
		result.BestBidPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "B"); err == nil {
		result.BestBidQuantity = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "a"); err == nil {
		result.BestAskPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "A"); err == nil {
		result.BestAskQuantity = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "o"); err == nil {
		result.OpenPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "h"); err == nil {
		result.HighPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "l"); err == nil {
		result.LowPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "v"); err == nil {
		result.TotalTradeBaseAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "q"); err == nil {
		result.TotalTradeQuoteAssetVolume = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "O"); err == nil {
		result.StatisticsOpenTime = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "C"); err == nil {
		result.StatisticsCloseTime = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "F"); err == nil {
		result.FirstTradeId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "L"); err == nil {
		result.LastTradeId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "n"); err == nil {
		result.TotalNumberOfTrades = v
	} else {
		return nil, err
	}
	return result, nil
}

func parseAllMarketTickersStreamHandler(b []byte) ([]*IndividualTickerStream, error) {
	results := make([]*IndividualTickerStream, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if v, _err := parseIndividualTickerStream(value); _err == nil {
			results = append(results, v)
		} else {
			return
		}
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

type IndividualBookTickerStream struct {
	OrderBookUpdateId int64  `json:"u"`
	Symbol            string `json:"s"`
	BestBidPrice      string `json:"b"`
	BestBidQuantity   string `json:"B"`
	BestAskPrice      string `json:"a"`
	BestAskQuantity   string `json:"A"`
}

func parseIndividualBookTickerStream(b []byte) (*IndividualBookTickerStream, error) {
	result := new(IndividualBookTickerStream)
	if v, err := jsonparser.GetInt(b, "u"); err == nil {
		result.OrderBookUpdateId = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "b"); err == nil {
		result.BestBidPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "B"); err == nil {
		result.BestBidQuantity = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "a"); err == nil {
		result.BestAskPrice = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "A"); err == nil {
		result.BestBidQuantity = v
	} else {
		return nil, err
	}
	return result, nil
}

type PartialBookDepthStream struct {
	LastUpdateId int64                               `json:"lastUpdateId"`
	Bids         []*PartialBookDepthStreamPriceLevel `json:"bids"`
	Asks         []*PartialBookDepthStreamPriceLevel `json:"asks"`
}

type PartialBookDepthStreamPriceLevel struct {
	Price    string `json:"p"`
	Quantity string `json:"quantity"`
}

func parsePartialBookDepthStream(b []byte) (*PartialBookDepthStream, error) {
	result := new(PartialBookDepthStream)
	if v, err := jsonparser.GetInt(b, "lastUpdateId"); err == nil {
		result.LastUpdateId = v
	} else {
		return nil, err
	}
	result.Bids = make([]*PartialBookDepthStreamPriceLevel, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		bid := new(PartialBookDepthStreamPriceLevel)
		if v, _err := jsonparser.GetString(value, "[0]"); _err == nil {
			bid.Price = v
		} else {
			return
		}
		if v, _err := jsonparser.GetString(value, "[1]"); _err == nil {
			bid.Quantity = v
		} else {
			return
		}
		result.Bids = append(result.Bids, bid)
	}, "bids")
	if err != nil {
		return nil, err
	}
	result.Asks = make([]*PartialBookDepthStreamPriceLevel, 0)
	_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		ask := new(PartialBookDepthStreamPriceLevel)
		if v, _err := jsonparser.GetString(value, "[0]"); _err == nil {
			ask.Price = v
		} else {
			return
		}
		if v, _err := jsonparser.GetString(value, "[1]"); _err == nil {
			ask.Quantity = v
		} else {
			return
		}
		result.Asks = append(result.Asks, ask)
	}, "asks")
	if err != nil {
		return nil, err
	}
	return result, nil
}

type DiffDepthStream struct {
	EventType            string                       `json:"e"`
	EventTime            time.Time                    `json:"E"`
	Symbol               string                       `json:"s"`
	FirstUpdateIdInEvent int64                        `json:"U"`
	FinalUpdateIdInEvent int64                        `json:"u"`
	Bids                 []*DiffDepthStreamPriceLevel `json:"b"`
	Asks                 []*DiffDepthStreamPriceLevel `json:"a"`
}

type DiffDepthStreamPriceLevel struct {
	Price    string `json:"p"`
	Quantity string `json:"quantity"`
}

func parseDiffDepthStream(b []byte) (*DiffDepthStream, error) {
	result := new(DiffDepthStream)
	if v, err := jsonparser.GetString(b, "e"); err == nil {
		result.EventType = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "E"); err == nil {
		result.EventTime = lib.ConvertIntToTime(v, 0)
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetString(b, "s"); err == nil {
		result.Symbol = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "U"); err == nil {
		result.FirstUpdateIdInEvent = v
	} else {
		return nil, err
	}
	if v, err := jsonparser.GetInt(b, "u"); err == nil {
		result.FinalUpdateIdInEvent = v
	} else {
		return nil, err
	}
	result.Bids = make([]*DiffDepthStreamPriceLevel, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		bid := new(DiffDepthStreamPriceLevel)
		if v, _err := jsonparser.GetString(value, "[0]"); _err == nil {
			bid.Price = v
		} else {
			return
		}
		if v, _err := jsonparser.GetString(value, "[1]"); _err == nil {
			bid.Quantity = v
		} else {
			return
		}
		result.Bids = append(result.Bids, bid)
	}, "bids")
	if err != nil {
		return nil, err
	}
	result.Asks = make([]*DiffDepthStreamPriceLevel, 0)
	_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		ask := new(DiffDepthStreamPriceLevel)
		if v, _err := jsonparser.GetString(value, "[0]"); _err == nil {
			ask.Price = v
		} else {
			return
		}
		if v, _err := jsonparser.GetString(value, "[1]"); _err == nil {
			ask.Quantity = v
		} else {
			return
		}
		result.Asks = append(result.Asks, ask)
	}, "asks")
	if err != nil {
		return nil, err
	}
	return result, nil
}
