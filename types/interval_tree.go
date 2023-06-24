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

// rightRotate performs a right rotation around a given node
func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.Height = max(height(y.Left), height(y.Right)) + 1
	y.Max = max(y.Interval.High, max(getMax(y.Left), getMax(y.Right)))
	x.Height = max(height(x.Left), height(x.Right)) + 1
	x.Max = max(x.Interval.High, max(getMax(x.Left), getMax(x.Right)))
	return x
}

// leftRotate performs a left rotation around a given node
func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	x.Height = max(height(x.Left), height(y.Right)) + 1
	x.Max = max(x.Interval.High, max(getMax(x.Left), getMax(x.Right)))
	y.Height = max(height(y.Left), height(y.Right)) + 1
	y.Max = max(y.Interval.High, max(getMax(y.Left), getMax(y.Right)))
	return y
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
		return rightRotate(node)
	}

	if balance < -1 && interval.Low > node.Right.Interval.Low {
		return leftRotate(node)
	}

	if balance > 1 && interval.Low > node.Left.Interval.Low {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && interval.Low < node.Right.Interval.Low {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
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
		if (node.Left == nil) || (node.Right == nil) {
			var temp *Node
			if temp = node.Left; temp == nil {
				temp = node.Right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
			temp = nil
		} else {
			temp := minValueNode(node.Right)
			node.Interval = temp.Interval
			node.Right = delete(node.Right, temp.Interval)
		}
	}

	if node == nil {
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	node.Max = max(node.Interval.High, max(getMax(node.Left), getMax(node.Right)))

	balance := getBalance(node)

	if balance > 1 && getBalance(node.Left) >= 0 {
		return rightRotate(node)
	}

	if balance < -1 && getBalance(node.Right) <= 0 {
		return leftRotate(node)
	}

	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
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

// FindFreeInterval finds and returns the smallest interval of size n that does not overlap with any intervals in the tree
func (it *IntervalTree) FindFreeInterval(n int) Interval {
	nodes := inOrder(it.Root)
	if len(nodes) == 0 {
		return Interval{1, n + 1}
	}

	// Check for space before the first interval
	if nodes[0].Interval.Low > n {
		return Interval{1, n + 1}
	}

	// Check for space between intervals
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Interval.Low-nodes[i-1].Interval.High > n {
			return Interval{nodes[i-1].Interval.High + 1, nodes[i-1].Interval.High + n + 1}
		}
	}

	// Check for space after the last interval
	return Interval{nodes[len(nodes)-1].Interval.High + 1, nodes[len(nodes)-1].Interval.High + n + 1}
}

// inOrder returns an array of pointers to nodes obtained by performing an in-order traversal of a given tree
func inOrder(node *Node) []*Node {
	nodes := []*Node{}
	if node != nil {
		nodes = append(nodes, inOrder(node.Left)...)
		nodes = append(nodes, node)
		nodes = append(nodes, inOrder(node.Right)...)
	}
	return nodes
}
