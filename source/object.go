package source

import (
	"time"

	"github.com/ab36245/go-codec"
	"github.com/ab36245/go-msgpack"
)

func Object(mp *msgpack.Decoder) *objectSource {
	return &objectSource{
		mp: mp,
	}
}

type objectSource struct {
	mp *msgpack.Decoder
}

func (s *objectSource) GetArray(name string) (codec.ArraySource, error) {
	return Array(s.mp), nil
}

func (s *objectSource) GetDate(name string) (time.Time, error) {
	return s.mp.GetTime(), nil
}

func (s *objectSource) GetInt(name string) (int, error) {
	return int(s.mp.GetInt()), nil
}

func (s *objectSource) GetObject(name string) (codec.ObjectSource, error) {
	return Object(s.mp), nil
}

func (s *objectSource) GetString(name string) (string, error) {
	return s.mp.GetString(), nil
}
