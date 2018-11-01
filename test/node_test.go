package test

import (
	"testing"

	"github.com/lunjon/tree"
)

func TestAdd(t *testing.T) {
	// Arange
	tree := tree.Create()

	// Act
	val := "val"
	node := tree.Root.Add(val)

	if node.Value != tree.Root.Children[0].Value {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	// Arange
	tree := tree.Create()
	node := tree.Root.Add(1)

	// Act
	node.Remove()

	// Assert
	if len(tree.Root.Children) != 0 {
		t.Fail()
	}
}

func TestHasChildren(t *testing.T) {
	// Arange
	tree := tree.Create()

	if tree.Root.HasChildren() {
		t.Fail()
	}

	tree.Root.Add(1)
	if !tree.Root.HasChildren() {
		t.Fail()
	}
}
