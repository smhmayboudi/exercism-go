package linkedlist

import (
	"errors"
	"slices"
)

// Define the List and Element types here.
type List struct {
	head *Node
}

type Node struct {
	song int
	next *Node
}

func New(elements []int) *List {
	out := List{}
	for i := 0; i < len(elements); i++ {
		v := elements[i]
		out.Push(v)
	}
	return &out
}

func (l *List) Size() int {
	i := 0
	node := l.head
	if node != nil {
		i++
	}
	for ; node != nil && node.next != nil; i++ {
		node = node.next
	}
	return i
}

func (l *List) Push(element int) {
	node := Node{
		song: element,
		next: nil,
	}
	node2 := l.head
	if node2 == nil {
		l.head = &node
	} else {
		for node2.next != nil {
			node2 = node2.next
		}
		node2.next = &node
	}
}

func (l *List) Pop() (int, error) {
	nodeOLD := l.head
	nodeNEW := l.head
	if nodeNEW == nil {
		return 0, errors.New("no node available")
	} else {
		if nodeNEW.next != nil {
			for nodeNEW.next != nil {
				nodeOLD = nodeNEW
				nodeNEW = nodeNEW.next
			}
			nodeOLD.next = nil
		} else {
			l.head = nil
		}
		return nodeNEW.song, nil
	}
}

func (l *List) Array() []int {
	out := []int{}
	node := l.head
	for i := 0; node != nil; i++ {
		out = append(out, node.song)
		if node.next == nil {
			break
		}
		node = node.next
	}
	return out
}

func (l *List) Reverse() *List {
	out := List{}
	in := l.Array()
	slices.Reverse(in)
	for i := 0; i < len(in); i++ {
		out.Push(in[i])
	}
	return &out
}
