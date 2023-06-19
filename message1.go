package stfc

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Message1Type uint32

const (
	Message1None Message1Type = 0
	Message1Json              = 42
)

var (
	ErrTypeNotFound = errors.New("type not found")
)

func (t Message1Type) String() string {
	switch t {
	case Message1Json: return "JSON"
	case Message1None: return "None"
	default:           return fmt.Sprintf("Message1Type(%d)", t)
	}
}

func getMessage1(body []byte, t Message1Type) ([]byte, error) {
	var generic Generic
	if err := proto.Unmarshal(body, &generic); err != nil {
		return nil, err
	}
	for _, d := range generic.Payload {
		if Message1Type(d.Type) == t {
			return []byte(d.Data), nil
		}
	}
	return nil, ErrTypeNotFound
}

func getMessage1JSON(body []byte, dest interface{}) error {
	b, err := getMessage1(body, Message1Json)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, dest); err != nil {
		return err
	}
	return nil
}

