package websocket

// guideline by
// https://gist.github.com/navono/d3742c4b0f26f68f1a48d86cf4556726

import (
	"errors"
	"fmt"
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	ErrNotConnected = errors.New("websocket not connected")
)

type Websocket struct {
	id  int64
	url string
	// default to 2 seconds
	reconnectIntervalMin time.Duration
	// default to 30 seconds
	reconnectIntervalMax time.Duration
	// interval, default to 1.5
	reconnectIntervalFactor float64
	// default to 2 seconds
	HandshakeTimeout time.Duration
	PingDuration     time.Duration
	PongDuration     time.Duration

	conn        *websocket.Conn
	dialer      *websocket.Dialer
	mu          sync.Mutex
	logger      *lib.BinanceLogger
	isConnected bool
	isDone      bool
	wg          sync.WaitGroup

	OnConnect func(ws *Websocket)
}

func (ws *Websocket) WriteJSON(v interface{}) error {
	err := ErrNotConnected
	if ws.IsNotDone() && ws.IsConnected() {
		err = ws.conn.WriteJSON(v)
		if err != nil {
			if ws.logger.CanDebug() {
				ws.logger.Info("write message error, try reconnect")
				if ws.logger.CanTrace() {
					ws.logger.Error(err.Error())
				}
			}
			ws.closeAndReconnect()
		}
	}

	return err
}

func (ws *Websocket) WriteMessage(messageType int, data []byte) error {
	err := ErrNotConnected
	if ws.IsNotDone() && ws.IsConnected() {
		err = ws.conn.WriteMessage(messageType, data)
		if err != nil {
			if ws.logger.CanDebug() {
				ws.logger.Info(fmt.Sprintf("websocket[%d]: write message error, try reconnect", ws.id))
				if ws.logger.CanTrace() {
					ws.logger.Error(err.Error())
				}
			}
			ws.closeAndReconnect()
		}
	}

	return err
}

func (ws *Websocket) ReadMessage() (messageType int, message []byte, err error) {
	err = ErrNotConnected
	if ws.IsNotDone() && ws.IsConnected() {
		messageType, message, err = ws.conn.ReadMessage()
		if err != nil {
			if ws.logger.CanDebug() {
				ws.logger.Debug(fmt.Sprintf("websocket[%d] read message error, try reconnect", ws.id))
				if ws.logger.CanTrace() {
					ws.logger.Error(err.Error())
				}
			}
			ws.closeAndReconnect()
		}
	}

	return
}

func (ws *Websocket) Dial(urlStr string) error {
	ws.url = urlStr
	ws.isConnected = false
	ws.setDefaults()

	ws.dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: ws.HandshakeTimeout,
	}

	hs := ws.HandshakeTimeout
	ws.Connect()

	// wait on first attempt
	time.Sleep(hs)

	return nil
}

func (ws *Websocket) Connect() {
	b := &backoff.Backoff{
		Min:    ws.reconnectIntervalMin,
		Max:    ws.reconnectIntervalMax,
		Factor: ws.reconnectIntervalFactor,
		Jitter: true,
	}

	// seed rand for backoff
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		nextInterval := b.Duration()

		wsConn, _, err := ws.dialer.Dial(ws.url, nil)

		ws.mu.Lock()
		ws.conn = wsConn
		ws.isConnected = err == nil
		ws.mu.Unlock()

		if err == nil {
			go ws.SetPingHandler()
			ws.logger.Info(fmt.Sprintf("websocket[%d] connection was successfully established with %s", ws.id, ws.url))
			if ws.OnConnect != nil {
				ws.OnConnect(ws)
			}
			ws.wg.Add(1)
			return
		} else {
			ws.logger.Error(fmt.Sprintf("websocket[%d] can't connect to %s, will try again in %v", ws.id, ws.url, nextInterval))
		}

		time.Sleep(nextInterval)
	}
}

func (ws *Websocket) SetPingHandler() {
	ticker := time.NewTicker(ws.PingDuration)
	defer func() {
		ticker.Stop()
	}()

	ws.wg.Add(1)
	for {
		if ws.isDone {
			ws.logger.Info(fmt.Sprintf("websocket[%d] stop ping handler", ws.id))
			ws.wg.Done()
			return
		}
		select {
		case <-ticker.C:
			_ = ws.conn.SetWriteDeadline(time.Now().Add(time.Second))
			if err := ws.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				ws.logger.Error(err.Error())
			}
		default:
			time.Sleep(time.Second * 5)
		}
	}
}

func (ws *Websocket) SetPongHandler() {
	_ = ws.conn.SetReadDeadline(time.Now().Add(ws.PongDuration))
	ws.conn.SetPongHandler(func(appData string) error {
		err := ws.conn.SetReadDeadline(time.Now().Add(ws.PongDuration))
		if err != nil {
			return err
		}
		return nil
	})
}

func (ws *Websocket) Close() {
	ws.mu.Lock()
	if ws.conn != nil {
		err := ws.conn.Close()
		if err == nil && ws.isConnected {
			ws.logger.Info(fmt.Sprintf("websocket[%d] disconnection was successfully", ws.id))
		}
		ws.conn = nil
	}
	ws.isConnected = false
	ws.wg.Done()
	ws.mu.Unlock()
}

func (ws *Websocket) closeAndReconnect() {
	if ws.IsNotDone() {
		ws.Close()
		ws.Connect()
	}
}

func (ws *Websocket) Shutdown() {
	ws.isDone = true
	ws.logger.Info(fmt.Sprintf("websocket[%d] is shutting down...", ws.id))
	ws.Close()

	ws.wg.Wait()
	ws.logger.Info(fmt.Sprintf("websocket[%d] is shutdown", ws.id))
}

func (ws *Websocket) IsConnected() bool {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.isConnected
}

func (ws *Websocket) IsNotDone() bool {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return !ws.isDone
}

func (ws *Websocket) setDefaults() {
	if ws.reconnectIntervalMin == 0 {
		ws.reconnectIntervalMin = 2 * time.Second
	}

	if ws.reconnectIntervalMax == 0 {
		ws.reconnectIntervalMax = 30 * time.Second
	}

	if ws.reconnectIntervalFactor == 0 {
		ws.reconnectIntervalFactor = 1.5
	}

	if ws.HandshakeTimeout == 0 {
		ws.HandshakeTimeout = 2 * time.Second
	}
}
