package msgs

import (
	"github.com/ab36245/go-msgpack"

	"github.com/ab36245/go-msgs/encoders"
)

func Encoder() *encoder {
	mp := msgpack.NewEncoder()
	return &encoder{
		ObjectEncoder: encoders.NewObjectEncoder(mp),
		mp:            mp,
	}
}

type encoder struct {
	*encoders.ObjectEncoder
	mp *msgpack.Encoder
}

func (e *encoder) Bytes() []byte {
	return e.mp.Bytes()
}

func (e *encoder) Id(id int) {
	e.mp.PutInt(int64(id))
}
