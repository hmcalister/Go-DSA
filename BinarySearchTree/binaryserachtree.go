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



	}

}

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for PreorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *BinarySearchTree[T]) ApplyTreePreorder(f func(item T)) {
	tree.root.ApplyPreorder(f)
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for InorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *BinarySearchTree[T]) ApplyTreeInorder(f func(item T)) {
	tree.root.ApplyInorder(f)
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for PostorderTraversalFold(tree.root, initialAccumulator, f)
func (tree *BinarySearchTree[T]) ApplyTreePostorder(f func(item T)) {
	tree.root.ApplyPostorder(f)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f over the tree preorder.
//
// This method is a wrapper for FoldPreorder(tree.root, initialAccumulator, f)
func FoldTreePreorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldPreorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Inorder.
//
// This method is a wrapper for FoldInorder(tree.root, initialAccumulator, f)
func FoldTreeInorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldInorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Postorder.
//
// This method is a wrapper for FoldPostorder(tree.root, initialAccumulator, f)
func FoldTreePostorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return FoldPostorder(tree.root, initialAccumulator, f)
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

// ----------------------------------------------------------------------------
// Remove Methods

// Remove an item from the tree.
//
// Returns an error if the item is not in the tree
func (tree *BinarySearchTree[T]) Remove(item T) error {
	// If the tree is empty, we cannot find the item
	if tree.root == nil {
		return &ItemNotFoundError[T]{item}
	}

	// Ensure we are not deleting the root
	if tree.comparatorFunction(item, tree.root.item) == 0 {
		// Special case when deleting the root! We cannot set the parent pointers

		currentNode := tree.root
		// If we have no children, simply remove the root
		if currentNode.left == nil && currentNode.right == nil {
			tree.root = nil
			return nil
		}

		// If we ONLY have a left child, we can replace root with the left child and be done
		if currentNode.left != nil && currentNode.right == nil {
			tree.root = currentNode.left
			tree.root.left = nil
			tree.root.right = nil
			return nil
		}

		// If we ONLY have a right child, we can replace root with the right child and be done
		if currentNode.left == nil && currentNode.right != nil {
			tree.root = currentNode.right
			tree.root.left = nil
			tree.root.right = nil
			return nil
		}

		// If we have two children we can simply replace with the successor (Case 3 below) and continue
		successorNode := currentNode.right
		for successorNode.left != nil {
			successorNode = successorNode.left
		}
		currentNode.item, successorNode.item = successorNode.item, currentNode.item
		// And now our target element for removal is present at the successor, which is not the root, and thus can be handled by the remaining logic
	}

	// We are now free to traverse the tree looking for item until
	// either we find it or find the leaf the item *would* be in

	// We know tree.root is not nil, so we can look at its children freely

	var currentCompare int
	currentNode := tree.root
	for {
		// If we have found a nil node, it means we have reached a leaf without encountering item
		// Therefore, the item is not in the tree, and hence cannot be removed
		if currentNode == nil {
			return &ItemNotFoundError[T]{item}
		}

		// We now know currentNode is not nil. The options are:
		// 1) currentNode has the item we want
		// 		So we break from the loop and process the deletion
		// 2) currentNode has an item larger than the target
		// 		So we loop again, looking at the left child
		// 3) currentNode has an item smaller than the target
		// 		So we loop again, looking at the left child

		currentCompare = tree.comparatorFunction(item, currentNode.item)
		if currentCompare == 0 {
			break
		}
		if currentCompare < 0 {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// Now, we know that currentNode has the item we are interested in
	// We now get to remove the node based on three conditions:
	// 1) If currentNode is a leaf, we can simply remove it
	// 2) If currentNode has one child, we can splice it out
	// 3) If currentNode has two children, we replace it with
	// 		the successor then delete the successor (which is now a leaf and hence case 1)

	// Function to be called on the parent of the removed node to fix up size and height
	fixupFunc := func(node *BinarySearchTreeNode[T]) {
		for node != nil {
			node.size += 1

			// The left and right height default to -1, in case the left or right child are nil
			leftHeight := -1
			if node.left != nil {
				leftHeight = node.left.height
			}
			rightHeight := -1
			if node.right != nil {
				rightHeight = node.right.height
			}
			node.height = max(leftHeight, rightHeight) + 1
			node = node.parent
		}
	}

	// Case 2 when we have only a left child
	if currentNode.left != nil && currentNode.right == nil {
		defer fixupFunc(currentNode.parent)
		// Since we know this cannot be the root node, we can access parent
		parentNode := currentNode.parent
		childNode := currentNode.left

		// Fix up pointer of parent to splice out currentNode
		if currentCompare < 0 {
			// We are a left child
			parentNode.left = childNode
		} else {
			// We are a right child
			parentNode.right = childNode
		}

		// Fix up pointer of child to finish splice
		childNode.parent = parentNode

		// Set the pointers of this node to nil to ensure no bugs creep in
		currentNode.parent = nil
		currentNode.right = nil
		currentNode.left = nil

		return nil
	}

	// Case 2 when we have only a right child
	if currentNode.left != nil && currentNode.right == nil {
		defer fixupFunc(currentNode.parent)
		// Since we know this cannot be the root node, we can access parent
		parentNode := currentNode.parent
		childNode := currentNode.right

		// Fix up pointer of parent to splice out currentNode
		if currentCompare < 0 {
			// We are a left child
			parentNode.left = childNode
		} else {
			// We are a right child
			parentNode.right = childNode
		}

		// Fix up pointer of child to finish splice
		childNode.parent = parentNode

		// Set the pointers of this node to nil to ensure no bugs creep in
		currentNode.parent = nil
		currentNode.right = nil
		currentNode.left = nil

		return nil
	}

	// Case 3: Reduce to Case 1
	// Once this statement is done currentNode will be the child node to delete

	if currentNode.left != nil && currentNode.right != nil {
		successorNode := currentNode.right

		for successorNode.left != nil {
			successorNode = successorNode.left
		}

		currentNode.item, successorNode.item = successorNode.item, currentNode.item
		currentNode = successorNode
	}

	// And now we must be in Case 1

	parentNode := currentNode.parent
	// Let's make sure we know if we are a left or right child
	if tree.comparatorFunction(item, parentNode.left.item) == 0 {
		// Remove the item by nil-ing the left pointer
		parentNode.left = nil
	} else {
		parentNode.right = nil
	}
	// Finish by removing the pointers of this node to avoid bugs
	currentNode.parent = nil
	currentNode.left = nil
	currentNode.right = nil

	return nil
}
