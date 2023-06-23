package fsm

import (
	"sort"

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
	nasties []dirtySegmentMetadata
	// store index of previously allocated dirty segments.
	cleanies *types.ConcurrentStack[int]
}

func NewCountingFreeSpaceManager() FreeSpaceManager {
	return &countingFreeSpaceManager{
		cleanies: types.NewConcurrentStack[int](),
	}
}

func (c *countingFreeSpaceManager) MarkBusy(startPos, n int) {
	defer sort.Slice(c.nasties, func(i, j int) bool {
		return c.nasties[i].index < c.nasties[j].index
	})
	// If we can use an idle entry from the stack let us do so.
	freeIndex, found := c.cleanies.Pop()
	// If it was precached just use the stack entry index.
	if found {
		c.nasties[freeIndex].index = startPos
		c.nasties[freeIndex].n = n
	}
	// we do not have any pre-allocated dirtySegmentMetadatas? noice, we just make one :D.
	c.nasties = append(c.nasties, dirtySegmentMetadata{
		index: startPos,
		n:     n,
	})
}

// MarkFree marks contiguous allocation as free, we can just use the start position and do simple linear search.
func (c *countingFreeSpaceManager) MarkFree(startPos, _ int) {
	defer sort.Slice(c.nasties, func(i, j int) bool {
		return c.nasties[i].index < c.nasties[j].index
	})
	i := sort.Search(len(c.nasties), func(i int) bool {
		return c.nasties[i].index == startPos
	})
	// Now it is clean :D.
	c.nasties[i].index = freeIndex
	c.cleanies.Add(i)
}

// Get an index which is free and accomadate for n bytes.
func (c *countingFreeSpaceManager) FirstFreeIndex(n int) int {
	// stands for begining of buffer, this is a first fit algo
	currentIndex := 0
	for i := range c.nasties {
		currentIndexEnd := currentIndex + n
		nextDirtyIndex := c.nasties[i].index
		if nextDirtyIndex < currentIndex || nextDirtyIndex > currentIndexEnd {
			break
		}
		currentIndexEnd = nextDirtyIndex + c.nasties[i].n
	}
	return currentIndex
}
