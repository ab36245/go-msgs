package decoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewMapDecoder(mp *msgpack.Decoder) *MapDecoder {
	return &MapDecoder{
		mp:     mp,
		length: mp.GetMapLength(),
	}
}

type MapDecoder struct {
	mp     *msgpack.Decoder
	length int
}

func (d *MapDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *MapDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *MapDecoder) GetInt() (int, error) {
	return decodeInt(d.mp)
}

func (d *MapDecoder) GetKey() (string, error) {
	return decodeString(d.mp)
}

func (d *MapDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *MapDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *MapDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *MapDecoder) GetString() (string, error) {
	return decodeString(d.mp)
}

func (d *MapDecoder) Length() int {
	return d.length
}
