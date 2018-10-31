package test

import (
	"testing"

	"github.com/lunjon/tree"
)

func TestCreate(t *testing.T) {
	// Arange
	tree := tree.Create()

	// Assert
	if tree == nil || len(tree.Root.Children) != 0 {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	// Arange
	tree := tree.Create()

	// Act and assert
	size := tree.Size()
	if size != 0 {
		t.Fail()
	}

	//Test size of the following tree:
	//	   Root
	//	  1     5
	//	 2 3
	//	4

	rem2 := tree.Root.Add(1)
	tree.Root.Children[0].Add(2)
	tree.Root.Children[0].Add(3)
	rem1 := tree.Root.Children[0].Children[0].Add(4)
	tree.Root.Add(5)

	size = tree.Size()
	if size != 5 {
		t.Log(size)
		t.Fail()
	}

	// Remove a leaf node
	rem1.Remove()
	size = tree.Size()
	if size != 4 {
		t.Log(size)
		t.Fail()
	}

	// Remove a node with descendants
	rem2.Remove()
	size = tree.Size()
	if size != 1 {
		t.Log(size)
		t.Fail()
	}
}
