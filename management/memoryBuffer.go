package management

const bufferExpansion = 1.2         // increase buffer by 20% of the maximum asked.
const initialBufferSize = 1_000_000 // 1 MB of initial buffer size is oke

type memoryBuffer struct {
	u []byte
}

func NewMemoryBuffer() *memoryBuffer {
	return &memoryBuffer{
		u: make([]byte, initialBufferSize),
	}
}

func (m *memoryBuffer) Get(startPos, n int) []byte {
	if startPos < 0 {
		panic("negative index")
	}
	if startPos+n >= len(m.u) {
		newSize := startPos + n
		tmp := m.u
		m.u = make([]byte, int(float64(newSize+1)*bufferExpansion))
		copy(m.u[:], tmp)
	}
	for i := startPos; i < startPos+n; i++ {
		m.u[i] = 0x00
	}
	return m.u[startPos : startPos+n]
}
