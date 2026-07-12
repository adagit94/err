package err

import (
	"fmt"
)

const (
	MethodNotRegisteredCode = iota
	RouteNotRegisteredCode
	HandlerNotFoundCode
	ChannelNotClosedCode
	WsReadMessageCode
	WsWriteMessageCode
	WsWritePingMessageCode
	WsInvalidMessageCode
	WsNoPongCode
)

type Err struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func (err *Err) Error() string {
	return fmt.Sprintf("[%d]: %s", err.Code, err.Message)
}
