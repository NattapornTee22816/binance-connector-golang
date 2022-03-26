package websocket

import (
	"errors"
	"fmt"
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"regexp"
	"sync"
	"time"
)

var (
	ErrNoStreamHandler     = errors.New("not found stream handler")
	ErrRequireStreamSymbol = errors.New("stream is required")
	ErrStreamSymbolInvalid = errors.New("stream invalid pattern")
)

type Stream struct {
	ws                                *Websocket
	requestId                         uint64
	streams                           map[string]StreamType
	aggTradeStreamHandler             []AggTradeStreamHandler
	tradeStreamHandler                []TradeStreamHandler
	klineStreamHandler                []KlineStreamHandler
	individualMiniTickerStreamHandler []IndividualMiniTickerStreamHandler
	allMarketMiniTickerStreamHandler  []AllMarketMiniTickerStreamHandler
	individualTickerStreamHandler     []IndividualTickerStreamHandler
	allMarketTickersStreamHandler     []AllMarketTickersStreamHandler
	individualBookTickerStreamHandler []IndividualBookTickerStreamHandler
	allBookTickerStreamHandler        []AllBookTickerStreamHandler
	partialBookDepthStreamHandler     []PartialBookDepthStreamHandler
	diffDepthStreamHandler            []DiffDepthStreamHandler
}

// NewWsStream
//
// https://binance-docs.github.io/apidocs/spot/en/#websocket-market-streams
func NewWsStream() (*Stream, error) {
	wss := &Stream{
		requestId:                         1,
		streams:                           make(map[string]string),
		aggTradeStreamHandler:             make([]AggTradeStreamHandler, 0),
		tradeStreamHandler:                make([]TradeStreamHandler, 0),
		klineStreamHandler:                make([]KlineStreamHandler, 0),
		individualMiniTickerStreamHandler: make([]IndividualMiniTickerStreamHandler, 0),
		allMarketMiniTickerStreamHandler:  make([]AllMarketMiniTickerStreamHandler, 0),
		individualTickerStreamHandler:     make([]IndividualTickerStreamHandler, 0),
		allMarketTickersStreamHandler:     make([]AllMarketTickersStreamHandler, 0),
		individualBookTickerStreamHandler: make([]IndividualBookTickerStreamHandler, 0),
		allBookTickerStreamHandler:        make([]AllBookTickerStreamHandler, 0),
		partialBookDepthStreamHandler:     make([]PartialBookDepthStreamHandler, 0),
		diffDepthStreamHandler:            make([]DiffDepthStreamHandler, 0),
	}
	wss.ws = &Websocket{
		id:           lib.RandomInt(),
		url:          "wss://stream.binance.com:9443/stream",
		PingDuration: 2 * time.Minute,
		PongDuration: 5 * time.Minute,
		mu:           sync.Mutex{},
		logger:       lib.NewLogger("ws-binance", lib.LogLevelDebug),
		wg:           sync.WaitGroup{},
		OnConnect:    wss.onWebsocketConnect,
	}

	wss.ws.Connect()
	go wss.readMessage()

	return wss, nil
}

func (s *Stream) onWebsocketConnect(ws *Websocket) {
	if len(s.streams) > 0 {
		streams := make([]string, 0)
		for key := range s.streams {
			streams = append(streams, key)
		}

		if ws.logger.CanDebug() {
			ws.logger.Debug("auto subscribe after connect")
		}

		err := s.subscribe(streams)
		if err != nil {
			ws.logger.Error(fmt.Sprintf("auto subscribe error: %s", err.Error()))
		}
	}
}

func (s *Stream) Shutdown() {
	s.ws.Shutdown()
}

func (s *Stream) validateStreams(streams []string, pattern string) error {
	if len(streams) == 0 {
		return ErrRequireStreamSymbol
	}
	regex := regexp.MustCompile(pattern)

	for _, stream := range streams {
		if !regex.MatchString(stream) {
			return ErrStreamSymbolInvalid
		}
	}

	return nil
}

func (s *Stream) appendStreams(streamType StreamType, streams []string) {
	for _, stream := range streams {
		s.streams[stream] = streamType
	}
}

func (s *Stream) subscribe(streams []string) error {
	command := newCommandSubscribe(s.requestId, streams)
	s.requestId = s.requestId + 1

	err := s.ws.WriteJSON(command)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stream) Unsubscribe(streams []string) error {
	command := newCommandUnSubscribe(s.requestId, streams)
	s.requestId = s.requestId + 1

	if s.ws.logger.CanDebug() {
		s.ws.logger.Debug(fmt.Sprintf("websocket[%d] unsubscribe stream", s.ws.id))
	}

	err := s.ws.WriteJSON(command)
	if err != nil {
		return err
	}

	for _, stream := range streams {
		if _, exists := s.streams[stream]; exists {
			delete(s.streams, stream)
		}
	}

	return nil
}

func (s *Stream) ListSubscription() {
	command := newCommandListSubscription(s.requestId)
	s.requestId = s.requestId + 1
	err := s.ws.WriteJSON(command)
	if err != nil {
		s.ws.logger.Error("list subscription error")
		if s.ws.logger.CanTrace() {
			s.ws.logger.Error(err.Error())
		}
	}
}

func (s *Stream) readMessage() {
	s.ws.SetPongHandler()
	s.ws.wg.Add(1)

	for {
		if !s.ws.IsNotDone() {
			s.ws.logger.Info(fmt.Sprintf("websocket[%d] stop read message", s.ws.id))
			s.ws.wg.Done()
			return
		}
		if s.ws.IsConnected() {
			_, message, err := s.ws.ReadMessage()
			if err == nil {
				streamData, err := parseStreamData(message)
				if err == nil {
					go s.messageHandler(streamData)
				} else {
					// on event subscribe, unsubscribe, listSubscription

				}
			}
		}
	}
}

func (s *Stream) messageHandler(streamData *StreamData) {
	if streamType, ok := s.streams[streamData.Stream]; ok {
		switch streamType {
		case AggregateTradeStreamType:
			if r, err := parseAggregateTradeStream(streamData.Data); err == nil {
				s.callAggTradeStreamHandler(streamData.Stream, r)
			}
		case TradeStreamType:
			if r, err := parseTradeStream(streamData.Data); err == nil {
				s.callTradeStreamHandler(streamData.Stream, r)
			}
		case KlineStreamType:
			if r, err := parseKlineStream(streamData.Data); err == nil {
				s.callKlineStreamHandler(streamData.Stream, r)
			}
		case IndividualMiniTickerStreamType:
			if r, err := parseIndividualMiniTickerStream(streamData.Data); err == nil {
				s.callIndividualMiniTickerStreamHandler(streamData.Stream, r)
			}
		case AllMarketMiniTickersStreamType:
			if r, err := parseAllMiniTickerStream(streamData.Data); err == nil {
				s.callAllMarketMiniTickerStreamHandler(streamData.Stream, r)
			}
		case IndividualTickerStreamType:
			if r, err := parseIndividualTickerStream(streamData.Data); err == nil {
				s.callIndividualTickerStreamHandler(streamData.Stream, r)
			}
		case AllMarketTickersStreamType:
			if r, err := parseAllMarketTickersStreamHandler(streamData.Data); err == nil {
				s.callAllMarketTickersStreamHandler(streamData.Stream, r)
			}
		case IndividualBookTickerStreamType:
			if r, err := parseIndividualBookTickerStream(streamData.Data); err == nil {
				s.callIndividualBookTickerStreamHandler(streamData.Stream, r)
			}
		case AllBookTickersStreamType:
			if r, err := parseIndividualBookTickerStream(streamData.Data); err == nil {
				s.callAllBookTickerStreamHandler(streamData.Stream, r)
			}
		case PartialBookDepthStreamType, PartialBookDepth100msStreamType:
			if r, err := parsePartialBookDepthStream(streamData.Data); err == nil {
				s.callPartialBookDepthStreamHandler(streamData.Stream, r)
			}
		case DiffDepthStreamType, DiffDepth100msStreamType:
			if r, err := parseDiffDepthStream(streamData.Data); err == nil {
				s.callDiffDepthStreamHandler(streamData.Stream, r)
			}
		}
	}
}
