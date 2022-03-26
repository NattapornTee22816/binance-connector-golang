package websocket

type ErrStreamParameterRequired struct {
	error
	Message string
}

func (e *ErrStreamParameterRequired) Error() string {
	return e.Message
}
