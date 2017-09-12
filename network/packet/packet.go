package network

type Packet interface {
	Init(p interface{}) bool
	Pack(data []byte) bool
	Unpack(data []byte) bool
}
