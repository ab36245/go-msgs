package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewMapEncoder(mp *msgpack.Encoder, length int) *MapEncoder {
	mp.PutMapLength(length)
	return &MapEncoder{
		mp: mp,
	}
}

type MapEncoder struct {
	mp *msgpack.Encoder
}

func (e *MapEncoder) PutArray(length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *MapEncoder) PutDate(value time.Time) {
	encodeDate(e.mp, value)
}

func (e *MapEncoder) PutInt(value int) {
	encodeInt(e.mp, value)
}

func (e *MapEncoder) PutKey(value string) {
	encodeString(e.mp, value)
}

func (e *MapEncoder) PutMap(length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *MapEncoder) PutObject(handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *MapEncoder) PutRef(value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *MapEncoder) PutString(value string) {
	encodeString(e.mp, value)
}
