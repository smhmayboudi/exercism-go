package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	LeftNode  *Node
	RightNode *Node
}

type List []Node

func NewList(elements ...any) *List {
	for i := 0; i < len(elements); i++ {

	}
	panic("Please implement the NewList function")
}

func (n *Node) Next() *Node {
	return n.RightNode
}

func (n *Node) Prev() *Node {
	return n.LeftNode
}

func (l *List) Unshift(v any) {
	for _, value := range *l {
		if value.RightNode == nil {
			node, _ := v.(Node)
			node.RightNode = nil
			node.LeftNode = &value
			value.RightNode = &node
		}
	}
}

func (l *List) Push(v any) {
	for _, value := range *l {
		if value.LeftNode == nil {
			node, _ := v.(Node)
			node.RightNode = &value
			node.LeftNode = nil
			value.LeftNode = &node
		}
	}
}

func (l *List) Shift() (any, error) {
	for _, value := range *l {
		if value.RightNode == nil {
			if value.LeftNode == nil {
				return nil, errors.New("nothing more")
			}
			value.LeftNode.RightNode = nil
		}
	}
	return l, nil
}

func (l *List) Pop() (any, error) {
	for _, value := range *l {
		if value.LeftNode == nil {
			if value.RightNode == nil {
				return nil, errors.New("nothing more")
			}
			value.RightNode.LeftNode = nil
		}
	}
	return l, nil
}

func (l *List) Reverse() {
	list := List{}
	for i := 0; i < len(*l); i++ {
		v := (*l)[i]
	}
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	panic("Please implement the First function")
}

func (l *List) Last() *Node {
	panic("Please implement the Last function")
}
