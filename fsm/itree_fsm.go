package fsm

import "github.com/Giulio2002/sharedpool/types"

type intervalTreeFreeSpaceManager struct {
	tree *types.IntervalTree
}

func NewIntervalTreeFreeSpaceManager() FreeSpaceManager {
	return &intervalTreeFreeSpaceManager{
		tree: &types.IntervalTree{},
	}
}

func (t *intervalTreeFreeSpaceManager) MarkBusy(startPos, count int) {
	t.tree.Insert(types.Interval{startPos, startPos + count})
}

func (t *intervalTreeFreeSpaceManager) MarkFree(startPos, count int) {
	t.tree.Delete(types.Interval{startPos, startPos + count})
}

func (t *intervalTreeFreeSpaceManager) FirstFreeIndex(count int) int {
	interval := t.tree.FindFreeInterval(count)

	return interval.Low
}
