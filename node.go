package tree

import (
	"fmt"
)

// Node is the data structure which the tree is consists of
type Node struct {
	Value    interface{}
	Parent   *Node
	Children []*Node
}

// Add adds a new node under the receiving node with the specified value
func (node *Node) Add(value interface{}) *Node {
	// TODO: Check for nil value and return an error?

	new := &Node{value, node, []*Node{}}
	node.Children = append(node.Children, new)
	return new
}

// Remove this node from the tree and all its descdendants
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

// ForEachDFS runs callback for each node using depth-first-search
func (node *Node) ForEachDFS(callback func(*Node)) {
	if !node.HasChildren() {
		callback(node)
	} else {
		recForEachDFS(node, callback)
	}
}

func recForEachDFS(node *Node, callback func(*Node)) {
	if !node.HasChildren() {
		callback(node)
	} else {
		for _, n := range node.Children {
			recForEachDFS(n, callback)
		}
		callback(node)
	}
}

// HasChildren returns true if node has any child nodes
func (node *Node) HasChildren() bool {
	return len(node.Children) != 0
}

func (node *Node) String() string {
	return fmt.Sprint(node.Value)
}
