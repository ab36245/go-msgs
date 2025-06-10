package target

import (
	"time"

	"github.com/ab36245/go-codec"
	"github.com/ab36245/go-msgpack"
)

func Array(mp *msgpack.Encoder, length int) *arrayTarget {
	mp.PutArrayLength(length)
	return &arrayTarget{
		mp:     mp,
		length: length,
	}
}

type arrayTarget struct {
	mp     *msgpack.Encoder
	length int
}

func (t *arrayTarget) PutArray(length int, f func(codec.ArrayTarget)) {
	array := Array(t.mp, length)
	f(array)
}

func (t *arrayTarget) PutDate(value time.Time) {
	t.mp.PutTime(value)
}

func (t *arrayTarget) PutInt(value int) {
	t.mp.PutInt(int64(value))
}

func (t *arrayTarget) PutObject(f func(codec.ObjectTarget)) {
	object := Object(t.mp)
	f(object)
}

func (t *arrayTarget) PutString(value string) {
	t.mp.PutString(value)
}
