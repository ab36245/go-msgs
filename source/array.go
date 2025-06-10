package source

import (
	"time"

	"github.com/ab36245/go-codec"
	"github.com/ab36245/go-msgpack"
)

func Array(mp *msgpack.Decoder) *arraySource {
	length := mp.GetArrayLength()
	return &arraySource{
		mp:     mp,
		length: length,
	}
}

type arraySource struct {
	mp     *msgpack.Decoder
	length int
}

func (s *arraySource) GetArray() (codec.ArraySource, error) {
	return Array(s.mp), nil
}

func (s *arraySource) GetDate() (time.Time, error) {
	return s.mp.GetTime(), nil
}

func (s *arraySource) GetInt() (int, error) {
	return int(s.mp.GetInt()), nil
}

func (s *arraySource) GetObject() (codec.ObjectSource, error) {
	return Object(s.mp), nil
}

func (s *arraySource) GetString() (string, error) {
	return s.mp.GetString(), nil
}

func (s *arraySource) Length() int {
	return s.length
}
