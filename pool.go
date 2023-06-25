package sharedbuffer

import (
	"sync"

	"github.com/Giulio2002/sharedbuffer/fsm"
	"github.com/Giulio2002/sharedbuffer/management"
)

type freeFunc func()

type ConcurrentSharedBuffer struct {
	fsm    fsm.FreeSpaceManager
	buffer management.Buffer

	mu sync.Mutex
}

func NewConcurrentSharedBuffer(fsm fsm.FreeSpaceManager, buffer management.Buffer) *ConcurrentSharedBuffer {
	return &ConcurrentSharedBuffer{
		fsm:    fsm,
		buffer: buffer,
	}
}

func (s *ConcurrentSharedBuffer) Make(n int) ([]byte, freeFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	offset, cancelFn := s.fsm.Dirty(n)
	return s.buffer.Get(offset, n), func() {
		s.mu.Lock()
		cancelFn()
		defer s.mu.Unlock()
	}
}

type SharedBuffer struct {
	fsm    fsm.FreeSpaceManager
	buffer management.Buffer
}

func NewSharedBuffer(fsm fsm.FreeSpaceManager, buffer management.Buffer) *SharedBuffer {
	return &SharedBuffer{
		fsm:    fsm,
		buffer: buffer,
	}
}

func (s *SharedBuffer) Make(n int) ([]byte, freeFunc) {
	offset, cancelFn := s.fsm.Dirty(n)
	return s.buffer.Get(offset, n), func() {
		cancelFn()
	}
}
