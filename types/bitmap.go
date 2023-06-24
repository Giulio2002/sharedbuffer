package types

type Bitmap struct {
	b []byte
}

const (
	BITS_PER_BYTE    = 8
	BITMAP_EXPANSION = 2.0 // Double bitmap every time the buffer length is exceeded
)

func NewBitmap() *Bitmap {
	return &Bitmap{}
}

func (b *Bitmap) Set(idx int, on bool) {
	byteIdx := idx / BITS_PER_BYTE
	bitIdx := idx % BITS_PER_BYTE
	if byteIdx >= len(b.b) {
		tmp := b.b
		b.b = make([]byte, int(float64(byteIdx+1)*BITMAP_EXPANSION))
		copy(b.b, tmp)
	}
	if on {
		b.b[byteIdx] = ((1 << bitIdx) | b.b[byteIdx])
		return
	}
	b.b[byteIdx] &= ^(1 << bitIdx)
}

func (b *Bitmap) Get(idx int) bool {
	byteIdx := idx / BITS_PER_BYTE
	bitIdx := idx % BITS_PER_BYTE
	return byteIdx < len(b.b) && b.b[byteIdx]&(1<<bitIdx) > 0
}
