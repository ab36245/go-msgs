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
		mp:            mp,
		id:            id,
	}
}

type decoder struct {
	*decoders.ObjectDecoder
	mp *msgpack.Decoder
	id int
}

func (d *decoder) Id() int {
	return d.id
}
