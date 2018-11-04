package node

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Arange
	root := Node{Value: nil, Parent: nil, Children: []*Node{}}

	// Act
	val := "val"
	node := root.Add(val)

	if node.Value != root.Children[0].Value {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	// Arange
	root := Node{Value: nil, Parent: nil, Children: []*Node{}}
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
	root := Node{Value: nil, Parent: nil, Children: []*Node{}}

	if root.HasChildren() {
		t.Fail()
	}

	root.Add(1)
	if !root.HasChildren() {
		t.Fail()
	}
}
