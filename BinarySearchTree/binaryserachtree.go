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

// Fold a function f over the tree Postorder.
//
// This method is a wrapper for PostorderTraversalFold(tree.root, initialAccumulator, f)
func TreePostorderTraversalFold[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return PostorderTraversalFold(tree.root, initialAccumulator, f)
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Postorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func PostorderTraversalFold[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if node.left != nil {
		currentAccumulator = InorderTraversalFold(node.left, currentAccumulator, f)
	}
	if node.right != nil {
		currentAccumulator = InorderTraversalFold(node.right, currentAccumulator, f)
	}
	currentAccumulator = f(node.item, currentAccumulator)

	return currentAccumulator
}

// ----------------------------------------------------------------------------
// Add Methods

// Insert a new item into the tree.
//
// Returns an error if the item already exists in the tree.
func (tree *BinarySearchTree[T]) Add(item T) error {
	// If the tree is currently empty, add this item to the root
	if tree.root == nil {
		tree.root = newNode(item)
	}

	// We know the tree is nonempty,
	// so we can now access root and walk the tree until this items position is found

	var currentCompare int
	var parentNode *BinarySearchTreeNode[T]
	currentNode := tree.root

	// We walk the tree until the node we are looking at is nil, i.e. we have reached a spot for this item
	// at which point parentNode will be the parent node we can access to make a new child
	for currentNode != nil {
		parentNode = currentNode
		currentCompare = tree.comparatorFunction(item, currentNode.item)

		// If the item we are inserting is the same as this node, we reject it and return an error
		if currentCompare == 0 {
			return &ItemAlreadyPresentError[T]{item}
		}

		// Otherwise, we can walk to this node's left or right child based on currentCompare
		// If currentCompare < 0, currentNode's item is *larger* than item so we walk left
		// If currentCompare > 0, currentNode's item is *smaller* than item so we walk right
		if currentCompare < 0 {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// We have no reached a nil node, meaning parentNode was the last non-nil node
	// (and since we ensured that the tree was not empty, parentNode is definitely non-nil)
	// The value of currentCompare will tell us if we are adding to currentNode's left or right

	newNode := newNode(item)
	newNode.parent = parentNode
	if currentCompare < 0 {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}

	// We now fix the size and height of each node by walking from parentNode up the tree
	// incrementing size at each step (to account for the newly added node)
	// and recomputing the height
	for parentNode != nil {
		parentNode.size += 1

		// The left and right height default to -1, in case the left or right child are nil
		leftHeight := -1
		if parentNode.left != nil {
			leftHeight = parentNode.left.height
		}
		rightHeight := -1
		if parentNode.right != nil {
			rightHeight = parentNode.right.height
		}
		parentNode.height = max(leftHeight, rightHeight) + 1
		parentNode = parentNode.parent
	}

	return nil
}

