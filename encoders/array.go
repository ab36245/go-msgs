package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewArrayEncoder(mp *msgpack.Encoder, length int) *ArrayEncoder {
	mp.PutArrayLength(length)
	return &ArrayEncoder{
		mp: mp,
	}
}

type ArrayEncoder struct {
	mp *msgpack.Encoder
}

func (e *ArrayEncoder) PutArray(length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *ArrayEncoder) PutDate(value time.Time) {
	encodeDate(e.mp, value)
}

func (e *ArrayEncoder) PutInt(value int) {
	encodeInt(e.mp, value)
}

func (e *ArrayEncoder) PutMap(length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *ArrayEncoder) PutObject(handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *ArrayEncoder) PutRef(value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *ArrayEncoder) PutString(value string) {
	encodeString(e.mp, value)
}
