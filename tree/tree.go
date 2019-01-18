package tree

// Tree represents a generic tree implementation
type Tree struct {
	Root *Node
}

// Node represents a node element in the above tree implementation.
// A node may be a leaf, a node with children or a root node.
type Node struct {
	Parent   *Node
	Children []*Node
	Value    interface{}
}

// IsLeaf returns whether the Node is a leaf or not
func (n *Node) IsLeaf() bool {
	return n.Children == nil
}

// IsRoot returns whether the Node is a the root Node of the tree or not
func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

// AppendNode inserts the child as child node of n in the tree
func (n *Node) AppendNode(child *Node) {
	child.Parent = n
	if n.Children == nil {
		n.Children = make([]*Node, 0, 2)
	}
	n.Children = append(n.Children, child)
}
