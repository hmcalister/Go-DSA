package redblacktree

import comparator "github.com/hmcalister/Go-DSA/Comparator"

type RedBlackTree[T any] struct {
	// The root of the tree
	root *RedBlackTreeNode[T]

	// Comparator function to compare and order the type T
	comparatorFunction comparator.ComparatorFunction[T]
}

// Create a new red-black tree of generic type.
//
// The comparator function (see github.com/hmcalister/Go-DSA/Comparator) defines how the items are ordered when creating the tree.
// This allows for trees that have any type, rather than just comparable types.
func New[T any](comparatorFunction comparator.ComparatorFunction[T]) *RedBlackTree[T] {
	return &RedBlackTree[T]{
		root:               nil,
		comparatorFunction: comparatorFunction,
	}
}

// Get the root the red-black search tree
func (tree *RedBlackTree[T]) Root() *RedBlackTreeNode[T] {
	return tree.root
}

// ----------------------------------------------------------------------------
// Find Methods

// Determines if a given item is present in the tree.
// If the item is present in the tree, the Node containing that item is returned with nil error.
// If the item is not present, nil is returned along with an error.
func (tree *RedBlackTree[T]) Find(item T) (*RedBlackTreeNode[T], error) {
	// If the root is nil, the item cannot be in the tree
	if tree.root == nil {
		return nil, &ItemNotFoundError[T]{item}
	}

	// Now we know the root is non-nil we can start traversing the tree

	currentNode := tree.root
	for currentNode != nil {
		currentCompare := tree.comparatorFunction(item, currentNode.item)

		if currentCompare == 0 {
			return currentNode, nil
		}

		if currentCompare < 0 {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// If we exit the loop, that means we have reached a leaf without finding the item
	return nil, &ItemNotFoundError[T]{item}
}
