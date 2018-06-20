package mysql

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/golang/protobuf/proto"
)

func decode(data []byte, t string) (io.Reader, proto.Message, error) {
	mt := proto.MessageType(t)
	if mt == nil {
		return nil, nil, fmt.Errorf("unknown message type %q", t)
	}
	msg, ok := reflect.New(mt.Elem()).Interface().(proto.Message)
	if !ok {
		return nil, nil, errors.New("failed to make a new instance of the type of expected object")
	}
	return bytes.NewReader(data), msg, nil
}
