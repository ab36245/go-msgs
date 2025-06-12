package msgs

import (
	"time"

	"github.com/ab36245/go-defs"
	"github.com/ab36245/go-msgpack"
)

func Decoder(b []byte) *decoder {
	mp := msgpack.NewDecoder(b)
	id := int(mp.GetInt())
	return &decoder{
		objectDecoder: *newObjectDecoder(mp),
		id:            id,
	}
}

type decoder struct {
	objectDecoder
	id int
}

func (d *decoder) Id() int {
	return d.id
}

func newObjectDecoder(mp *msgpack.Decoder) *objectDecoder {
	return &objectDecoder{
		mp: mp,
	}
}

type objectDecoder struct {
	mp *msgpack.Decoder
}

func (s *objectDecoder) GetArray(name string) (defs.ArrayDecoder, error) {
	return decodeArray(s.mp)
}

func (s *objectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(s.mp)
}

func (s *objectDecoder) GetInt(name string) (int, error) {
	return decodeInt(s.mp)
}

func (s *objectDecoder) GetObject(name string) (defs.ObjectDecoder, error) {
	return decodeObject(s.mp)
}

func (s *objectDecoder) GetString(name string) (string, error) {
	return decodeString(s.mp)
}

func newArrayDecoder(mp *msgpack.Decoder) *arrayDecoder {
	return &arrayDecoder{
		mp:     mp,
		length: mp.GetArrayLength(),
	}
}

type arrayDecoder struct {
	mp     *msgpack.Decoder
	length int
}

func (s *arrayDecoder) GetArray() (defs.ArrayDecoder, error) {
	return decodeArray(s.mp)
}

func (s *arrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(s.mp)
}

func (s *arrayDecoder) GetInt() (int, error) {
	return decodeInt(s.mp)
}

func (s *arrayDecoder) GetObject() (defs.ObjectDecoder, error) {
	return decodeObject(s.mp)
}

func (s *arrayDecoder) GetString() (string, error) {
	return decodeString(s.mp)
}

func (s *arrayDecoder) Length() int {
	return s.length
}

func decodeDate(mp *msgpack.Decoder) (time.Time, error) {
	return mp.GetTime(), nil
}

func decodeArray(mp *msgpack.Decoder) (*arrayDecoder, error) {
	return newArrayDecoder(mp), nil
}

func decodeInt(mp *msgpack.Decoder) (int, error) {
	return int(mp.GetInt()), nil
}

func decodeObject(mp *msgpack.Decoder) (*objectDecoder, error) {
	return newObjectDecoder(mp), nil
}

func decodeString(mp *msgpack.Decoder) (string, error) {
	return mp.GetString(), nil
}
