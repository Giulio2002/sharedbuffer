package sharedbuffer

import (
	"sync"

	"github.com/Giulio2002/sharedbuffer/fsm"
	"github.com/Giulio2002/sharedbuffer/management"
)

var globalBuffer SharedBuffer

func init() {
	globalBuffer = NewConcurrentSharedBuffer(fsm.NewIntervalTreeFreeSpaceManager(), management.NewMemoryBuffer())
}

func SetGlobalBuffer(b SharedBuffer) {
	globalBuffer = b
}

func Make(size int) ([]byte, freeFunc) {
	return globalBuffer.Make(size)
}

type SharedBuffer interface {
	Make(n int) ([]byte, freeFunc)
}

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

type SimpleSharedBuffer struct {
	fsm    fsm.FreeSpaceManager
	buffer management.Buffer
}

func NewSimpleSharedBuffer(fsm fsm.FreeSpaceManager, buffer management.Buffer) *SimpleSharedBuffer {
	return &SimpleSharedBuffer{
		fsm:    fsm,
		buffer: buffer,
	}
}

func (s *SimpleSharedBuffer) Make(n int) ([]byte, freeFunc) {
	offset, cancelFn := s.fsm.Dirty(n)
	return s.buffer.Get(offset, n), freeFunc(cancelFn)
}
