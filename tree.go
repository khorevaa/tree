package tree

import (
	"fmt"
	"strings"
)

// Tree represent a tree structure and holds a reference to the root
type Tree struct {
	Root *Node
}

// Node is the data structure which the tree is consists of
type Node struct {
	Value    fmt.Stringer
	Parent   *Node
	Children []*Node
}

// Create returns a new empty tree
func Create() *Tree {
	root := Node{nil, nil, []*Node{}}
	return &Tree{&root}
}

// Size returns the number of nodes in the tree, the root excluded
func (tree *Tree) Size() int {
	if len(tree.Root.Children) == 0 {
		return 0
	}
	return recSize(tree.Root) - 1
}

func recSize(node *Node) int {
	if len(node.Children) == 0 {
		return 1
	}

	sum := 1
	for _, child := range node.Children {
		sum += recSize(child)
	}

	return sum
}

func (tree *Tree) String() string {
	if len(tree.Root.Children) == 0 {
		return ""
	}

	fmt.Println("Root")
	var builder strings.Builder
	recString(tree.Root, 0, &builder)
	return strings.TrimRight(builder.String(), "\n")
}

func recString(node *Node, nbrIndents int, b *strings.Builder) {
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

// Add adds a new node under the receiving node with the specified value
func (node *Node) Add(value fmt.Stringer) *Node {
	// TODO: Check for nil value and return an error?

	new := &Node{value, node, []*Node{}}
	node.Children = append(node.Children, new)
	return new
}

// Remove deletes this node from the tree and all its descdendants
func (node *Node) Remove() {
	// Check if it's the root
	if node.Parent == nil {
		return // Do nothing
	}

	// Remove this node from its parent
	parent := node.Parent
	for i, child := range parent.Children {
		if child == node {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			break
		}
	}
}

// HasChildren returns true if node has any child nodes
func (node *Node) HasChildren() bool {
	return len(node.Children) != 0
}

func (node *Node) String() string {
	return fmt.Sprint(node.Value)
}
