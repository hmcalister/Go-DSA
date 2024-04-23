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

// Create a new binary tree of generic type.
//
// The comparator function (see github.com/hmcalister/Go-DSA/Comparator) defines how the items are ordered when creating the tree.
// This allows for trees that have any type, rather than just comparable types.
func New[T any](comparatorFunction comparator.ComparatorFunction[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root:               nil,
		comparatorFunction: comparatorFunction,
	}
}

// Create a new tree wrapper around a Node, allowing for a "subtree"
func NewTreeFromNode[T any](rootNode *BinarySearchTreeNode[T], comparatorFunction comparator.ComparatorFunction[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root:               rootNode,
		comparatorFunction: comparatorFunction,
	}
}

// Get the root the binary search tree
func (tree *BinarySearchTree[T]) Root() *BinarySearchTreeNode[T] {
	return tree.root
}

