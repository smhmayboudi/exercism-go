package linkedlist

import (
	"errors"
)

var ErrEmptyList = errors.New("empty list")

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
	list := &List{}
	for i := 0; i < len(elements); i++ {
		list.Push(elements[i])
	}
	return list
}

func (n *Node) Next() *Node {
	return n.NextNode
}

func (n *Node) Prev() *Node {
	return n.PrevNode
}

func (l *List) Unshift(v any) {
	node := &Node{Value: v}
	node.NextNode = l.FirstNode
	if l.FirstNode != nil {
		l.FirstNode.PrevNode = node
	}
	l.FirstNode = node
	if l.LastNode == nil {
		l.LastNode = node
	}
}

func (l *List) Push(v any) {
	node := &Node{Value: v}
	node.PrevNode = l.LastNode
	if l.LastNode != nil {
		l.LastNode.NextNode = node
	}
	l.LastNode = node
	if l.FirstNode == nil {
		l.FirstNode = node
	}
}

func (l *List) Shift() (any, error) {
	if l.FirstNode == nil {
		return nil, ErrEmptyList
	}
	val := l.FirstNode.Value
	l.FirstNode = l.FirstNode.Next()
	if l.FirstNode == nil {
		l.LastNode = nil
	} else {
		l.FirstNode.PrevNode = nil
	}
	return val, nil

}

func (l *List) Pop() (any, error) {
	if l.LastNode == nil {
		return nil, ErrEmptyList
	}
	val := l.LastNode.Value
	l.LastNode = l.LastNode.Prev()
	if l.LastNode == nil {
		l.FirstNode = nil
	} else {
		l.LastNode.NextNode = nil
	}
	return val, nil
}

func (l *List) Reverse() {
	vals := []any{}
	for {
		v, err := l.Pop()
		if err != nil {
			break
		}
		vals = append(vals, v)
	}
	for _, v := range vals {
		l.Push(v)
	}
}

func (l *List) First() *Node {
	return l.FirstNode
}

func (l *List) Last() *Node {
	return l.LastNode
}
