package tree

// Tree represent a tree structure and holds a reference to the root
type Tree struct {
	Root *Node
}

// Node is the data structure which the tree is consists of
type Node struct {
	Value    interface{}
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

	return stringRec(tree.Root)
}

func stringRec(node *Node) string {
	return ""
}

// Add adds a new node under the receiving node with the specified value
func (node *Node) Add(value interface{}) *Node {
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
