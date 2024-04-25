package redblacktree

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

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
// Misc / Helper methods

// Rotate right around the given node.
//
// Given the node G in the diagram:
//
//	        G
//	      /   \
//	     P     U
//	   /  \   /  \
//	  X   3  4   5
//	 / \
//	1   2
//
// Shift it into the form:
//
//	       P
//	    /    \
//	   X      G
//	 /  \    /  \
//	1    2  3    U
//	            / \
//	           4   5
//
// Should NEVER be called on a node that has no left child.
// If rotate fails returns an error.
func (tree *RedBlackTree[T]) rotateRight(node *RedBlackTreeNode[T]) error {
	// Ensure rotation is possible
	if node.left == nil {
		return &RotationNotPossible[int]{}
	}

	// Use same notation as diagram

	G := node
	P := node.left

	// Fix parent connection (not shown on the diagram) -----------------------

	// Fix pointer from G.parent down, will now point to P
	// Allow parent being nil in case node is root
	if G.parent != nil {
		tree.root = P
	} else {
		if G.parent.left == G {
			G.parent.left = P
		} else {
			G.parent.right = P
		}
	}

	// Fix pointer from P up
	P.parent = G.parent

	// Correct position of 3 --------------------------------------------------

	// Move child node's right to current node's left. (See where 3 goes in the diagram)
	// Note we cannot overwrite / lose currentNode.left, as it is help by childNode variable.
	// Also childNode.right may be null -- that's fine

	// Move 3 to be G's left, fixing G down to 3
	G.left = P.right

	// Fix 3 up to G if not nil
	if G.left != nil {
		G.left.parent = G
	}

	// Fix relationship between G and P ---------------------------------------

	// Fix G up
	G.parent = P

	// Fix P down
	P.right = G

	return nil
}

// Rotate right around the given node.
//
// Given the node G in the diagram:
//
//			G
//		  /	  \
//		U		P
//	  /	 \	   / \
//	 1	  2	  3   X
//				 / \
//				4   5
//
// Shift it into the form:
//
//				P
//			  /	  \
//			G		X
//		  /	 \	   / \
//		 U	  3	  4   5
//		/ \
//	   1   2
//
// Should NEVER be called on a node that has no left child.
// If rotate fails returns an error.
func (tree *RedBlackTree[T]) rotateLeft(node *RedBlackTreeNode[T]) error {
	// Ensure rotation is possible
	if node.right == nil {
		return &RotationNotPossible[int]{}
	}

	// Use same notation as diagram

	G := node
	P := node.right

	// Fix parent connection (not shown on the diagram) -----------------------

	// Fix pointer from G.parent down, will now point to P
	// Allow parent being nil in case node is root
	if G.parent != nil {
		tree.root = P
	} else {
		if G.parent.left == G {
			G.parent.left = P
		} else {
			G.parent.right = P
		}
	}

	// Fix pointer from P up
	P.parent = G.parent

	// Correct position of 3 --------------------------------------------------

	// Move child node's right to current node's left. (See where 3 goes in the diagram)
	// Note we cannot overwrite / lose currentNode.left, as it is help by childNode variable.
	// Also childNode.right may be null -- that's fine

	// Move 3 to be G's left, fixing G down to 3
	G.right = P.left

	// Fix 3 up to G if not nil
	if G.right != nil {
		G.right.parent = G
	}

	// Fix relationship between G and P ---------------------------------------

	// Fix G up
	G.parent = P

	// Fix P down
	P.left = G

	return nil
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

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
// This method is a wrapper for PreorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *RedBlackTree[T]) ApplyTreePreorder(f func(item T)) {
	tree.root.ApplyNodePreorder(f)
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
// This method is a wrapper for InorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *RedBlackTree[T]) ApplyTreeInorder(f func(item T)) {
	tree.root.ApplyNodeInorder(f)
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
// This method is a wrapper for PostorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *RedBlackTree[T]) ApplyTreePostorder(f func(item T)) {
	tree.root.ApplyNodePostorder(f)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f over the tree preorder.
//
// This method is a wrapper for FoldPreorder(tree.root, initialAccumulator, f)
func FoldTreePreorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldNodePreorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Inorder.
//
// This method is a wrapper for FoldInorder(tree.root, initialAccumulator, f)
func FoldTreeInorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldNodeInorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Postorder.
//
// This method is a wrapper for FoldPostorder(tree.root, initialAccumulator, f)
func FoldTreePostorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldNodePostorder(tree.root, initialAccumulator, f)
}

// ----------------------------------------------------------------------------
// Add Methods

// Insert a new item into the tree.
//
// Returns an error if the item already exists in the tree.
func (tree *RedBlackTree[T]) Add(item T) error {
	return nil
}

// ----------------------------------------------------------------------------
// Remove Methods

// Remove an item from the tree.
//
// Returns an error if the item is not in the tree
func (tree *RedBlackTree[T]) Remove(item T) error {
	return nil
}
