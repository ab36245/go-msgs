package msgs

import (
	"github.com/ab36245/go-msgpack"

	"github.com/ab36245/go-msgs/decoders"
)

func Decoder(b []byte) *decoder {
	mp := msgpack.NewDecoder(b)
	id := int(mp.GetInt())
	return &decoder{
		ObjectDecoder: decoders.NewObjectDecoder(mp),
		id:            id,
		mp:            mp,
	}
}

type decoder struct {
	*decoders.ObjectDecoder
	id int
	mp *msgpack.Decoder
}

func (d *decoder) Id() int {
	return d.id
}
