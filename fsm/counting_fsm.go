package fsm

import (
	"github.com/Giulio2002/sharedpool/types"
)

// Magic number for segment metadata objects which were allocated in the past and are just fleeting
const freeIndex = -1

type dirtySegmentMetadata struct {
	index int
	n     int
}

type countingFreeSpaceManager struct {
	// store location of dirty segments.
	nasties *types.SortedList[dirtySegmentMetadata]
}

func NewCountingFreeSpaceManager() FreeSpaceManager {
	return &countingFreeSpaceManager{
		nasties: types.NewSortedList[dirtySegmentMetadata](func(a, b dirtySegmentMetadata) bool {
			return a.index < b.index
		}),
	}
}

func (c *countingFreeSpaceManager) MarkBusy(startPos, n int) {
	c.nasties.Add(dirtySegmentMetadata{
		index: startPos,
		n:     n,
	})
}

// MarkFree marks contiguous allocation as free, we can just use the start position and do simple linear search.
func (c *countingFreeSpaceManager) MarkFree(startPos, _ int) {
	_, idx, found := c.nasties.Search(func(a dirtySegmentMetadata) bool {
		return a.index >= startPos
	})
	if !found {
		panic("bad")
	}

	// Now it is clean :D.
	c.nasties.Remove(idx)
}

// Get an index which is free and accomadate for n bytes.
func (c *countingFreeSpaceManager) FirstFreeIndex(n int) int {
	// stands for begining of buffer, this is a first fit algo
	currentIndex := 0
	c.nasties.Range(func(val dirtySegmentMetadata, idx, l int) bool {
		if val.index == freeIndex {
			return true
		}
		currentIndexEnd := currentIndex + n
		nextDirtyIndex := val.index
		nextDirtyIndexEnd := val.index + val.n
		if currentIndexEnd < nextDirtyIndex || nextDirtyIndexEnd < currentIndex {
			return false
		}
		currentIndex = nextDirtyIndexEnd
		return true
	})

	return currentIndex
}
