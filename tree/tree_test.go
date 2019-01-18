package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNode_IsLeaf(t *testing.T) {
	leaf := Node{}
	if !leaf.IsLeaf() {
		t.Error("Node should be leaf but isn't")
	}

	node := Node{
		Children: []*Node{&Node{}},
	}
	if node.IsLeaf() {
		t.Error("Node is leaf, but shouldn't")
	}
}

func TestNode_IsRoot(t *testing.T) {
	node := Node{
		Children: []*Node{&Node{}},
		Parent:   &Node{},
	}
	if node.IsRoot() {
		t.Error("Node is root, but shouldn't")
	}

	root := Node{}
	if !root.IsRoot() {
		t.Error("Node should be root, but isn't")
	}

}

func TestNode_AppendNode(t *testing.T) {
	root := Node{}
	firstNode := Node{}
	secondNode := Node{}

	root.AppendNode(&firstNode)
	root.AppendNode(&secondNode)

	fmt.Printf("Root Children: %v\n", root.Children)
	if !reflect.DeepEqual(*firstNode.Parent, root) {
		t.Errorf("Parent of firstNode has not been set properly, got %v, expected %v", *firstNode.Parent, root)
	}
	if !reflect.DeepEqual(*secondNode.Parent, root) {
		t.Errorf("Parent of secondNode has not been set properly, got %v, expected %v", *secondNode.Parent, root)
	}

	if !reflect.DeepEqual(*root.Children[0], firstNode) {
		t.Error("First Child of root has not been set properly")
	}
	if !reflect.DeepEqual(*root.Children[1], secondNode) {
		t.Error("Second Child of root has not been set properly")
	}
}
