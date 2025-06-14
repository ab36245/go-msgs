package decoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewObjectDecoder(mp *msgpack.Decoder) *ObjectDecoder {
	return &ObjectDecoder{
		mp: mp,
	}
}

type ObjectDecoder struct {
	mp *msgpack.Decoder
}

func (d *ObjectDecoder) GetArray(name string) (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *ObjectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *ObjectDecoder) GetInt(name string) (int, error) {
	return decodeInt(d.mp)
}

func (d *ObjectDecoder) GetMap(name string) (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *ObjectDecoder) GetObject(name string) (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *ObjectDecoder) GetRef(name string) (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *ObjectDecoder) GetString(name string) (string, error) {
	return decodeString(d.mp)
}
