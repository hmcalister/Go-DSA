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

// A helper method to transplant two nodes, such that old is replaced by new
// (oldNode is removed from the tree)
func (tree *RedBlackTree[T]) replaceNode(oldNode, newNode *RedBlackTreeNode[T]) {
	if oldNode.parent == nil {
		tree.root = newNode
	} else if oldNode == oldNode.parent.left {
		oldNode.parent.left = newNode
	} else {
		oldNode.parent.right = newNode
	}
	if newNode != nil {
		newNode.parent = oldNode.parent
	}
}

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
	if G.parent == nil {
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

	// Fix the size and heights -----------------------------------------------

	G.fixSize()
	G.fixHeight()
	P.fixSize()
	P.fixHeight()

	return nil
}

// Rotate right around the given node.
//
// Given the node G in the diagram:
//
//	      G
//	    /  \
//	  U      P
//	 / \    / \
//	1   2  3   X
//	          / \
//	         4   5
//
// Shift it into the form:
//
//	       P
//	     /  \
//	    G     X
//	   / \    / \
//	  U   3  4   5
//	 / \
//	1   2
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
	if G.parent == nil {
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

	// Fix the size and heights -----------------------------------------------

	G.fixSize()
	G.fixHeight()
	P.fixSize()
	P.fixHeight()

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
	if tree.root == nil {
		return
	}
	tree.root.ApplyNodePreorder(f)
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
// This method is a wrapper for InorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *RedBlackTree[T]) ApplyTreeInorder(f func(item T)) {
	if tree.root == nil {
		return
	}
	tree.root.ApplyNodeInorder(f)
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
// This method is a wrapper for PostorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *RedBlackTree[T]) ApplyTreePostorder(f func(item T)) {
	if tree.root == nil {
		return
	}
	tree.root.ApplyNodePostorder(f)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f over the tree preorder.
//
// This method is a wrapper for FoldPreorder(tree.root, initialAccumulator, f)
func FoldTreePreorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodePreorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Inorder.
//
// This method is a wrapper for FoldInorder(tree.root, initialAccumulator, f)
func FoldTreeInorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodeInorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Postorder.
//
// This method is a wrapper for FoldPostorder(tree.root, initialAccumulator, f)
func FoldTreePostorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodePostorder(tree.root, initialAccumulator, f)
}

// ----------------------------------------------------------------------------
// Add Methods

// Insert a new item into the tree.
//
// Returns an error if the item already exists in the tree.
func (tree *RedBlackTree[T]) Add(item T) error {
	// If the tree is empty we can simply add a new node as the root
	if tree.root == nil {
		tree.root = newNode(item)
		return nil
	}

	// Otherwise, the root exists so we must find where to insert this item
	// (or, if the item is in the tree, return an error)

	var traverseCompare int
	var parentNode *RedBlackTreeNode[T]
	currentNode := tree.root

	for currentNode != nil {
		parentNode = currentNode
		traverseCompare = tree.comparatorFunction(item, currentNode.item)

		// If the item is the same as the current node, return an error and do not insert
		if traverseCompare == 0 {
			return &ItemAlreadyPresentError[T]{item}
		}

		// Otherwise, we can walk to this node's left or right child based on currentCompare
		// If currentCompare < 0, currentNode's item is *larger* than item so we walk left
		// If currentCompare > 0, currentNode's item is *smaller* than item so we walk right
		if traverseCompare < 0 {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// We have now found a nil node, meaning parentNode was the last non-nil node
	// (and since we ensured that the tree was not empty, parentNode is definitely non-nil)
	// The value of currentCompare will tell us if we are adding to currentNode's left or right

	newNode := newNode(item)
	newNode.color = color_RED
	newNode.parent = parentNode
	if traverseCompare < 0 {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}

	// Account for the new nodes size and height
	currentNode = newNode.parent
	for currentNode != nil {
		currentNode.fixSize()
		currentNode.fixHeight()
		currentNode = currentNode.parent
	}

	// Finally, fix up the tree
	tree.addFix(newNode)

	return nil
}

// Fix the red black tree given the newly inserted leaf node newNode
//
// This method may restructure the tree, and will ensure size and height are fixed too.
//
// From https://www.programiz.com/dsa/red-black-tree
func (tree *RedBlackTree[T]) addFix(newNode *RedBlackTreeNode[T]) {
	currentNode := newNode

	for currentNode != tree.root && currentNode.parent.color == color_RED {
		// note because the root is ALWAYS black we can access currentNode.parent.parent
		// since currentNode.parent.color is red
		if currentNode.parent == currentNode.parent.parent.left {
			// parent is a left child
			uncleNode := currentNode.parent.parent.right
			if uncleNode != nil && uncleNode.color == color_RED {
				// Case I
				uncleNode.color = color_BLACK
				currentNode.parent.color = color_BLACK
				currentNode.parent.parent.color = color_RED
				currentNode = currentNode.parent.parent
			} else {
				if currentNode == currentNode.parent.right {
					// Case II
					currentNode = currentNode.parent
					tree.rotateLeft(currentNode)
				}
				// Case III
				currentNode.parent.color = color_BLACK
				currentNode.parent.parent.color = color_RED
				tree.rotateRight(currentNode.parent.parent)
			}
		} else {
			// parent is a right child
			uncleNode := currentNode.parent.parent.left
			if uncleNode != nil && uncleNode.color == color_RED {
				// Case I
				uncleNode.color = color_BLACK
				currentNode.parent.color = color_BLACK
				currentNode.parent.parent.color = color_RED
				currentNode = currentNode.parent.parent
			} else {
				if currentNode == currentNode.parent.left {
					// Case II
					currentNode = currentNode.parent
					tree.rotateRight(currentNode)
				}
				// Case III
				currentNode.parent.color = color_BLACK
				currentNode.parent.parent.color = color_RED
				tree.rotateLeft(currentNode.parent.parent)
			}
		}
	}
	tree.root.color = color_BLACK
}

// ----------------------------------------------------------------------------
// Remove Methods

// A helper method to transplant two nodes, such that old is replaced by new
// (oldNode is removed from the tree)
func (tree *RedBlackTree[T]) replaceNode(oldNode, newNode *RedBlackTreeNode[T]) {
	if oldNode.parent == nil {
		tree.root = newNode
	} else if oldNode == newNode.parent.left {
		oldNode.parent.left = newNode
	} else {
		oldNode.parent.right = newNode
	}
	if newNode != nil {
		newNode.parent = oldNode.parent
	}
}

// Remove an item from the tree.
//
// Returns an error if the item is not in the tree
func (tree *RedBlackTree[T]) Remove(item T) error {
	// Find the node to delete. If we cannot find the node, it does not exist in the tree

	Z, err := tree.Find(item)
	if err != nil {
		return err
	}

	Y := Z
	YorigColor := Y.color
	var X *RedBlackTreeNode[T]
	// If this node has only a left child, transplant it
	if Z.left != nil && Z.right == nil {
		X = Z.right
		tree.replaceNode(Z, X)
	} else if Z.left == nil && Z.right != nil {
		// Same for is we only have a right child

		X = Z.left
		tree.replaceNode(Z, X)
	} else {
		// We find the minimum of the right subtree (similar to successor)

		Y = Z.right
		for Y.left != nil {
			Y = Y.left
		}
		YorigColor = Y.color
		X = Y.right

		if Y.parent == Z {
			X.parent = Y
		} else {
			tree.replaceNode(Y, Y.right)
		}

		tree.replaceNode(Z, Y)
		Y.left = Z.left
		Y.left.parent = Y
		Y.color = Z.color
	}
	currentNode := X
	for currentNode != nil {
		currentNode.fixSize()
		currentNode.fixHeight()
		currentNode = currentNode.parent
	}

	if YorigColor == color_BLACK {
		tree.deleteFix(X)
	}

	return nil
}

// Helper method to fix the tree after deleting nodes
func (tree *RedBlackTree[T]) deleteFix(currentNode *RedBlackTreeNode[T]) {
	var siblingNode *RedBlackTreeNode[T]
	for currentNode != tree.root && currentNode.color == color_BLACK {
		if currentNode == currentNode.parent.left {
			siblingNode = currentNode.parent.right
			if siblingNode.color == color_RED {
				siblingNode.color = color_BLACK
				currentNode.parent.color = color_RED
				tree.rotateLeft(currentNode.parent)
				siblingNode = currentNode.parent.right
			}

			if siblingNode.left.color == color_BLACK && siblingNode.right.color == color_BLACK {
				siblingNode.color = color_RED
				currentNode = currentNode.parent
			} else {
				if siblingNode.right.color == color_BLACK {
					siblingNode.left.color = color_BLACK
					siblingNode.color = color_RED
					tree.rotateRight(siblingNode)
					siblingNode = currentNode.parent.right
				}

				siblingNode.color = currentNode.parent.color
				currentNode.parent.color = color_BLACK
				siblingNode.right.color = color_BLACK
				tree.rotateLeft(currentNode.parent)
				break
			}
		} else {
			siblingNode = currentNode.parent.left
			if siblingNode.color == color_RED {
				siblingNode.color = color_BLACK
				currentNode.parent.color = color_RED
				tree.rotateRight(currentNode.parent)
				siblingNode = currentNode.parent.left
			}

			if siblingNode.left.color == color_BLACK && siblingNode.right.color == color_BLACK {
				siblingNode.color = color_RED
				currentNode = currentNode.parent
			} else {
				if siblingNode.left.color == color_BLACK {
					siblingNode.right.color = color_BLACK
					siblingNode.color = color_RED
					tree.rotateLeft(siblingNode)
					siblingNode = currentNode.parent.left
				}

				siblingNode.color = currentNode.parent.color
				currentNode.parent.color = color_BLACK
				siblingNode.left.color = color_BLACK
				tree.rotateRight(currentNode.parent)
				break
			}
		}
	}
	tree.root.color = color_BLACK
}
