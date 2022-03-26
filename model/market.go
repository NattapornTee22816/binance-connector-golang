package model

import (
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"github.com/buger/jsonparser"
	"time"
)

// ExchangeInformation

type ExchangeInformationParam struct {
	Symbol  string   `json:"symbol" param:"symbol"`
	Symbols []string `json:"symbols" param:"symbols"`
}

type ExchangeInformation struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int64  `json:"intervalNum"`
		Limit         int64  `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []struct{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                     string   `json:"symbol"`
		Status                     string   `json:"status"`
		BaseAsset                  string   `json:"baseAsset"`
		BaseAssetPrecision         int64    `json:"baseAssetPrecision"`
		QuoteAsset                 string   `json:"quoteAsset"`
		QuotePrecision             int64    `json:"quotePrecision"`
		QuoteAssetPrecision        int64    `json:"quoteAssetPrecision"`
		BaseCommissionPrecision    int64    `json:"baseCommissionPrecision"`
		QuoteCommissionPrecision   int64    `json:"quoteCommissionPrecision"`
		OrderTypes                 []string `json:"orderTypes"`
		IcebergAllowed             bool     `json:"icebergAllowed"`
		OcoAllowed                 bool     `json:"ocoAllowed"`
		QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
		AllowTrailingStop          bool     `json:"allowTrailingStop"`
		IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
		Filters                    []struct {
			FilterType       string `json:"filterType"`
			MinPrice         string `json:"minPrice,omitempty"`
			MaxPrice         string `json:"maxPrice,omitempty"`
			TickSize         string `json:"tickSize,omitempty"`
			MultiplierUp     string `json:"multiplierUp,omitempty"`
			MultiplierDown   string `json:"multiplierDown,omitempty"`
			AvgPriceMins     int64  `json:"avgPriceMins,omitempty"`
			MinQty           string `json:"minQty,omitempty"`
			MaxQty           string `json:"maxQty,omitempty"`
			StepSize         string `json:"stepSize,omitempty"`
			MinNotional      string `json:"minNotional,omitempty"`
			ApplyToMarket    bool   `json:"applyToMarket,omitempty"`
			Limit            int64  `json:"limit,omitempty"`
			MaxNumOrders     int64  `json:"maxNumOrders,omitempty"`
			MaxNumAlgoOrders int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
		Permissions []string `json:"permissions"`
	} `json:"symbols"`
}

// OrderBook

type OrderBookParam struct {
	Symbol string `json:"symbol" param:"symbol" validate:"required"`
	Limit  int64  `json:"limit" param:"limit" validate:"min=0,max=5000"`
}

type OrderBook struct {
	LastUpdateId int64             `json:"lastUpdateId"`
	Bids         []*OrderBookPrice `json:"bids"`
	Asks         []*OrderBookPrice `json:"asks"`
}

type OrderBookPrice struct {
	Price string `json:"price"`
	Qty   string `json:"qty"`
}

func (r *Parser) ParseOrderBook(b []byte) (*OrderBook, error) {
	result := new(OrderBook)
	bids := make([]*OrderBookPrice, 0)
	asks := make([]*OrderBookPrice, 0)

	if v, err := jsonparser.GetInt(b, "lastUpdateId"); err == nil {
		result.LastUpdateId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}

	parseOrderBookType := 0
	parseOrderBook := func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		orderBook := new(OrderBookPrice)

		if v, err := jsonparser.GetString(value, "[0]"); err == nil {
			orderBook.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[1]"); err == nil {
			orderBook.Qty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		switch parseOrderBookType {
		case 0: // bids
			bids = append(bids, orderBook)
		case 1:
			asks = append(asks, orderBook)
		}
	}

	if _, err := jsonparser.ArrayEach(b, parseOrderBook, "bids"); err == nil {
		result.Bids = bids
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}

	parseOrderBookType = 1
	if _, err := jsonparser.ArrayEach(b, parseOrderBook, "asks"); err == nil {
		result.Asks = asks
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// RecentTradesList

type RecentTradeParam struct {
	Symbol string `json:"symbol" param:"symbol" validate:"required"`
	Limit  int64  `json:"limit" param:"limit" validate:"min=0,max=1000"`
}

type RecentTrade struct {
	Id           int64     `json:"id"`
	Price        string    `json:"price"`
	Qty          string    `json:"qty"`
	QuoteQty     string    `json:"quoteQty"`
	Time         time.Time `json:"time"`
	IsBuyerMaker bool      `json:"isBuyerMaker"`
	IsBestMatch  bool      `json:"isBestMatch"`
}

func (r *Parser) ParseRecentTrade(b []byte) ([]*RecentTrade, error) {
	results := make([]*RecentTrade, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(RecentTrade)

		if v, err := jsonparser.GetInt(value, "id"); err == nil {
			item.Id = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "qty"); err == nil {
			item.Qty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "quoteQty"); err == nil {
			item.QuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "time"); err == nil {
			item.Time = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBuyerMaker"); err == nil {
			item.IsBuyerMaker = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBestMatch"); err == nil {
			item.IsBestMatch = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// OldTradeLookup

type OldTradeLookupParam struct {
	Symbol string `json:"symbol" param:"symbol" validate:"required"`
	Limit  int64  `json:"limit" param:"limit" validate:"min=0,max=1000"`
	FromId int64  `json:"from_id" param:"fromId"`
}

type OldTradeLookup struct {
	Id           int64     `json:"id"`
	Price        string    `json:"price"`
	Qty          string    `json:"qty"`
	QuoteQty     string    `json:"quoteQty"`
	Time         time.Time `json:"time"`
	IsBuyerMaker bool      `json:"isBuyerMaker"`
	IsBestMatch  bool      `json:"isBestMatch"`
}

func (r *Parser) ParseOldTradeLookup(b []byte) ([]*OldTradeLookup, error) {
	results := make([]*OldTradeLookup, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(OldTradeLookup)

		if v, err := jsonparser.GetInt(value, "id"); err == nil {
			item.Id = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "qty"); err == nil {
			item.Qty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "quoteQty"); err == nil {
			item.QuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "time"); err == nil {
			item.Time = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBuyerMaker"); err == nil {
			item.IsBuyerMaker = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBestMatch"); err == nil {
			item.IsBestMatch = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// Compressed/Aggregate Trades List

type AggregateTradeParam struct {
	Symbol    string    `json:"symbol" param:"symbol" validate:"required"`
	FromId    int64     `json:"fromId" param:"fromId"`
	StartTime time.Time `json:"startTime" param:"startTime"`
	EndTime   time.Time `json:"endTime" param:"endTime"`
	Limit     int64     `json:"limit" param:"limit" validate:"min=0,max=1000"`
}

type AggregateTrade struct {
	TradeId      int64     `json:"a"`
	Price        string    `json:"p"`
	Quantity     string    `json:"q"`
	FirstTradeId int64     `json:"f"`
	LastTradeId  int64     `json:"l"`
	Timestamp    time.Time `json:"T"`
	IsBuyerMaker bool      `json:"m"`
	IsBestMatch  bool      `json:"M"`
}

func (r *Parser) ParseAggregateTrade(b []byte) ([]*AggregateTrade, error) {
	results := make([]*AggregateTrade, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(AggregateTrade)

		if v, err := jsonparser.GetInt(b, "a"); err == nil {
			item.TradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "p"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "q"); err == nil {
			item.Quantity = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "f"); err == nil {
			item.FirstTradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "l"); err == nil {
			item.LastTradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "T"); err == nil {
			item.Timestamp = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(b, "m"); err == nil {
			item.IsBuyerMaker = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(b, "M"); err == nil {
			item.IsBestMatch = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// Klines (Kline/Candlestick Data)

type KlineParam struct {
	Symbol    string    `json:"symbol" param:"symbol" validate:"required"`
	Interval  Interval  `json:"interval" param:"interval" validate:"required"`
	StartTime time.Time `json:"startTime" param:"startTime"`
	EndTime   time.Time `json:"endTime" param:"endTime"`
	Limit     int64     `json:"limit" param:"limit" validate:"min=0,max=1000"`
}

type Kline struct {
	OpenTime                 time.Time `json:"openTime"`
	Open                     string    `json:"open"`
	High                     string    `json:"high"`
	Low                      string    `json:"low"`
	Close                    string    `json:"close"`
	Volume                   string    `json:"volume"`
	CloseTime                time.Time `json:"closeTime"`
	QuoteAssetVolume         string    `json:"quoteAssetVolume"`
	NumberOfTrades           int64     `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string    `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string    `json:"takerBuyQuoteAssetVolume"`
}

func (r *Parser) ParseKline(b []byte) ([]*Kline, error) {
	results := make([]*Kline, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(Kline)

		if v, err := jsonparser.GetInt(value, "[0]"); err == nil {
			item.OpenTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[1]"); err == nil {
			item.Open = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[2]"); err == nil {
			item.High = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[3]"); err == nil {
			item.Low = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[4]"); err == nil {
			item.Close = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[5]"); err == nil {
			item.Volume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "[6]"); err == nil {
			item.CloseTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[7]"); err == nil {
			item.QuoteAssetVolume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "[8]"); err == nil {
			item.NumberOfTrades = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[9]"); err == nil {
			item.TakerBuyBaseAssetVolume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "[10]"); err == nil {
			item.TakerBuyQuoteAssetVolume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// AveragePrice (Current Average Price)

type AveragePriceParam struct {
	Symbol string `json:"symbol" param:"symbol" validate:"required"`
}

type AveragePrice struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

// Ticker24hr

type Ticker24hrParam struct {
	Symbol string `json:"symbol" param:"symbol"`
}

type Ticker24hr struct {
	Symbol             string    `json:"symbol"`
	PriceChange        string    `json:"priceChange"`
	PriceChangePercent string    `json:"priceChangePercent"`
	WeightedAvgPrice   string    `json:"weightedAvgPrice"`
	PrevClosePrice     string    `json:"prevClosePrice"`
	LastPrice          string    `json:"lastPrice"`
	LastQty            string    `json:"lastQty"`
	BidPrice           string    `json:"bidPrice"`
	BidQty             string    `json:"bidQty"`
	AskPrice           string    `json:"askPrice"`
	AskQty             string    `json:"askQty"`
	OpenPrice          string    `json:"openPrice"`
	HighPrice          string    `json:"highPrice"`
	LowPrice           string    `json:"lowPrice"`
	Volume             string    `json:"volume"`
	QuoteVolume        string    `json:"quoteVolume"`
	OpenTime           time.Time `json:"openTime"`
	CloseTime          time.Time `json:"closeTime"`
	FirstTradeId       int64     `json:"firstId"`
	LastTradeId        int64     `json:"lastId"`
	TradeCount         int64     `json:"count"`
}

func (r *Parser) ParseTicker24hr(b []byte) ([]*Ticker24hr, error) {
	b = lib.BytesToJsonArray(b)
	results := make([]*Ticker24hr, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(Ticker24hr)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "priceChange"); err == nil {
			item.PriceChange = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "priceChangePercent"); err == nil {
			item.PriceChangePercent = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "weightedAvgPrice"); err == nil {
			item.WeightedAvgPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "prevClosePrice"); err == nil {
			item.PrevClosePrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "lastPrice"); err == nil {
			item.LastPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "lastQty"); err == nil {
			item.LastQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "bidPrice"); err == nil {
			item.BidPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "bidQty"); err == nil {
			item.BidQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "askPrice"); err == nil {
			item.AskPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "askQty"); err == nil {
			item.AskQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "openPrice"); err == nil {
			item.OpenPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "highPrice"); err == nil {
			item.HighPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "lowPrice"); err == nil {
			item.LowPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "volume"); err == nil {
			item.Volume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "quoteVolume"); err == nil {
			item.QuoteVolume = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "openTime"); err == nil {
			item.OpenTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "closeTime"); err == nil {
			item.CloseTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "firstId"); err == nil {
			item.FirstTradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "lastId"); err == nil {
			item.LastTradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "count"); err == nil {
			item.TradeCount = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// TickerPrice

type TickerPriceParam struct {
	Symbol string `json:"symbol" param:"symbol"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (r *Parser) ParseTickerPrice(b []byte) ([]*TickerPrice, error) {
	b = lib.BytesToJsonArray(b)
	results := make([]*TickerPrice, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(TickerPrice)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// BookTicker

type BookTickerParam struct {
	Symbol string `json:"symbol" param:"symbol"`
}

type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func (r *Parser) ParseBookTicker(b []byte) ([]*BookTicker, error) {
	b = lib.BytesToJsonArray(b)
	results := make([]*BookTicker, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(BookTicker)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "bidPrice"); err == nil {
			item.BidPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "bidQty"); err == nil {
			item.BidQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "askPrice"); err == nil {
			item.AskPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "askQty"); err == nil {
			item.AskQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}
