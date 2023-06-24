package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTree(t *testing.T) {
	assert := assert.New(t)
	it := &IntervalTree{}

	// Test insertions
	it.Insert(Interval{15, 20})
	assert.NotNil(it.Root, "Root should not be nil after insert")
	assert.Equal(Interval{15, 20}, it.Root.Interval, "Root interval should be [15, 20] after insert")

	it.Insert(Interval{10, 30})
	assert.Equal(Interval{15, 20}, it.Root.Interval, "Root interval should be [15, 20] after insert")

	it.Insert(Interval{17, 19})
	it.Insert(Interval{5, 20})
	it.Insert(Interval{12, 15})
	it.Insert(Interval{30, 40})

	// Test deletions
	it.Delete(Interval{10, 30})
	assert.Equal(Interval{15, 20}, it.Root.Interval, "Root interval should be [15, 20] after delete")
}

func TestEmptyIntervalTree(t *testing.T) {
	assert := assert.New(t)
	it := &IntervalTree{}

	// Test deletion from an empty tree
	it.Delete(Interval{10, 30})
	assert.Nil(it.Root, "Root should be nil when deleting from an empty tree")
}

// helper function to get absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// isBalanced checks if the tree rooted at node is balanced or not
func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	return abs(leftHeight-rightHeight) <= 1 && isBalanced(node.Left) && isBalanced(node.Right)
}

func TestBalancing(t *testing.T) {
	assert := assert.New(t)
	it := &IntervalTree{}

	// Test insertions
	it.Insert(Interval{15, 20})
	it.Insert(Interval{10, 30})
	it.Insert(Interval{17, 19})
	it.Insert(Interval{5, 20})
	it.Insert(Interval{12, 15})
	it.Insert(Interval{30, 40})

	assert.True(isBalanced(it.Root), "Tree should be balanced after insertions")

	// Test deletions
	it.Delete(Interval{10, 30})

	assert.True(isBalanced(it.Root), "Tree should be balanced after deletions")
}

func TestFindFreeInterval(t *testing.T) {
	assert := assert.New(t)
	it := &IntervalTree{}

	// Test empty tree
	freeInterval := it.FindFreeInterval(10)
	assert.Equal(Interval{1, 11}, freeInterval, "Free interval should be [1, 11] in an empty tree")

	// Add intervals with a gap in the middle
	it.Insert(Interval{15, 20})
	it.Insert(Interval{7, 12})
	it.Insert(Interval{21, 30})
	it.Insert(Interval{30, 40})
	freeInterval = it.FindFreeInterval(5)
	assert.Equal(Interval{1, 6}, freeInterval, "Free interval should be [21, 26]")

	// Fill the gap
	it.Insert(Interval{1, 6})
	freeInterval = it.FindFreeInterval(5)
	assert.Equal(Interval{41, 46}, freeInterval, "Free interval should be [41, 46]")
}
