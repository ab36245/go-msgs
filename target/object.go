package target

import (
	"time"

	"github.com/ab36245/go-codec"
	"github.com/ab36245/go-msgpack"
)

func Object(mp *msgpack.Encoder) *objectTarget {
	return &objectTarget{
		mp: mp,
	}
}

type objectTarget struct {
	mp *msgpack.Encoder
}

func (t *objectTarget) PutArray(name string, length int, f func(codec.ArrayTarget)) {
	array := Array(t.mp, length)
	f(array)
}

func (t *objectTarget) PutDate(name string, value time.Time) {
	t.mp.PutTime(value)
}

func (t *objectTarget) PutInt(name string, value int) {
	t.mp.PutInt(int64(value))
}

func (t *objectTarget) PutObject(name string, f func(codec.ObjectTarget)) {
	object := Object(t.mp)
	f(object)
}

func (t *objectTarget) PutString(name string, value string) {
	t.mp.PutString(value)
}
