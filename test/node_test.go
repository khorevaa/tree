package test

import (
	"log"
	"testing"

	"github.com/lunjon/tree"
)

func TestAdd(t *testing.T) {
	// Arange
	root := tree.Node{Value: nil, Parent: nil, Children: []*tree.Node{}}

	// Act
	val := "val"
	node := root.Add(val)

	if node.Value != root.Children[0].Value {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	// Arange
	root := tree.Node{Value: nil, Parent: nil, Children: []*tree.Node{}}
	node := root.Add(1)

	// Act
	node.Remove()

	// Assert
	if len(root.Children) != 0 {
		t.Fail()
	}
}

func TestHasChildren(t *testing.T) {
	// Arange
	root := tree.Node{Value: nil, Parent: nil, Children: []*tree.Node{}}

	if root.HasChildren() {
		t.Fail()
	}

	root.Add(1)
	if !root.HasChildren() {
		t.Fail()
	}
}

func TestForEachDSF(t *testing.T) {
	root := tree.Node{Value: 0, Parent: nil, Children: []*tree.Node{}}
	one := root.Add(1)
	two := one.Add(2)
	two.Add(3)

	values := make([]int, 0)
	callback := func(node *tree.Node) {
		values = append(values, node.Value.(int))
	}

	root.ForEachDFS(callback)

	if len(values) != 4 {
		log.Printf("len == %d", len(values))
		t.Fail()
	}

	if values[0] != 3 {
		t.Fail()
	}
	if values[1] != 2 {
		t.Fail()
	}
}
