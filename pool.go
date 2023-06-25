package sharedpool

import (
	"sync"

	"github.com/Giulio2002/sharedpool/fsm"
	"github.com/Giulio2002/sharedpool/management"
)

type freeFunc func()

type SharedPool struct {
	fsm    fsm.FreeSpaceManager
	buffer management.Buffer

	mu sync.Mutex
}

func NewSharedPool(fsm fsm.FreeSpaceManager, buffer management.Buffer) *SharedPool {
	return &SharedPool{
		fsm:    fsm,
		buffer: buffer,
	}
}

func (s *SharedPool) Make(n int) ([]byte, freeFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	offset, cancelFn := s.fsm.Dirty(n)
	return s.buffer.Get(offset, n), func() {
		s.mu.Lock()
		cancelFn()
		defer s.mu.Unlock()
	}
}
