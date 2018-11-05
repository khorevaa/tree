package test

import (
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
