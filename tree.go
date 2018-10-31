package tree

type Tree struct {
	Root *Node
}

type Node struct {
	Value    interface{}
	Parent   *Node
	Children []*Node
}

func Create() *Tree {
	root := Node{nil, nil, []*Node{}}
	return &Tree{&root}
}

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
