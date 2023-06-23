package types

type ByteList struct {
	u []byte

	l, c int
}

func NewByteList() *ByteList {
	return &ByteList{}
}
