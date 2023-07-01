package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTree_InsertAndInOrderTraversal(t *testing.T) {
	tree := AVLTree{}
	tree.Insert(5, 10)
	tree.Insert(3, 20)
	tree.Insert(7, 30)
	tree.Insert(2, 40)
	tree.Insert(4, 50)

	expectedOrder := []int{2, 3, 4, 5, 7}
	traversalResult := make([]int, 0)
	traversalFunc := func(node *Node) bool {
		traversalResult = append(traversalResult, node.Key)
		return true
	}

	tree.InOrderTraversal(traversalFunc)

	assert.Equal(t, expectedOrder, traversalResult, "In-order traversal should match expected order")
}

func TestAVLTree_DeleteAndInOrderTraversal(t *testing.T) {
	tree := AVLTree{}
	tree.Insert(5, 10)
	tree.Insert(3, 20)
	tree.Insert(7, 30)
	tree.Insert(2, 40)
	tree.Insert(4, 50)

	tree.Delete(3)

	expectedOrder := []int{2, 4, 5, 7}
	traversalResult := make([]int, 0)
	traversalFunc := func(node *Node) bool {
		traversalResult = append(traversalResult, node.Key)
		return true
	}

	tree.InOrderTraversal(traversalFunc)

	assert.Equal(t, expectedOrder, traversalResult, "In-order traversal after deletion should match expected order")
}

func TestAVLTree_InsertDuplicateKey(t *testing.T) {
	tree := AVLTree{}
	tree.Insert(5, 10)
	tree.Insert(3, 20)
	tree.Insert(7, 30)
	tree.Insert(2, 40)
	tree.Insert(4, 50)
	tree.Insert(3, 60) // Inserting with duplicate key

	expectedOrder := []int{2, 3, 4, 5, 7}
	traversalResult := make([]int, 0)
	traversalFunc := func(node *Node) bool {
		traversalResult = append(traversalResult, node.Key)
		return true
	}

	tree.InOrderTraversal(traversalFunc)

	assert.Equal(t, expectedOrder, traversalResult, "In-order traversal should ignore duplicate keys")
}

func TestAVLTree_DeleteNonExistentKey(t *testing.T) {
	tree := AVLTree{}
	tree.Insert(5, 10)
	tree.Insert(3, 20)
	tree.Insert(7, 30)
	tree.Insert(2, 40)
	tree.Insert(4, 50)

	tree.Delete(8) // Deleting non-existent key

	expectedOrder := []int{2, 3, 4, 5, 7}
	traversalResult := make([]int, 0)
	traversalFunc := func(node *Node) bool {
		traversalResult = append(traversalResult, node.Key)
		return true
	}

	tree.InOrderTraversal(traversalFunc)

	assert.Equal(t, expectedOrder, traversalResult, "In-order traversal should not be affected by deleting non-existent key")
}
