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

// ----------------------------------------------------------------------------
// Traversal Methods

// Fold a function f over the tree preorder.
//
// This method is a wrapper for PreorderTraversalFold(tree.root, initialAccumulator, f)
func TreePreorderTraversalFold[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return PreorderTraversalFold(tree.root, initialAccumulator, f)
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Preorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func PreorderTraversalFold[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	currentAccumulator = f(node.item, currentAccumulator)
	if node.left != nil {
		currentAccumulator = PreorderTraversalFold(node.left, currentAccumulator, f)
	}
	if node.right != nil {
		currentAccumulator = PreorderTraversalFold(node.right, currentAccumulator, f)
	}

	return currentAccumulator
}

// Fold a function f over the tree Inorder.
//
// This method is a wrapper for InorderTraversalFold(tree.root, initialAccumulator, f)
func TreeInorderTraversalFold[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return InorderTraversalFold(tree.root, initialAccumulator, f)
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Inorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func InorderTraversalFold[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if node.left != nil {
		currentAccumulator = InorderTraversalFold(node.left, currentAccumulator, f)
	}
	currentAccumulator = f(node.item, currentAccumulator)
	if node.right != nil {
		currentAccumulator = InorderTraversalFold(node.right, currentAccumulator, f)
	}

	return currentAccumulator
}

