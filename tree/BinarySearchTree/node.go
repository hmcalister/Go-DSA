package binarysearchtree

import (
	"iter"
)

type BinarySearchTreeNode[T any] struct {
	// The item of this node
	item T

	// The size of this node, the number of nodes in this subtree
	// (count of this node and all children)
	size int

	// The height of this node, the number of steps to the furthest leaf node
	height int

	// The parent of this node
	parent *BinarySearchTreeNode[T]

	// The left child of this node
	left *BinarySearchTreeNode[T]

	// The right child of this node
	right *BinarySearchTreeNode[T]
}

// Create a new node from an item.
func newNode[T any](item T) *BinarySearchTreeNode[T] {
	return &BinarySearchTreeNode[T]{
		item:   item,
		size:   1,
		height: 0,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

// Get the item of this tree node
//
// BEWARE: Mutating this item (e.g. if this item is a struct, array, etc...) may break the tree structure!
// Only mutate the result of node.Item() if:
// i) The type of T is a primitive, such as int, float... in which case the result is copied anyway
// ii) You can ensure your mutation will not change the ordering based on the tree's ComparatorFunction
func (node *BinarySearchTreeNode[T]) Item() T {
	return node.item
}

// Get the size of this Node, the number of items in the subtree rooted at this node
//
// A leaf node has size 1.
func (node *BinarySearchTreeNode[T]) Size() int {
	return node.size
}

// Get the height of this node, the number of steps from this node to the furthest leaf node.
//
// A leaf node has height 0.
func (node *BinarySearchTreeNode[T]) Height() int {
	return node.height
}

// Get the parent of this node. May be nil
//
// The root node has a nil parent.
func (node *BinarySearchTreeNode[T]) Parent() *BinarySearchTreeNode[T] {
	return node.parent
}

// Get the left child of this node. May be nil.
func (node *BinarySearchTreeNode[T]) Left() *BinarySearchTreeNode[T] {
	return node.left
}

// Get the right child of this node. May be nil.
func (node *BinarySearchTreeNode[T]) Right() *BinarySearchTreeNode[T] {
	return node.right
}

// ----------------------------------------------------------------------------
// Successor and Predecessor methods

// Return the successor of this node, or nil if there is no successor
func (node *BinarySearchTreeNode[T]) Successor() *BinarySearchTreeNode[T] {
	// If node has a right child, successor is one right then as far left as possible
	if node.right != nil {
		successorNode := node.right
		for successorNode.left != nil {
			successorNode = successorNode.left
		}
		return successorNode
	}

	// If node has a parent, find the first parent that is a left child and return that
	// If no such parent exists, this node has no successor

	currentNode := node.parent
	for currentNode != nil {
		parentNode := currentNode.parent

		// If the current parent is not nil and the current node is a left child, we are done
		if parentNode != nil && parentNode.left == currentNode {
			return currentNode
		}

		// Otherwise, continue to step up the tree
		currentNode = parentNode
	}

	// We did not find a parent node that was a left child
	return nil
}

// Return the predecessor of this node, or nil if there is no successor
func (node *BinarySearchTreeNode[T]) Predecessor() *BinarySearchTreeNode[T] {
	// If node has a left child, successor is one left then as far right as possible
	if node.left != nil {
		predecessorNode := node.left
		for predecessorNode.right != nil {
			predecessorNode = predecessorNode.right
		}
		return predecessorNode
	}

	// If node has a parent, find the first parent that is a right child and return that
	// If no such parent exists, this node has no predecessor

	currentNode := node.parent
	for currentNode != nil {
		parentNode := currentNode.parent

		// If the current parent is not nil and the current node is a right child, we are done
		if parentNode != nil && parentNode.right == currentNode {
			return currentNode
		}

		// Otherwise, continue to step up the tree
		currentNode = parentNode
	}

	// We did not find a parent node that was a right child
	return nil
}

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodePreorder[T any](node *BinarySearchTreeNode[T], f func(item T)) {
	f(node.item)
	if node.left != nil {
		ApplyNodePreorder(node.left, f)
	}
	if node.right != nil {
		ApplyNodePreorder(node.right, f)
	}
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodeInorder[T any](node *BinarySearchTreeNode[T], f func(item T)) {
	if node.left != nil {
		ApplyNodeInorder(node.left, f)
	}
	f(node.item)
	if node.right != nil {
		ApplyNodeInorder(node.right, f)
	}
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodePostorder[T any](node *BinarySearchTreeNode[T], f func(item T)) {
	if node.left != nil {
		ApplyNodePostorder(node.left, f)
	}
	if node.right != nil {
		ApplyNodePostorder(node.right, f)
	}
	f(node.item)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f (taking the current node item and the accumulator value) across the tree Preorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodePreorder[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	currentAccumulator = f(node.item, currentAccumulator)
	if node.left != nil {
		currentAccumulator = FoldNodePreorder(node.left, currentAccumulator, f)
	}
	if node.right != nil {
		currentAccumulator = FoldNodePreorder(node.right, currentAccumulator, f)
	}

	return currentAccumulator
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Inorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodeInorder[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if node.left != nil {
		currentAccumulator = FoldNodeInorder(node.left, currentAccumulator, f)
	}
	currentAccumulator = f(node.item, currentAccumulator)
	if node.right != nil {
		currentAccumulator = FoldNodeInorder(node.right, currentAccumulator, f)
	}

	return currentAccumulator
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Postorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodePostorder[T, G any](node *BinarySearchTreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if node.left != nil {
		currentAccumulator = FoldNodeInorder(node.left, currentAccumulator, f)
	}
	if node.right != nil {
		currentAccumulator = FoldNodeInorder(node.right, currentAccumulator, f)
	}
	currentAccumulator = f(node.item, currentAccumulator)

	return currentAccumulator
}

// ----------------------------------------------------------------------------
// Iterator Methods

// Iterate over each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func IteratorNodePreorder[T any](node *BinarySearchTreeNode[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if node == nil {
			return
		}

		if !yield(node.item) {
			return
		}
		for item := range IteratorNodePreorder(node.left) {
			if !yield(item) {
				return
			}
		}
		for item := range IteratorNodePreorder(node.right) {
			if !yield(item) {
				return
			}
		}
	}
}

// Iterate over each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func IteratorNodeInorder[T any](node *BinarySearchTreeNode[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if node == nil {
			return
		}

		for item := range IteratorNodeInorder(node.left) {
			if !yield(item) {
				return
			}
		}
		if !yield(node.item) {
			return
		}
		for item := range IteratorNodeInorder(node.right) {
			if !yield(item) {
				return
			}
		}
	}
}

// Iterate over each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func IteratorNodePostorder[T any](node *BinarySearchTreeNode[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if node == nil {
			return
		}

		for item := range IteratorNodePostorder(node.left) {
			if !yield(item) {
				return
			}
		}
		for item := range IteratorNodePostorder(node.right) {
			if !yield(item) {
				return
			}
		}
		if !yield(node.item) {
			return
		}
	}
}
