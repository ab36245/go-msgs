package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewObjectEncoder(mp *msgpack.Encoder) *ObjectEncoder {
	return &ObjectEncoder{
		mp: mp,
	}
}

type ObjectEncoder struct {
	mp *msgpack.Encoder
}

func (e *ObjectEncoder) PutArray(name string, length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *ObjectEncoder) PutDate(name string, value time.Time) {
	encodeDate(e.mp, value)
}

func (e *ObjectEncoder) PutInt(name string, value int) {
	encodeInt(e.mp, value)
}

func (e *ObjectEncoder) PutMap(name string, length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *ObjectEncoder) PutObject(name string, handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *ObjectEncoder) PutRef(name string, value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *ObjectEncoder) PutString(name string, value string) {
	encodeString(e.mp, value)
}
