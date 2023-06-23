package fsm

type FreeSpaceManager interface {
	MarkBusy(startPos, count int)
	MarkFree(startPos, count int)
	FirstFreeIndex(count int) int
}
