package msgs

import (
	"time"

	"github.com/ab36245/go-model"
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

func (d *objectDecoder) GetArray(name string) (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *objectDecoder) GetDate(name string) (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *objectDecoder) GetInt(name string) (int, error) {
	return decodeInt(d.mp)
}

func (d *objectDecoder) GetMap(name string) (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *objectDecoder) GetObject(name string) (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *objectDecoder) GetRef(name string) (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *objectDecoder) GetString(name string) (string, error) {
	return decodeString(d.mp)
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

func (d *arrayDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *arrayDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *arrayDecoder) GetInt() (int, error) {
	return decodeInt(d.mp)
}

func (d *arrayDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *arrayDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *arrayDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *arrayDecoder) GetString() (string, error) {
	return decodeString(d.mp)
}

func (d *arrayDecoder) Length() int {
	return d.length
}

func newMapDecoder(mp *msgpack.Decoder) *mapDecoder {
	return &mapDecoder{
		mp:     mp,
		length: mp.GetMapLength(),
	}
}

type mapDecoder struct {
	mp     *msgpack.Decoder
	length int
}

func (d *mapDecoder) GetArray() (model.ArrayDecoder, error) {
	return decodeArray(d.mp)
}

func (d *mapDecoder) GetDate() (time.Time, error) {
	return decodeDate(d.mp)
}

func (d *mapDecoder) GetInt() (int, error) {
	return decodeInt(d.mp)
}

func (d *mapDecoder) GetKey() (string, error) {
	return decodeString(d.mp)
}

func (d *mapDecoder) GetMap() (model.MapDecoder, error) {
	return decodeMap(d.mp)
}

func (d *mapDecoder) GetObject() (model.ObjectDecoder, error) {
	return decodeObject(d.mp)
}

func (d *mapDecoder) GetRef() (model.Ref, error) {
	return decodeRef(d.mp)
}

func (d *mapDecoder) GetString() (string, error) {
	return decodeString(d.mp)
}

func (d *mapDecoder) Length() int {
	return d.length
}

func decodeArray(mp *msgpack.Decoder) (*arrayDecoder, error) {
	return newArrayDecoder(mp), nil
}

func decodeDate(mp *msgpack.Decoder) (time.Time, error) {
	return mp.GetTime(), nil
}

func decodeInt(mp *msgpack.Decoder) (int, error) {
	return int(mp.GetInt()), nil
}

func decodeMap(mp *msgpack.Decoder) (*mapDecoder, error) {
	return newMapDecoder(mp), nil
}

func decodeObject(mp *msgpack.Decoder) (*objectDecoder, error) {
	return newObjectDecoder(mp), nil
}

func decodeRef(mp *msgpack.Decoder) (model.Ref, error) {
	return model.Ref(mp.GetString()), nil
}

func decodeString(mp *msgpack.Decoder) (string, error) {
	return mp.GetString(), nil
}
