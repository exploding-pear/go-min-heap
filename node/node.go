// Package node is for the implementation
// of the min-heap nodes
package node

import "fmt"

// Node is the base data structure
// for a min heap. It stores information
// about the value, its parent and children
type Node struct {
	LeftChild  *Node
	RightChild *Node
	Parent     *Node
	value      int
}

// Implementing the String function so
// nodes can be printed out
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// GetValue returns node's value
func (n *Node) GetValue() int {
	return n.value
}

// NewOrphan creates a new Node
// with no parents or children
func NewOrphan(value int) Node {
	return Node{
		LeftChild:  nil,
		RightChild: nil,
		Parent:     nil,
		value:      value,
	}
}

// NewChild attaches a child to the node
// returns false if both children are already taken
// or if one if the child's children is the caller
func (n *Node) NewChild(child *Node) bool {
	switch {
	// ensure that the child's children are not the caller
	case child.LeftChild == n || child.RightChild == n:
		return false
	// ensure that the child is not the parent
	case child == n:
		return false
	// check to see if left child is empty first
	case n.LeftChild == nil:
		n.LeftChild = child
		child.Parent = n
		return true
	// check to see if right child is empty after left
	case n.RightChild == nil:
		n.RightChild = child
		child.Parent = n
		return true
	// returns false if the node already has children
	default:
		return false
	}
}
