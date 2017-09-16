package network

import (
	"container/list"
)

type FixedLengthPacket struct {
	fixed_len uint16
	data_read uint16
	data      []byte
}

func (this *FixedLengthPacket) Init(p interface{}) bool {
	this.fixed_len = p.(uint16)
	return true
}

func (this *FixedLengthPacket) Pack(data []byte, packet_list *list.List) (err error) {
	data_len := uint16(len(data))
	if data == nil || len(data) == 0 {
		return
	}

	offset := uint16(0)
	for {
		can_read := data_len - offset
		if can_read <= 0 {
			break
		}
		need_read := uint16(len(this.data)) - this.data_read
		if can_read >= need_read {
			copy(this.data[this.data_read:], data[offset:offset+need_read])
			packet_list.PushBack(this.data)
			this.data_read = 0
			offset += need_read
		} else {
			copy(this.data[this.data_read:], data[offset:])
			offset += can_read
		}
	}
	return
}

func (this *FixedLengthPacket) Unpack(data []byte) []byte {
	return nil
}
