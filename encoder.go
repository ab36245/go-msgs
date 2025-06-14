package msgs

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func Encoder() *encoder {
	mp := msgpack.NewEncoder()
	return &encoder{
		objectEncoder: *newObjectEncoder(mp),
	}
}

type encoder struct {
	objectEncoder
}

func (e *encoder) Bytes() []byte {
	return e.mp.Bytes()
}

func (e *encoder) Id(id int) {
	e.mp.PutInt(int64(id))
}

func newObjectEncoder(mp *msgpack.Encoder) *objectEncoder {
	return &objectEncoder{
		mp: mp,
	}
}

type objectEncoder struct {
	mp *msgpack.Encoder
}

func (e *objectEncoder) PutArray(name string, length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *objectEncoder) PutDate(name string, value time.Time) {
	encodeDate(e.mp, value)
}

func (e *objectEncoder) PutInt(name string, value int) {
	encodeInt(e.mp, value)
}

func (e *objectEncoder) PutMap(name string, length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *objectEncoder) PutObject(name string, handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *objectEncoder) PutRef(name string, value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *objectEncoder) PutString(name string, value string) {
	encodeString(e.mp, value)
}

func newArrayEncoder(mp *msgpack.Encoder, length int) *arrayEncoder {
	mp.PutArrayLength(length)
	return &arrayEncoder{
		mp: mp,
	}
}

type arrayEncoder struct {
	mp *msgpack.Encoder
}

func (e *arrayEncoder) PutArray(length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *arrayEncoder) PutDate(value time.Time) {
	encodeDate(e.mp, value)
}

func (e *arrayEncoder) PutInt(value int) {
	encodeInt(e.mp, value)
}

func (e *arrayEncoder) PutMap(length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *arrayEncoder) PutObject(handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *arrayEncoder) PutRef(value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *arrayEncoder) PutString(value string) {
	encodeString(e.mp, value)
}

func newMapEncoder(mp *msgpack.Encoder, length int) *mapEncoder {
	mp.PutMapLength(length)
	return &mapEncoder{
		mp: mp,
	}
}

type mapEncoder struct {
	mp *msgpack.Encoder
}

func (e *mapEncoder) PutArray(length int, handler func(model.ArrayEncoder)) {
	encodeArray(e.mp, length, handler)
}

func (e *mapEncoder) PutDate(value time.Time) {
	encodeDate(e.mp, value)
}

func (e *mapEncoder) PutInt(value int) {
	encodeInt(e.mp, value)
}

func (e *mapEncoder) PutKey(value string) {
	encodeString(e.mp, value)
}

func (e *mapEncoder) PutMap(length int, handler func(model.MapEncoder)) {
	encodeMap(e.mp, length, handler)
}

func (e *mapEncoder) PutObject(handler func(model.ObjectEncoder)) {
	encodeObject(e.mp, handler)
}

func (e *mapEncoder) PutRef(value model.Ref) {
	encodeRef(e.mp, value)
}

func (e *mapEncoder) PutString(value string) {
	encodeString(e.mp, value)
}

func encodeArray(mp *msgpack.Encoder, length int, handler func(model.ArrayEncoder)) {
	handler(newArrayEncoder(mp, length))
}

func encodeDate(mp *msgpack.Encoder, value time.Time) {
	mp.PutTime(value)
}

func encodeInt(mp *msgpack.Encoder, value int) {
	mp.PutInt(int64(value))
}

func encodeMap(mp *msgpack.Encoder, length int, handler func(model.MapEncoder)) {
	handler(newMapEncoder(mp, length))
}

func encodeObject(mp *msgpack.Encoder, handler func(model.ObjectEncoder)) {
	handler(newObjectEncoder(mp))
}

func encodeRef(mp *msgpack.Encoder, value model.Ref) {
	mp.PutString(string(value))
}

func encodeString(mp *msgpack.Encoder, value string) {
	mp.PutString(value)
}
