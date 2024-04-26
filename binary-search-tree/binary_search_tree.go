package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if bst.data < i {
		if bst.right == nil {
			bst.right = &BinarySearchTree{data: i}
		} else {
			bst.right.Insert(i)
		}
	} else {
		if bst.left == nil {
			bst.left = &BinarySearchTree{data: i}
		} else {
			bst.left.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	if bst == nil {
		return []int{}
	}
	return append(append(bst.left.SortedData(), bst.data), bst.right.SortedData()...)
}
