package test

import (
	"fmt"
	"testing"

	"github.com/lunjon/tree"
)

func TestCreateDefault(t *testing.T) {
	// Arange
	tree := tree.Create("root")

	// Assert
	if tree == nil || len(tree.Root.Children) != 0 {
		t.Fail()
	}
}

func TestCreateWithRootValue(t *testing.T) {
	// Arange
	tree := tree.Create("root")

	// Assert
	if tree == nil || len(tree.Root.Children) != 0 {
		t.Fail()
	}

	if tree.Root.Value != "root" {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	// Arange
	tree := tree.Create("root")

	// Act and assert
	size := tree.Size()
	if size != 1 {
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
	if size != 6 {
		t.Log(size)
		t.Fail()
	}

	// Remove a leaf node
	rem1.Remove()
	size = tree.Size()
	if size != 5 {
		t.Log(size)
		t.Fail()
	}

	// Remove a node with descendants
	rem2.Remove()
	size = tree.Size()
	if size != 2 {
		t.Log(size)
		t.Fail()
	}
}

func TestString(t *testing.T) {
	// Arange
	tree := tree.Create("root")

	// Act
	n1 := tree.Root.Add(1)
	tree.Root.Add(2)
	n1.Add(3)

	// Assert
	str := fmt.Sprint(tree)

	if str == "" {
		t.Fail()
	}
}

func TestForEachBSF(t *testing.T) {
	// Arange
	tr := buildTree()

	n := 0
	callback := func(node *tree.Node) {
		n++
	}

	tr.ForEachBSF(callback)

	if n != tr.Size() {
		t.Logf("Expected %d but got %d", tr.Size(), n)
		t.Fail()
	}
}

func buildTree() *tree.Tree {
	// Build the following tree:
	//	   Root
	//	  1     5
	//	 2 3
	//	4

	tree := tree.Create("root")
	tree.Root.Add(1)
	tree.Root.Children[0].Add(2)
	tree.Root.Children[0].Add(3)
	tree.Root.Children[0].Children[0].Add(4)
	tree.Root.Add(5)
	return tree
}
