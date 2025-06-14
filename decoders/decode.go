package decoders

import (
	"time"

	"github.com/ab36245/go-model"
	"github.com/ab36245/go-msgpack"
)

func decodeArray(mp *msgpack.Decoder) (*ArrayDecoder, error) {
	return NewArrayDecoder(mp), nil
}

func decodeDate(mp *msgpack.Decoder) (time.Time, error) {
	return mp.GetTime(), nil
}

func decodeInt(mp *msgpack.Decoder) (int, error) {
	return int(mp.GetInt()), nil
}

func decodeMap(mp *msgpack.Decoder) (*MapDecoder, error) {
	return NewMapDecoder(mp), nil
}

func decodeObject(mp *msgpack.Decoder) (*ObjectDecoder, error) {
	return NewObjectDecoder(mp), nil
}

func decodeRef(mp *msgpack.Decoder) (model.Ref, error) {
	return model.Ref(mp.GetString()), nil
}

func decodeString(mp *msgpack.Decoder) (string, error) {
	return mp.GetString(), nil
}
