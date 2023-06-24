package management

import "fmt"

const bufferExpansion = 1.8 // increase buffer by 20% of the max asked.

type memoryBuffer struct {
	u []byte
}

func NewMemoryBuffer() *memoryBuffer {
	return &memoryBuffer{}
}

func (m *memoryBuffer) Get(startPos, n int) []byte {
	if startPos < 0 {
		panic("negative index")
	}
	if startPos+n >= len(m.u) {
		newSize := startPos + n
		tmp := m.u
		m.u = make([]byte, int(float64(newSize)*bufferExpansion))
		copy(m.u[:], tmp)
	}
	for i := startPos; i < startPos+n; i++ {
		m.u[i] = 0x00
	}
	fmt.Println("taken", startPos)
	return m.u[startPos : startPos+n]
}
