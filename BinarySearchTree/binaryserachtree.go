package binarysearchtree

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

type BinarySearchTree[T any] struct {
	// The root of the tree
	root *BinarySearchTreeNode[T]

	// Comparator function to compare and order the type T
	comparatorFunction comparator.ComparatorFunction[T]
}

}
