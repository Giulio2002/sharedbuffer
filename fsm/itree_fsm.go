package fsm

import (
	"github.com/google/btree"
)

type Item struct {
	offset int // Key of the Btree
	size   int // Value of the btree
}

func (i Item) Less(than btree.Item) bool {
	return i.offset < than.(Item).offset
}

type intervalTreeFreeSpaceManager struct {
	tree *btree.BTree
}

func NewIntervalTreeFreeSpaceManager() FreeSpaceManager {
	return &intervalTreeFreeSpaceManager{
		tree: btree.New(16),
	}
}

func (t *intervalTreeFreeSpaceManager) Dirty(size int) (offset int, c cancelFunc) {
	item := Item{
		size: size,
	}
	offset = 0
	t.tree.Ascend(func(it btree.Item) bool {
		obj := it.(Item)
		if obj.offset-offset >= size {
			return false
		}

		offset = obj.offset + obj.size
		return true
	})
	item.offset = offset
	t.tree.ReplaceOrInsert(item)
	c = func() {
		t.tree.Delete(item)
	}
	return
}
