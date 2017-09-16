package util

func PACK_UINT16_TO_DATA(n uint16, data []byte) {
	data[0] = byte((n >> 8) & 0xff)
	data[1] = byte((n) & 0xff)
}

func UNPACK_DATA_TO_UINT16(data []byte) uint16 {
	n := uint16(uint16(data[0]<<8) & 0xff00)
	n += uint16(data[1] & 0xff)
	return n
}

func PACK_UINT32_TO_DATA(n uint32, data []byte) {
	data[0] = byte((n >> 24) & 0xff)
	data[1] = byte((n >> 16) & 0xff)
	data[2] = byte((n >> 8) & 0xff)
	data[3] = byte(n & 0xff)
}

func UNPACK_DATA_TO_UINT32(data []byte) uint32 {
	n := uint32(uint32(data[0]<<24) & 0xff000000)
	n += uint32(uint32(data[1]<<16) & 0xff0000)
	n += uint32(uint32(data[2]<<8) & 0xff00)
	n += uint32(data[3] & 0xff)
	return n
}

func PACK_UINT64_TO_DATA(n uint64, data []byte) {
	data[0] = byte((n >> 56) & 0xff)
	data[1] = byte((n >> 48) & 0xff)
	data[2] = byte((n >> 40) & 0xff)
	data[3] = byte((n >> 32) & 0xff)
	data[4] = byte((n >> 24) & 0xff)
	data[5] = byte((n >> 16) & 0xff)
	data[6] = byte((n >> 8) & 0xff)
	data[7] = byte((n) & 0xff)
}

func UNPACK_DATA_TO_UINT64(data []byte) uint64 {
	n := uint64(uint64(data[0]<<56) & 0xff00000000000000)
	n += uint64(uint64(data[1]<<48) & 0xff000000000000)
	n += uint64(uint64(data[2]<<40) & 0xff0000000000)
	n += uint64(uint64(data[3]<<32) & 0xff00000000)
	n += uint64(uint64(data[4]<<24) & 0xff000000)
	n += uint64(uint64(data[5]<<16) & 0xff0000)
	n += uint64(uint64(data[6]<<8) & 0xff00)
	n += uint64(data[7] & 0xff)
	return n
}
