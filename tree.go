package tree

import (
	"fmt"
	"io"
	"strings"
)

// Tree represent a tree structure and holds a reference to the root
type Tree struct {
	Root *Node
}

// Create returns a new empty tree
func Create(value interface{}) *Tree {
	root := Node{Value: value, Parent: nil, Children: []*Node{}}
	return &Tree{&root}
}

// Size returns the number of nodes in the tree, the root excluded
func (tree *Tree) Size() int {
	if !tree.Root.HasChildren() {
		return 1
	}
	return recSize(tree.Root)
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

// ForEachBSF iterates over this tree using breadth first
// search and calls callback for each node
func (tree *Tree) ForEachBSF(callback func(*Node)) {
	if !tree.Root.HasChildren() {
		callback(tree.Root)
	}

	recForEachBSF(tree.Root, callback)
}

func recForEachBSF(node *Node, callback func(*Node)) {
	callback(node)
	for i := range node.Children {
		recForEachBSF(node.Children[i], callback)
	}
}

// Implement the fmt.Stringer interface
func (tree *Tree) String() string {
	if !tree.Root.HasChildren() {
		return ""
	}

	var builder strings.Builder
	var levelsEnded []int

	fmt.Println(tree.Root)
	printNodes(&builder, 0, levelsEnded, tree.Root.Children)
	return builder.String()
}

func printNodes(wr io.Writer, level int, levelsEnded []int, nodes []*Node) {
	for i, node := range nodes {
		edge := midEdge
		if i == len(nodes)-1 {
			levelsEnded = append(levelsEnded, level)
			edge = endEdge
		}
		printValues(wr, level, levelsEnded, edge, node)
		if len(node.Children) > 0 {
			printNodes(wr, level+1, levelsEnded, node.Children)
		}
	}
}

func printValues(wr io.Writer, level int, levelsEnded []int, edge string, node *Node) {
	for i := 0; i < level; i++ {
		if isEnded(levelsEnded, i) {
			fmt.Fprint(wr, "    ")
			continue
		}
		fmt.Fprintf(wr, "%s   ", linkEdge)
	}
	fmt.Fprintf(wr, "%s %v\n", edge, node)
}

func isEnded(levelsEnded []int, level int) bool {
	for _, l := range levelsEnded {
		if l == level {
			return true
		}
	}
	return false
}

const (
	linkEdge string = "│"
	midEdge  string = "├──"
	endEdge  string = "└──"
)

func recString(node *Node, indent string, b *strings.Builder) {
	if !node.HasChildren() {
		fmt.Fprintf(b, "%s├-- %v\n", indent, node)
	} else {
		if node.Value != nil {
			fmt.Fprintf(b, "%s├-- %v\n", indent, node)
		}
		indent = fmt.Sprintf("%s|   ", indent)
		for _, child := range node.Children {
			recString(child, indent, b)
		}
	}
}
