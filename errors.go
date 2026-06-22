package err

import (
	"fmt"
)

const (
	MethodNotRegisteredCode = iota
	RouteNotRegisteredCode
	HandlerNotFoundCode
	ChannelNotClosed
)

type Err struct {
	Code    int
	Message string
}

func (err *Err) Format() string {
	return fmt.Sprintf("[%d]: %s", err.Code, err.Message)
}
