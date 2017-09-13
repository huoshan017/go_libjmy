package network

import (
	"container/list"
)

type Packet interface {
	Init(p interface{}) bool
	Pack(data []byte, packet_list *list.List) error
	Unpack(data []byte) []byte
}
