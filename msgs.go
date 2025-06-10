package msgs

// import (
// 	"github.com/aivoicesystems/aivoice/common/msgpack"
// )

// func New(decoders Decoders) *Msgs {
// 	return &Msgs{
// 		decoders: decoders,
// 	}
// }

// type Msgs struct {
// 	decoders Decoders
// }

// func (m Msgs) Decode(bytes []byte) (Msg, error) {
// 	mp := msgpack.NewDecoder(bytes)
// 	id := mp.GetInt()
// 	return m.decoders(id, mp)
// }

// func (m Msgs) DecodeSocket(socket websocket.Socket) *stream.Output[Msg] {
// 	return m.DecodeStream(socket.StreamBinary())
// }

// func (m Msgs) DecodeStream(upstream *stream.Output[[]byte]) *stream.Output[Msg] {
// 	return stream.Add(
// 		upstream,
// 		func(input <-chan []byte, output func(Msg)) error {
// 			for {
// 				bytes, ok := <-input
// 				if len(bytes) != 0 {
// 					msg, err := m.Decode(bytes)
// 					if err != nil {
// 						return err
// 					}
// 					output(msg)
// 				}
// 				if !ok {
// 					return nil
// 				}
// 			}
// 		},
// 	)
// }

// func (m Msgs) Encode(msg Msg) []byte {
// 	mp := msgpack.NewEncoder()
// 	mp.PutInt(msg.MsgId())
// 	msg.MsgEncode(mp)
// 	return mp.Bytes()
// }

// func (m Msgs) EncodeSocket(
// 	upstream *stream.Output[Msg],
// 	socket websocket.Socket,
// ) *stream.Output[int] {
// 	return socket.SinkBinary(m.EncodeStream(upstream))
// }

// func (m Msgs) EncodeStream(upstream *stream.Output[Msg]) *stream.Output[[]byte] {
// 	return stream.Add(
// 		upstream,
// 		func(input <-chan Msg, output func([]byte)) error {
// 			for {
// 				msg, ok := <-input
// 				if msg != nil {
// 					bytes := m.Encode(msg)
// 					output(bytes)
// 				}
// 				if !ok {
// 					return nil
// 				}
// 			}
// 		},
// 	)
// }

// func (m Msgs) TransformSocket(
// 	socket websocket.Socket,
// 	handler func(<-chan Msg, func(Msg)) error,
// ) *stream.Output[int] {
// 	i := m.DecodeSocket(socket)
// 	t := stream.Add(i, handler)
// 	o := m.EncodeSocket(t, socket)
// 	return o
// }

// type Decoder func(*msgpack.Decoder) (Msg, error)

// type Decoders func(int64, *msgpack.Decoder) (Msg, error)
