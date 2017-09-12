package network

type FixedLengthPacket struct {
}

func (this *FixedLengthPacket) Init(p interface{}) bool {
	return true
}

func (this *FixedLengthPacket) Pack(data []byte) bool {
	return true
}

func (this *FixedLengthPacket) Unpack(data []byte) bool {
	return true
}
