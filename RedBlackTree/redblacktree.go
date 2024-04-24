package redblacktree

import comparator "github.com/hmcalister/Go-DSA/Comparator"

type RedBlackTree[T any] struct {
	// The root of the tree
	root *RedBlackTreeNode[T]

	// Comparator function to compare and order the type T
	comparatorFunction comparator.ComparatorFunction[T]
}
