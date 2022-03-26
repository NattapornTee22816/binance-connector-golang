package model

import "testing"

// OrderBook
func TestParser_ParseOrderBook(t *testing.T) {
	bytes := []byte(`{"lastUpdateId":1027024,"bids":[["4.00000000","431.00000000"]],"asks":[["4.00000200","12.00000000"]]}`)

	parser := NewParser()
	_, err := parser.ParseOrderBook(bytes)
	if err != nil {
		t.Error(err)
	}
}

// RecentTradesList
func TestParser_ParseRecentTrade(t *testing.T) {
	bytes := []byte(`[{"id":28457,"price":"4.00000100","qty":"12.00000000","quoteQty":"48.000012","time":1499865549590,"isBuyerMaker":true,"isBestMatch":true}]`)

	parser := NewParser()
	_, err := parser.ParseRecentTrade(bytes)
	if err != nil {
		t.Error(err)
	}
}

// OldTradeLookup
func TestParser_ParseOldTradeLookup(t *testing.T) {
	bytes := []byte(`[{"id":28457,"price":"4.00000100","qty":"12.00000000","quoteQty":"48.000012","time":1499865549590,"isBuyerMaker":true,"isBestMatch":true}]`)

	parser := NewParser()
	_, err := parser.ParseOldTradeLookup(bytes)
	if err != nil {
		t.Error(err)
	}
}

// Compressed/Aggregate Trades List
func TestParser_ParseAggregateTrade(t *testing.T) {
	bytes := []byte(`[{"a":26129,"p":"0.01633102","q":"4.70443515","f":27781,"l":27781,"T":1498793709153,"m":true,"M":true}]`)

	parser := NewParser()
	_, err := parser.ParseAggregateTrade(bytes)
	if err != nil {
		t.Error(err)
	}
}

// Klines (Kline/Candlestick Data)
func (r *Parser) TestParser_ParseKline(t *testing.T) {
	bytes := []byte(`[[1499040000000,"0.01634790","0.80000000","0.01575800","0.01577100","148976.11427815",1499644799999,"2434.19055334",308,"1756.87402397","28.46694368","17928899.62484339"]]`)

	parser := NewParser()
	_, err := parser.ParseKline(bytes)
	if err != nil {
		t.Error(err)
	}
}

// Ticker24hr
func TestParser_ParseTicker24hr_SingleItem(t *testing.T) {
	bytes := []byte(`{"symbol":"BNBBTC","priceChange":"-94.99999800","priceChangePercent":"-95.960","weightedAvgPrice":"0.29628482","prevClosePrice":"0.10002000","lastPrice":"4.00000200","lastQty":"200.00000000","bidPrice":"4.00000000","bidQty":"100.00000000","askPrice":"4.00000200","askQty":"100.00000000","openPrice":"99.00000000","highPrice":"100.00000000","lowPrice":"0.10000000","volume":"8913.30000000","quoteVolume":"15.30000000","openTime":1499783499040,"closeTime":1499869899040,"firstId":28385,"lastId":28460,"count":76}`)

	parser := NewParser()
	_, err := parser.ParseTicker24hr(bytes)
	if err != nil {
		t.Error(err)
	}
}

func TestParser_ParseTicker24hr_ArrayItem(t *testing.T) {
	bytes := []byte(`[{"symbol":"BNBBTC","priceChange":"-94.99999800","priceChangePercent":"-95.960","weightedAvgPrice":"0.29628482","prevClosePrice":"0.10002000","lastPrice":"4.00000200","lastQty":"200.00000000","bidPrice":"4.00000000","bidQty":"100.00000000","askPrice":"4.00000200","askQty":"100.00000000","openPrice":"99.00000000","highPrice":"100.00000000","lowPrice":"0.10000000","volume":"8913.30000000","quoteVolume":"15.30000000","openTime":1499783499040,"closeTime":1499869899040,"firstId":28385,"lastId":28460,"count":76}]`)

	parser := NewParser()
	_, err := parser.ParseTicker24hr(bytes)
	if err != nil {
		t.Error(err)
	}
}

// TickerPrice
func TestParser_ParseTickerPrice_SingleItem(t *testing.T) {
	bytes := []byte(`{"symbol":"LTCBTC","price":"4.00000200"}`)

	parser := NewParser()
	_, err := parser.ParseTickerPrice(bytes)
	if err != nil {
		t.Error(err)
	}
}

func TestParser_ParseTickerPrice_ArrayItem(t *testing.T) {
	bytes := []byte(`[{"symbol":"LTCBTC","price":"4.00000200"},{"symbol":"ETHBTC","price":"0.07946600"}]`)

	parser := NewParser()
	_, err := parser.ParseTickerPrice(bytes)
	if err != nil {
		t.Error(err)
	}
}

// BookTicker
func TestParser_ParseBookTicker_SingleItem(t *testing.T) {
	bytes := []byte(`{"symbol":"LTCBTC","bidPrice":"4.00000000","bidQty":"431.00000000","askPrice":"4.00000200","askQty":"9.00000000"}`)

	parser := NewParser()
	_, err := parser.ParseBookTicker(bytes)
	if err != nil {
		t.Error(err)
	}
}

func TestParser_ParseBookTicker_ArrayItem(t *testing.T) {
	bytes := []byte(`[{"symbol":"LTCBTC","bidPrice":"4.00000000","bidQty":"431.00000000","askPrice":"4.00000200","askQty":"9.00000000"},{"symbol":"ETHBTC","bidPrice":"0.07946700","bidQty":"9.00000000","askPrice":"100000.00000000","askQty":"1000.00000000"}]`)

	parser := NewParser()
	_, err := parser.ParseBookTicker(bytes)
	if err != nil {
		t.Error(err)
	}
}
