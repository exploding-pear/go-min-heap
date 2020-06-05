package node

import "testing"

// Test that NewOrphan works as expected
func TestNewOrphan(t *testing.T) {
	newNode := NewOrphan(4)

	if newNode.LeftChild != nil {
		t.Error("left child not nil")
	}
	if newNode.RightChild != nil {
		t.Error("right child not nil")
	}
	if newNode.Parent != nil {
		t.Error("parent not nil")
	}
	if newNode.value != 4 {
		t.Error("value of node is not what was expected")
	}
}

// naive test of NewChild
func TestNewChild(t *testing.T) {
	n1 := NewOrphan(5)
	n2 := NewOrphan(12)
	n1.NewChild(&n2)

	if n1.LeftChild.value != 12 {
		t.Error("assignment of parent to child failed")
	}
}

// test that a child's children cannot be a child's parent
func TestGrandchildNotParent(t *testing.T) {
	n1 := NewOrphan(1)
	n2 := NewOrphan(2)

	// making n1 the parent of n2
	if !n1.NewChild(&n2) {
		t.Error("normal child creation failed")
	}

	// expected to fail - making child's child the child's parent
	if n2.NewChild(&n1) {
		t.Error("child's child was assigned to the child's parent")
	}
}
