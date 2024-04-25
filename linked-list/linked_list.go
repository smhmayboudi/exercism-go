package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value    any
	PrevNode *Node
	NextNode *Node
}

type List struct {
	FirstNode *Node
	LastNode  *Node
}

func NewList(elements ...any) *List {
	list := List{}
	for index := 0; index < len(elements); index++ {
		value := elements[index]
		node := Node{
			Value: value,
		}
		if 1 < index {
			node.PrevNode = list.LastNode
		}
		list.LastNode = &node
	}
	return &list
}

func (n *Node) Next() *Node {
	return n.NextNode
}

func (n *Node) Prev() *Node {
	return n.PrevNode
}

func (l *List) Unshift(v any) {
	node := Node{
		Value:    v,
		PrevNode: l.LastNode,
		NextNode: nil,
	}
	l.LastNode = &node
}

func (l *List) Push(v any) {
	node := Node{
		Value:    v,
		PrevNode: nil,
		NextNode: l.FirstNode,
	}
	l.FirstNode = &node
}

func (l *List) Shift() (any, error) {
	frontNode := l.LastNode
	if frontNode == nil {
		return nil, errors.New("empty list")
	}
	if frontNode.Prev() == nil {
		return nil, errors.New("last item")
	}
	l.LastNode = frontNode.Prev()
	return frontNode.Value, nil
}

func (l *List) Pop() (any, error) {
	backNode := l.FirstNode
	if backNode == nil {
		return nil, errors.New("empty list")
	}
	if backNode.Next() == nil {
		return nil, errors.New("last item")
	}
	l.FirstNode = backNode.Next()
	return backNode.Value, nil
}

func (l *List) Reverse() {
	list := List{}
	node := l.First()
	for {
		list.Unshift(node.Value)
		if node.Next() == nil {
			break
		}
		node = node.Next()
	}
	l = &list
}

func (l *List) First() *Node {
	return l.FirstNode
}

func (l *List) Last() *Node {
	return l.LastNode
}
