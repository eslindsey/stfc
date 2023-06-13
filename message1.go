package stfc

import (
	"fmt"
)

type Message1Type uint32

const (
	Message1None Message1Type = 0
	Message1Json              = 42
)

func (t Message1Type) String() string {
	switch t {
	case Message1Json: return "JSON"
	case Message1None: return "None"
	default:           return fmt.Sprintf("Message1Type(%d)", t)
	}
}

