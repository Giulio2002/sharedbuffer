package fsm

import "github.com/Giulio2002/sharedbuffer/types"

type intervalTreeFreeSpaceManager struct {
	tree *types.AVLTree
}

func NewIntervalTreeFreeSpaceManager() FreeSpaceManager {
	return &intervalTreeFreeSpaceManager{
		tree: &types.AVLTree{},
	}
}

func (t *intervalTreeFreeSpaceManager) Dirty(size int) (offset int, c cancelFunc) {
	offset = 0
	t.tree.InOrderTraversal(func(n *types.Node) bool {
		if n.Key-offset >= size {
			return false
		}

		offset = n.Key + n.Value
		return true
	})
	t.tree.Insert(offset, size)
	c = func() {
		t.tree.Delete(offset)
	}
	return
}
