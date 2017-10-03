package binary

import (
	"errors"
)

// Node is main type
type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

// Tree is root node
type Tree struct {
	Root *Node
}

// Insert is method for Node
func (n *Node) Insert(value, data string) error {
	// If node is nil -> Error
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case value == n.Value:
		// can't determine left or right, but error is nil
		return nil
	case value < n.Value:
		if n.Left == nil {
			// If left is nil then create new pointer
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			// If right is nil then create new pointer
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

// Search is method for Node
func (n *Node) Search(s string) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Search(s)
	default:
		return n.Right.Search(s)
	}
}

// Insert for Tree
func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}

	return t.Root.Insert(value, data)
}

// Search for Tree
func (t *Tree) Search(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Search(s)
}
