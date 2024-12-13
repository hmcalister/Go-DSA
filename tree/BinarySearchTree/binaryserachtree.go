package binarysearchtree

import (
	"iter"

	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

// Implement a binary search tree.
//
// A BST has items stored in nodes, such that all left/right children are respectively smaller/larger than the parent node.
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

// Get the root the binary search tree
func (tree *BinarySearchTree[T]) Root() *BinarySearchTreeNode[T] {
	return tree.root
}

// Find Methods

// Determines if a given item is present in the tree.
// If the item is present in the tree, the Node containing that item is returned with nil error.
// If the item is not present, nil is returned along with a dsa_error.ErrorItemNotFound.
func (tree *BinarySearchTree[T]) Find(item T) (*BinarySearchTreeNode[T], error) {
	// If the root is nil, the item cannot be in the tree
	if tree.root == nil {
		return nil, dsa_error.ErrorItemNotFound
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
	return nil, dsa_error.ErrorItemNotFound
}

// Get all items from the tree. This method allocates an array of length equal to the number of items.
// Items may not be present in the order they were inserted.
func (tree *BinarySearchTree[T]) Items() []T {
	items := make([]T, tree.root.size)
	ApplyTreeInorder(tree, func(item T) { items = append(items, item) })
	return items
}

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Idiomatic Go should likely use IteratorTreePreorder() rather than functional methods.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
//
// This method is a wrapper for ApplyNodePreorder(tree.root, f)
func ApplyTreePreorder[T any](tree *BinarySearchTree[T], f func(item T)) {
	if tree.root == nil {
		return
	}
	ApplyNodePreorder(tree.root, f)
}

// Apply a function f to each node in a tree Inorder.
//
// Idiomatic Go should likely use IteratorTreeInorder() rather than functional methods.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
//
// This method is a wrapper for ApplyNodeInorder(tree.root, f)
func ApplyTreeInorder[T any](tree *BinarySearchTree[T], f func(item T)) {
	if tree.root == nil {
		return
	}
	ApplyNodeInorder(tree.root, f)
}

// Apply a function f to each node in a tree Postorder.
//
// Idiomatic Go should likely use IteratorTreePostorder() rather than functional methods.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
//
// This method is a wrapper for ApplyNodePostorder(tree.root, f)
func ApplyTreePostorder[T any](tree *BinarySearchTree[T], f func(item T)) {
	if tree.root == nil {
		return
	}
	ApplyNodePostorder(tree.root, f)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f over the tree preorder.
//
// Idiomatic Go should likely use IteratorTreePreorder() rather than functional methods.
//
// This method is a wrapper for FoldNodePreorder(tree.root, initialAccumulator, f)
func FoldTreePreorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodePreorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Inorder.
//
// Idiomatic Go should likely use IteratorTreeInorder() rather than functional methods.
//
// This method is a wrapper for FoldNodeInorder(tree.root, initialAccumulator, f)
func FoldTreeInorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodeInorder(tree.root, initialAccumulator, f)
}

// Fold a function f over the tree Postorder.
//
// Idiomatic Go should likely use IteratorTreePostorder() rather than functional methods.
//
// This method is a wrapper for FoldNodePostorder(tree.root, initialAccumulator, f)
func FoldTreePostorder[T, G any](tree *BinarySearchTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	if tree.root == nil {
		return initialAccumulator
	}
	return FoldNodePostorder(tree.root, initialAccumulator, f)
}

// ----------------------------------------------------------------------------
// Iterator Methods

// Iterate over the tree Preorder.
//
// This method is a wrapper for IteratorNodePreorder(tree.root)
func IteratorTreePreorder[T any](tree *BinarySearchTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodePreorder(tree.root)
}

// Iterate over the tree Inorder.
//
// This method is a wrapper for IteratorNodeInorder(tree.root)
func IteratorTreeInorder[T any](tree *BinarySearchTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodeInorder(tree.root)
}

// Iterate over the tree Postorder.
//
// This method is a wrapper for IteratorNodePostorder(tree.root)
func IteratorTreePostorder[T any](tree *BinarySearchTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodePostorder(tree.root)
}

// ----------------------------------------------------------------------------
// Add Methods

// Insert a new item into the tree.
//
// Returns a dsa_error.ItemAlreadyPresent error if the item already exists in the tree.
func (tree *BinarySearchTree[T]) Add(item T) error {
	// If the tree is currently empty, add this item to the root
	if tree.root == nil {
		tree.root = newNode(item)
		return nil
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
			return dsa_error.ErrorItemAlreadyPresent
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

// A helper method to fixup nodes from a deleted node up to the root.
func (tree *BinarySearchTree[T]) removeFixupHelper(node *BinarySearchTreeNode[T]) {
	for node != nil {
		node.size -= 1

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

// Remove an item from the tree.
//
// Returns a dsa_error.ErrorItemNotFound if item is not present in the tree.
func (tree *BinarySearchTree[T]) Remove(item T) error {
	// If the tree is empty, we cannot find the item
	if tree.root == nil {
		return dsa_error.ErrorItemNotFound
	}

	// If we are trying to delete the root node, we must handle it slightly differently since parent is nil
	if tree.comparatorFunction(item, tree.root.item) == 0 {
		tree.removeRoot()
		return nil
	}

	// Now we know the root is not the node to delete, we can traverse the tree.
	// We will traverse the tree to find the node with the item.
	// If we do not find a node with the item, we return an error instead

	currentNode := tree.root
	for currentNode != nil {
		currentComparison := tree.comparatorFunction(item, currentNode.item)

		// If we find the node with an item, we are done traversing!
		if currentComparison == 0 {
			break
		}

		if currentComparison < 0 {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	if currentNode == nil {
		return dsa_error.ErrorItemNotFound
	}

	// We now have the node to delete held in currentNode.
	// We have three cases depending on the state of currentNode:
	// 1) If currentNode is a leaf, i.e. has no children, just remove the node
	// 2) If the currentNode has only one child, replace this node by the single child
	// 3) If the currentNode has two children, replace this node by the successor and then delete that node

	// Case 1: No Children
	if currentNode.left == nil && currentNode.right == nil {
		// We know parent is non-nil as currentNode is not the root
		parentNode := currentNode.parent

		// Fix pointers
		if parentNode.left == currentNode {
			parentNode.left = nil
		} else {
			parentNode.right = nil
		}

		// nil current node to avoid bugs
		currentNode.parent = nil

		tree.removeFixupHelper(parentNode)
		return nil
	}

	// Case 2: One child

	if currentNode.left != nil && currentNode.right == nil {
		// We know parent is non-nil as currentNode is not the root
		parentNode := currentNode.parent
		childNode := currentNode.left

		// Fix pointers
		childNode.parent = parentNode
		if parentNode.left == currentNode {
			parentNode.left = childNode
		} else {
			parentNode.right = childNode
		}

		// nil current node to avoid bugs
		currentNode.parent = nil
		currentNode.left = nil

		tree.removeFixupHelper(parentNode)
		return nil
	}

	if currentNode.left == nil && currentNode.right != nil {
		// We know parent is non-nil as currentNode is not the root
		parentNode := currentNode.parent
		childNode := currentNode.right

		// Fix pointers
		childNode.parent = parentNode
		if parentNode.left == currentNode {
			parentNode.left = childNode
		} else {
			parentNode.right = childNode
		}

		// nil current node to avoid bugs
		currentNode.parent = nil
		currentNode.right = nil

		tree.removeFixupHelper(parentNode)
		return nil
	}

	// Case 3: Two children

	successorNode := currentNode.Successor()
	currentNode.item, successorNode.item = successorNode.item, currentNode.item
	// For consistency, rename successorNode to currentNode
	currentNode = successorNode

	// Now we have replaced the node with its successor, we can more easily delete currentNode.
	// We must either have a leaf node, or a node with only a right child.
	// If the successor had a left child, it would have been chosen as the successor.
	if currentNode.right != nil {
		// Replace currentNode with its right child and be done with it, like case 2

		// Note we can access parent because currentNode cannot be root
		parent := currentNode.parent
		child := currentNode.right

		// Fix pointers
		if parent.left == currentNode {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent

		// nil out current node to avoid bugs
		currentNode.parent = nil
		currentNode.left = nil
		currentNode.right = nil

		tree.removeFixupHelper(parent)
		return nil
	}

	// The successor is a leaf with no children, still easy

	// Note we can access parent because currentNode cannot be root
	parent := currentNode.parent

	// Fix pointers
	if parent.left == currentNode {
		parent.left = nil
	} else {
		parent.right = nil
	}

	// nil out current node to avoid bugs
	currentNode.parent = nil
	currentNode.left = nil
	currentNode.right = nil

	tree.removeFixupHelper(parent)

	return nil
}

// Remove the root node.
func (tree *BinarySearchTree[T]) removeRoot() {
	// if the root is the ONLY node simply remove it
	if tree.root.left == nil && tree.root.right == nil {
		tree.root = nil
		return
	}

	// if the root has only one child, replace the root with that child
	if tree.root.left != nil && tree.root.right == nil {
		// We have only a left child
		removedRoot := tree.root
		tree.root = tree.root.left

		tree.root.parent = nil
		removedRoot.left = nil
		removedRoot.right = nil
		return
	}
	if tree.root.left == nil && tree.root.right != nil {
		// We have only a left child
		removedRoot := tree.root
		tree.root = tree.root.right

		tree.root.parent = nil
		removedRoot.left = nil
		removedRoot.right = nil
		return
	}

	// If the root has two children, swap the item with the successor and delete that
	// Note that we must move down the tree, as here root must have a right child

	rootSuccessor := tree.root.Successor()
	tree.root.item, rootSuccessor.item = rootSuccessor.item, tree.root.item

	// To delete rootSuccessor we must either have a leaf node, or a node with only a right child
	// If the node had a left child, it would have been chosen as the successor.
	if rootSuccessor.right != nil {
		// Replace rootSuccessor with its right child and be done with it

		// Note we can access parent because rootSuccessor cannot be root
		parent := rootSuccessor.parent
		child := rootSuccessor.right

		if parent.left == rootSuccessor {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent

		rootSuccessor.parent = nil
		rootSuccessor.left = nil
		rootSuccessor.right = nil

		tree.removeFixupHelper(parent)
		return
	}

	// The rootSuccessor is a leaf with no children

	// Note we can access parent because rootSuccessor cannot be root
	parent := rootSuccessor.parent
	if parent.left == rootSuccessor {
		parent.left = nil
	} else {
		parent.right = nil
	}
	rootSuccessor.parent = nil
	rootSuccessor.left = nil
	rootSuccessor.right = nil

	tree.removeFixupHelper(parent)
}
