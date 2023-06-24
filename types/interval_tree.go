package types

// Interval is a simple structure with two integer fields, Low and High
type Interval struct {
	Low, High int
}

// Node is the main structure in the IntervalTree, which stores an Interval, a Max value, a Height,
// and pointers to two other nodes (Left and Right).
type Node struct {
	Interval    Interval
	Max         int
	Height      int
	Left, Right *Node
}

// IntervalTree is a self-balancing binary search tree (AVL Tree), that stores Nodes
type IntervalTree struct {
	Root *Node
}

// NewNode creates a new Node with given Interval, and its height set to 1
func NewNode(interval Interval) *Node {
	return &Node{
		Interval: interval,
		Max:      interval.High,
		Height:   1,
	}
}

// max returns the maximum between two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// height returns the height of a given node, or 0 if the node is nil
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// getBalance calculates and returns the balance factor of a node,
// which is the difference in height between its left child and right child
func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// rotate performs a rotation (either left or right) around a given node
func rotate(node *Node, isLeftRotation bool) *Node {
	var newRoot *Node

	if isLeftRotation {
		newRoot = node.Right
		node.Right = newRoot.Left
		newRoot.Left = node
	} else {
		newRoot = node.Left
		node.Left = newRoot.Right
		newRoot.Right = node
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1
	node.Max = max(node.Interval.High, max(getMax(node.Left), getMax(node.Right)))

	newRoot.Height = max(height(newRoot.Left), height(newRoot.Right)) + 1
	newRoot.Max = max(newRoot.Interval.High, max(getMax(newRoot.Left), getMax(newRoot.Right)))

	return newRoot
}

// insert is a recursive function that inserts a new interval in the tree, maintaining the AVL property
func insert(node *Node, interval Interval) *Node {
	// If the node is nil, return a new node with the given interval
	if node == nil {
		return NewNode(interval)
	}

	// Otherwise, insert the new interval in the left or right subtree depending on the value of its low end
	if interval.Low < node.Interval.Low {
		node.Left = insert(node.Left, interval)
	} else if interval.Low > node.Interval.Low {
		node.Right = insert(node.Right, interval)
	} else {
		return node
	}

	// Update the height and max value of the current node
	node.Height = 1 + max(height(node.Left), height(node.Right))
	node.Max = max(node.Interval.High, max(getMax(node.Left), getMax(node.Right)))

	// Check the balance factor and perform rotations if the node became unbalanced
	balance := getBalance(node)

	if balance > 1 && interval.Low < node.Left.Interval.Low {
		return rotate(node, false)
	}

	if balance < -1 && interval.Low > node.Right.Interval.Low {
		return rotate(node, true)
	}

	if balance > 1 && interval.Low > node.Left.Interval.Low {
		node.Left = rotate(node.Left, true)
		return rotate(node, false)
	}

	if balance < -1 && interval.Low < node.Right.Interval.Low {
		node.Right = rotate(node.Right, false)
		return rotate(node, true)
	}

	return node
}

// minValueNode returns the node with the smallest low value in a given tree
func minValueNode(node *Node) *Node {
	current := node

	for current.Left != nil {
		current = current.Left
	}

	return current
}

// delete removes a node with a given interval from the tree and maintains the AVL property
func delete(node *Node, interval Interval) *Node {
	if node == nil {
		return node
	}

	if interval.Low < node.Interval.Low {
		node.Left = delete(node.Left, interval)
	} else if interval.Low > node.Interval.Low {
		node.Right = delete(node.Right, interval)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		tmp := minValueNode(node.Right)
		node.Interval = tmp.Interval
		node.Right = delete(node.Right, tmp.Interval)
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	node.Max = max(node.Interval.High, max(getMax(node.Left), getMax(node.Right)))

	balance := getBalance(node)

	if balance > 1 {
		if getBalance(node.Left) < 0 {
			node.Left = rotate(node.Left, true)
		}
		return rotate(node, false)
	}

	if balance < -1 {
		if getBalance(node.Right) > 0 {
			node.Right = rotate(node.Right, false)
		}
		return rotate(node, true)
	}

	return node
}

// getMax returns the Max value of a node, or -1 if the node is nil
func getMax(node *Node) int {
	if node == nil {
		return -1
	}
	return node.Max
}

// Insert is a wrapper function for insert that adds a new interval to the tree
func (it *IntervalTree) Insert(interval Interval) {
	it.Root = insert(it.Root, interval)
}

// Delete is a wrapper function for delete that removes an interval from the tree
func (it *IntervalTree) Delete(interval Interval) {
	it.Root = delete(it.Root, interval)
}

// findFreeIntervalHelper finds a free interval of size n within the subtree rooted at node.
func findFreeIntervalHelper(node *Node, n int, previousHigh int) Interval {
	if node == nil {
		return Interval{previousHigh + 1, previousHigh + n + 1} // If the subtree is empty, return the interval after the previous high
	}

	if node.Left != nil && node.Left.Max >= n {
		// If the maximum endpoint in the left subtree is greater than or equal to n,
		// recursively search in the left subtree.
		return findFreeIntervalHelper(node.Left, n, previousHigh)
	}

	if node.Interval.Low-previousHigh > n {
		// If the gap between the current interval and the previous high is greater than n,
		// we have found a free interval.
		return Interval{previousHigh + 1, previousHigh + n + 1}
	}

	// Otherwise, recursively search in the right subtree.
	return findFreeIntervalHelper(node.Right, n, node.Interval.High)
}

// FindFreeInterval finds and returns the smallest interval of size n that does not overlap with any intervals in the tree.
func (it *IntervalTree) FindFreeInterval(n int) Interval {
	return findFreeIntervalHelper(it.Root, n, 0)
}
