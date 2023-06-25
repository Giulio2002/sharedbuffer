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

func (t *intervalTreeFreeSpaceManager) Dirty(size int) (offset int, c cancelFunc) {
	offset = t.tree.FindFreeInterval(size).Low
	t.tree.Insert(types.Interval{offset, offset + size})
	c = func() {
		t.tree.Delete(types.Interval{offset, offset + size})
	}
	return
}
