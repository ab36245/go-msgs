package encoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func encodeArray(mp *msgpack.Encoder, length int, handler func(model.ArrayEncoder)) {
	handler(NewArrayEncoder(mp, length))
}

func encodeDate(mp *msgpack.Encoder, value time.Time) {
	mp.PutTime(value)
}

func encodeInt(mp *msgpack.Encoder, value int) {
	mp.PutInt(int64(value))
}

func encodeMap(mp *msgpack.Encoder, length int, handler func(model.MapEncoder)) {
	handler(NewMapEncoder(mp, length))
}

func encodeObject(mp *msgpack.Encoder, handler func(model.ObjectEncoder)) {
	handler(NewObjectEncoder(mp))
}

func encodeRef(mp *msgpack.Encoder, value model.Ref) {
	mp.PutString(string(value))
}

func encodeString(mp *msgpack.Encoder, value string) {
	mp.PutString(value)
}
