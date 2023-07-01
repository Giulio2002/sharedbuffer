package types

import (
	"math"
)

type Node struct {
	Key    int
	Value  int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) Insert(key, value int) {
	if t.Root == nil {
		t.Root = &Node{Key: key, Value: value, Height: 1}
		return
	}

	stack := []*Node{}
	curr := t.Root

	for curr != nil {
		stack = append(stack, curr)

		if key < curr.Key {
			if curr.Left == nil {
				curr.Left = &Node{Key: key, Value: value, Height: 1}
				break
			}
			curr = curr.Left
		} else if key > curr.Key {
			if curr.Right == nil {
				curr.Right = &Node{Key: key, Value: value, Height: 1}
				break
			}
			curr = curr.Right
		} else {
			// Update the value if the key already exists
			curr.Value = value
			return
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		node.Height = int(math.Max(float64(t.getHeight(node.Left)), float64(t.getHeight(node.Right)))) + 1

		balance := t.getBalance(node)

		if balance > 1 && key < node.Left.Key {
			node = t.rotateRight(node)
		}

		if balance < -1 && key > node.Right.Key {
			node = t.rotateLeft(node)
		}

		if balance > 1 && key > node.Left.Key {
			node.Left = t.rotateLeft(node.Left)
			node = t.rotateRight(node)
		}

		if balance < -1 && key < node.Right.Key {
			node.Right = t.rotateRight(node.Right)
			node = t.rotateLeft(node)
		}

		if i > 0 {
			parent := stack[i-1]
			if parent.Left == node {
				parent.Left = node
			} else {
				parent.Right = node
			}
		} else {
			t.Root = node
		}
	}
}

func (t *AVLTree) Delete(key int) {
	t.Root = t.deleteNode(t.Root, key)
}

func (t *AVLTree) deleteNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.Key {
		node.Left = t.deleteNode(node.Left, key)
	} else if key > node.Key {
		node.Right = t.deleteNode(node.Right, key)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			minNode := t.findMin(node.Right)
			node.Key = minNode.Key
			node.Value = minNode.Value
			node.Right = t.deleteNode(node.Right, minNode.Key)
		}
	}

	if node != nil {
		node.Height = int(math.Max(float64(t.getHeight(node.Left)), float64(t.getHeight(node.Right)))) + 1

		balance := t.getBalance(node)

		if balance > 1 && t.getBalance(node.Left) >= 0 {
			node = t.rotateRight(node)
		}

		if balance > 1 && t.getBalance(node.Left) < 0 {
			node.Left = t.rotateLeft(node.Left)
			node = t.rotateRight(node)
		}

		if balance < -1 && t.getBalance(node.Right) <= 0 {
			node = t.rotateLeft(node)
		}

		if balance < -1 && t.getBalance(node.Right) > 0 {
			node.Right = t.rotateRight(node.Right)
			node = t.rotateLeft(node)
		}
	}

	return node
}

func (t *AVLTree) findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *AVLTree) rotateLeft(z *Node) *Node {
	y := z.Right
	T2 := y.Left

	y.Left = z
	z.Right = T2

	z.Height = int(math.Max(float64(t.getHeight(z.Left)), float64(t.getHeight(z.Right)))) + 1
	y.Height = int(math.Max(float64(t.getHeight(y.Left)), float64(t.getHeight(y.Right)))) + 1

	return y
}

func (t *AVLTree) rotateRight(z *Node) *Node {
	y := z.Left
	T3 := y.Right

	y.Right = z
	z.Left = T3

	z.Height = int(math.Max(float64(t.getHeight(z.Left)), float64(t.getHeight(z.Right)))) + 1
	y.Height = int(math.Max(float64(t.getHeight(y.Left)), float64(t.getHeight(y.Right)))) + 1

	return y
}

func (t *AVLTree) getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (t *AVLTree) getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.Left) - t.getHeight(node.Right)
}

func (t *AVLTree) InOrderTraversal(traversalFunc func(*Node) bool) {
	stack := []*Node{}
	curr := t.Root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		shouldContinue := traversalFunc(curr)
		if !shouldContinue {
			return
		}

		curr = curr.Right
	}
}
