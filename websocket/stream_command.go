package websocket

type StreamCommand struct {
	Method string   `json:"method,omitempty"`
	Params []string `json:"params,omitempty"`
	Id     uint64   `json:"id,omitempty"`
}

func newCommandSubscribe(id uint64, streams []string) *StreamCommand {
	return &StreamCommand{
		Method: "SUBSCRIBE",
		Params: streams,
		Id:     id,
	}
}

func newCommandUnSubscribe(id uint64, streams []string) *StreamCommand {
	return &StreamCommand{
		Method: "UNSUBSCRIBE",
		Params: streams,
		Id:     id,
	}
}

func newCommandListSubscription(id uint64) *StreamCommand {
	return &StreamCommand{
		Method: "LIST_SUBSCRIPTIONS",
		Id:     id,
	}
}
