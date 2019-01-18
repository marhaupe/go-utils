package tree

type Tree struct {
	root *Node
}

type Node struct {
	parent   *Node
	children []*Node
	val      interface{}
}

func (n Node) IsLeaf() bool {
	return n.children == nil
}

func (n Node) IsRoot() bool {
	return n.parent == nil
}
