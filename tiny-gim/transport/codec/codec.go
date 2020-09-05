package codec

import (
	"github.com/vmihailenco/msgpack"
)

type SerializeType byte

type Codec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte, value interface{}) error
}

type MessagePackCodec struct{}

func (c MessagePackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c MessagePackCodec) Decode(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
