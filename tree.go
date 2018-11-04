package tree

import (
	"fmt"
	"strings"

	"github.com/lunjon/tree/internal/node"
)

// Tree represent a tree structure and holds a reference to the root
type Tree struct {
	Root *node.Node
}

// Create returns a new empty tree
func Create() *Tree {
	root := node.Node{Value: nil, Parent: nil, Children: []*node.Node{}}
	return &Tree{&root}
}

// Size returns the number of nodes in the tree, the root excluded
func (tree *Tree) Size() int {
	if len(tree.Root.Children) == 0 {
		return 0
	}
	return recSize(tree.Root) - 1
}

func recSize(node *node.Node) int {
	if len(node.Children) == 0 {
		return 1
	}

	sum := 1
	for _, child := range node.Children {
		sum += recSize(child)
	}

	return sum
}

// Implement the fmt.Stringer interface
func (tree *Tree) String() string {
	if len(tree.Root.Children) == 0 {
		return ""
	}

	fmt.Println("Root")
	var builder strings.Builder
	recString(tree.Root, 0, &builder)
	return strings.TrimRight(builder.String(), "\n")
}

func recString(node *node.Node, nbrIndents int, b *strings.Builder) {
	indent := strings.Repeat("\t", nbrIndents)
	if !node.HasChildren() {
		fmt.Fprintf(b, "%s%v\n", indent, node)
	} else {
		if node.Value != nil {
			fmt.Fprintf(b, "%s%v\n", indent, node)
		}
		for _, child := range node.Children {
			recString(child, nbrIndents+1, b)
		}
	}
}
