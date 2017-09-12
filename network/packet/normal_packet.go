package network

import (
	"container/list"
)

type NormalPacket struct {
}

func (this *NormalPacket) Init() bool {
	return true
}

func (this *NormalPacket) Unpack(data []byte) bool {
	return true
}

func (this *NormalPacket) Pack(data []byte) bool {
	return true
}
