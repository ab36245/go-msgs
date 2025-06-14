package decoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func NewArrayDecoder(mp *msgpack.Decoder) *ArrayDecoder {
	return &ArrayDecoder{
		mp:     mp,
		length: mp.GetArrayLength(),
	}
}

type ArrayDecoder struct {
	mp     *msgpack.Decoder
	length int
}

func (d *ArrayDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *ArrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *ArrayDecoder) GetInt() (int, error) {
	return decodeInt(d.mp)
}

func (d *ArrayDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *ArrayDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *ArrayDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *ArrayDecoder) GetString() (string, error) {
	return decodeString(d.mp)
}

func (d *ArrayDecoder) Length() int {
	return d.length
}
