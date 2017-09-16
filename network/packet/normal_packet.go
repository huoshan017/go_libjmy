package network

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"

	"github.com/huoshan017/go_libjmy/util"
)

const (
	MAX_PACKET_LENGTH           = 65536
	NORMAL_PACKET_HEADER_LENGTH = 2
)

type NormalPacket struct {
	header_read uint8
	header      []byte
	data_len    uint16
	data_read   uint16
	data        []byte
	max_len     uint16
}

func (this *NormalPacket) Init(p interface{}) bool {
	this.max_len = p.(uint16)
	this.header = make([]byte, NORMAL_PACKET_HEADER_LENGTH)
	return true
}

func (this *NormalPacket) unpack(data []byte) (error, []byte) {
	offset := uint16(0)
	data_len := uint16(len(data))
	if data_len == 0 {
		return nil, nil
	}

	// header read incomplete
	if this.header_read < NORMAL_PACKET_HEADER_LENGTH {
		// copy the data to header
		copy(this.header[this.header_read:], data)
		offset = uint16(NORMAL_PACKET_HEADER_LENGTH) - uint16(this.header_read)
		// not enough header
		if data_len < offset {
			this.header_read += uint8(data_len)
			fmt.Printf("not enough header data to read")
			return nil, nil
		}
		// header read complete
		this.header_read = NORMAL_PACKET_HEADER_LENGTH
		// get data len
		this.data_len = util.UNPACK_DATA_TO_UINT16(this.header)
		if this.data_len == 0 { // packet length cant be zero
			err := errors.New("packet length is zero")
			return err, nil
		} else if this.data_len > this.max_len { // to large packet not support
			err_str := "too large packet data: " + strconv.Itoa(int(this.data_len))
			err := errors.New(err_str)
			return err, nil
		}
		// make buffer to cache data
		this.data = make([]byte, this.data_len)
	}

	// to length can read
	can_read := data_len - offset
	if can_read <= 0 {
		return nil, nil
	}

	copy(this.data[this.data_len:], data[offset:])

	// need read
	need_read := this.data_len - this.data_read
	if can_read < need_read {
		this.data_read += can_read
		return nil, nil
	}

	this.header_read = 0
	this.data_read = 0
	this.data_len = 0

	return nil, this.data
}

func (this *NormalPacket) Unpack(data []byte, packet_list *list.List) error {
	offset := uint16(0)
	var err error
	var p []byte
	for {
		err, p = this.unpack(data[:offset])
		if err != nil || (err == nil && p == nil) {
			break
		}
		packet_list.PushBack(p)
		offset += uint16(len(p))
	}
	return err
}

func (this *NormalPacket) Pack(data []byte) []byte {
	data_len := uint16(len(data))
	if data == nil || data_len == 0 {
		return nil
	}
	d := make([]byte, data_len+NORMAL_PACKET_HEADER_LENGTH)
	util.PACK_UINT16_TO_DATA(data_len, d)
	copy(d[:NORMAL_PACKET_HEADER_LENGTH], data)
	return d
}
