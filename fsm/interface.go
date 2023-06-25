package fsm

type cancelFunc func()

// The free space manager just keeps track of the size of the free space.
type FreeSpaceManager interface {
	Dirty(size int) (offset int, cancel cancelFunc)
}
