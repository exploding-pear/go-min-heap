// This package is a go implementation of a min-heap
package main

import (
	"fmt"

	"github.com/exploding-pear/go-min-heap/node"
)

// PrintTree prints out tree
func PrintTree(n *node.Node) {
	ch := make(chan int)
	list := make([]int, 0)
	go func() {
		defer close(ch)
		traverseTree(n, ch)
	}()
	for i := range ch {
		list = append(list, i)
	}
	fmt.Println(list)
}

func traverseTree(n *node.Node, c chan int) {
	if n == nil {
		return
	}
	c <- n.GetValue()
	traverseTree(n.LeftChild, c)
	traverseTree(n.RightChild, c)
}

func main() {
	n1 := node.NewOrphan(4)
	n2 := node.NewOrphan(10)
	n3 := node.NewOrphan(15)
	n4 := node.NewOrphan(43)
	n5 := node.NewOrphan(71)
	n6 := node.NewOrphan(400)

	n1.NewChild(&n2)
	n1.NewChild(&n3)
	n2.NewChild(&n4)
	n2.NewChild(&n5)
	n3.NewChild(&n6)

	PrintTree(&n1)
}
