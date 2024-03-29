package sharedbuffer

import (
	"encoding/binary"
	"sync"
	"testing"

	"github.com/Giulio2002/sharedbuffer/fsm"
	"github.com/Giulio2002/sharedbuffer/management"
	"github.com/stretchr/testify/assert"
)

func TestMakeSimpleBitmap(t *testing.T) {
	p := NewConcurrentSharedBuffer(fsm.NewBitmapFreeSpaceManager(), management.NewMemoryBuffer())
	buf, cancelfunc := p.Make(100)
	defer cancelfunc()
	assert.Equal(t, len(buf), 100)
}

func TestMakeConcurrentBitmap(t *testing.T) {
	p := NewConcurrentSharedBuffer(fsm.NewBitmapFreeSpaceManager(), management.NewMemoryBuffer())
	base := uint32(0xffffff)
	var wg sync.WaitGroup

	for i := uint32(1); i <= 200_000; i++ {
		wg.Add(1)

		i := i

		go func(base, i uint32) {
			defer wg.Done()
			n := base + i
			buf, cancelfunc := p.Make(4)

			binary.LittleEndian.PutUint32(buf, uint32(n))
			num := binary.LittleEndian.Uint32(buf)
			assert.Equal(t, num, n)
			cancelfunc()
		}(base, i)
	}

	wg.Wait()
}
