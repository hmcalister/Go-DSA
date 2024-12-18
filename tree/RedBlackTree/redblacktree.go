package redblacktree

import (
	"iter"

	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

// Implement a red black tree.
//
// Like a binary search tree, items stored in nodes, such that all left/right children are respectively smaller/larger than the parent node.
// Unlike a binary search tree, which may become unbalanced and reduce to a linked list, a red black tree will rebalance itself
// after each addition or removal such that the maximum height is bounded by the log of the number of items.
// (BST's are worst case bounded linearly by the number of items).
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

// A helper method to transplant two nodes, such that old is replaced by new.
// (oldNode is removed from the tree).
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
	// Use same notation as diagram

	G := node
	P := node.left

	if P == nil {
		return ErrorRotationNotPossible
	}

	tree.replaceNode(G, P)

	// Fix pointers of 3
	G.left = P.right
	if G.left != nil {
		G.left.parent = G
	}

	// Fix pointers between G and P
	P.right = G
	G.parent = P

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
	// Use same notation as diagram

	G := node
	P := node.right

	if P == nil {
		return ErrorRotationNotPossible
	}

	tree.replaceNode(G, P)

	// Fix pointers of 3
	G.right = P.left
	if G.right != nil {
		G.right.parent = G
	}

	// Fix pointers between G and P
	P.left = G
	G.parent = P

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
// If the item is not present, nil is returned along with a dsa_error.ErrorItemNotFound.
func (tree *RedBlackTree[T]) Find(item T) (*RedBlackTreeNode[T], error) {
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
func (tree *RedBlackTree[T]) Items() []T {
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
// Apply should not change the item in a Node, as this could affect the tree structure.
//
// This method is a wrapper for ApplyNodePreorder(tree.root, f)
func ApplyTreePreorder[T any](tree *RedBlackTree[T], f func(item T)) {
	if tree.root == nil {
		return
	}
	ApplyNodePreorder(tree.root, f)
}

// Apply a function f to each node in a tree Inorder.
//
// Idiomatic Go should likely use IteratorTreeInorder() rather than functional methods.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
//
// This method is a wrapper for ApplyNodeInorder(tree.root, f)
func ApplyTreeInorder[T any](tree *RedBlackTree[T], f func(item T)) {
	if tree.root == nil {
		return
	}
	ApplyNodeInorder(tree.root, f)
}

// Apply a function f to each node in a tree Postorder.
//
// Idiomatic Go should likely use IteratorTreePostorder() rather than functional methods.
//
// Apply should not change the item in a Node, as this could affect the tree structure.
//
// This method is a wrapper for ApplyNodePostorder(tree.root, f)
func ApplyTreePostorder[T any](tree *RedBlackTree[T], f func(item T)) {
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
func FoldTreePreorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
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
func FoldTreeInorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
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
func FoldTreePostorder[T, G any](tree *RedBlackTree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
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
func IteratorTreePreorder[T any](tree *RedBlackTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodePreorder(tree.root)
}

// Iterate over the tree Inorder.
//
// This method is a wrapper for IteratorNodeInorder(tree.root)
func IteratorTreeInorder[T any](tree *RedBlackTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodeInorder(tree.root)
}

// Iterate over the tree Postorder.
//
// This method is a wrapper for IteratorNodePostorder(tree.root)
func IteratorTreePostorder[T any](tree *RedBlackTree[T]) iter.Seq[T] {
	if tree.root == nil {
		return func(yield func(T) bool) {}
	}
	return IteratorNodePostorder(tree.root)
}

// ----------------------------------------------------------------------------
// Add Methods

func (tree *RedBlackTree[T]) addCase1(node *RedBlackTreeNode[T]) {
	if node.parent == nil {
		node.color = color_BLACK
	} else {
		tree.addCase2(node)
	}
}

func (tree *RedBlackTree[T]) addCase2(node *RedBlackTreeNode[T]) {
	if getNodeColor(node.parent) == color_BLACK {
		return
	}

	tree.addCase3(node)
}

func (tree *RedBlackTree[T]) addCase3(node *RedBlackTreeNode[T]) {
	uncle := node.parent.getSibling()
	if getNodeColor(uncle) == color_RED {
		node.parent.color = color_BLACK
		uncle.color = color_BLACK
		node.parent.parent.color = color_RED
		tree.addCase1(node.parent.parent)
	} else {
		tree.addCase4(node)
	}
}

func (tree *RedBlackTree[T]) addCase4(node *RedBlackTreeNode[T]) {
	grandparent := node.parent.parent
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.addCase5(node)
}

func (tree *RedBlackTree[T]) addCase5(node *RedBlackTreeNode[T]) {
	grandparent := node.parent.parent
	grandparent.color = color_RED
	node.parent.color = color_BLACK
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

// Insert a new item into the tree.
//
// Returns a dsa_error.ErrorItemAlreadyPresent error if the item already exists in the tree.
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
			return dsa_error.ErrorItemAlreadyPresent
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

	// Fix up the tree
	tree.addCase1(newNode)

	// Account for the new nodes size and height
	currentNode = newNode.parent
	for currentNode != nil {
		currentNode.fixSize()
		currentNode.fixHeight()
		currentNode = currentNode.parent
	}

	return nil
}

// ----------------------------------------------------------------------------
// Remove Methods

func (tree *RedBlackTree[T]) removeCase1(node *RedBlackTreeNode[T]) {
	if node.parent == nil {
		return
	}
	tree.removeCase2(node)
}

func (tree *RedBlackTree[T]) removeCase2(node *RedBlackTreeNode[T]) {
	siblingNode := node.getSibling()
	if getNodeColor(siblingNode) == color_RED {
		node.parent.color = color_RED
		siblingNode.color = color_BLACK
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.removeCase3(node)
}

func (tree *RedBlackTree[T]) removeCase3(node *RedBlackTreeNode[T]) {
	siblingNode := node.getSibling()
	if getNodeColor(node.parent) == color_BLACK &&
		getNodeColor(siblingNode) == color_BLACK &&
		getNodeColor(siblingNode.left) == color_BLACK &&
		getNodeColor(siblingNode.right) == color_BLACK {
		siblingNode.color = color_RED
		tree.removeCase1(node.parent)
	} else {
		tree.removeCase4(node)
	}
}

func (tree *RedBlackTree[T]) removeCase4(node *RedBlackTreeNode[T]) {
	siblingNode := node.getSibling()
	if getNodeColor(node.parent) == color_RED &&
		getNodeColor(siblingNode) == color_BLACK &&
		getNodeColor(siblingNode.left) == color_BLACK &&
		getNodeColor(siblingNode.right) == color_BLACK {
		siblingNode.color = color_RED
		node.parent.color = color_BLACK
	} else {
		tree.removeCase5(node)
	}
}

func (tree *RedBlackTree[T]) removeCase5(node *RedBlackTreeNode[T]) {
	siblingNode := node.getSibling()
	if node == node.parent.left &&
		getNodeColor(siblingNode) == color_BLACK &&
		getNodeColor(siblingNode.left) == color_RED &&
		getNodeColor(siblingNode.right) == color_BLACK {
		siblingNode.color = color_RED
		siblingNode.left.color = color_BLACK
		tree.rotateRight(siblingNode)
	} else if node == node.parent.right &&
		getNodeColor(siblingNode) == color_BLACK &&
		getNodeColor(siblingNode.left) == color_BLACK &&
		getNodeColor(siblingNode.right) == color_RED {
		siblingNode.color = color_RED
		siblingNode.right.color = color_BLACK
		tree.rotateLeft(siblingNode)
	}
	tree.removeCase6(node)
}

func (tree *RedBlackTree[T]) removeCase6(node *RedBlackTreeNode[T]) {
	siblingNode := node.getSibling()
	siblingNode.color = getNodeColor(node.parent)
	node.parent.color = color_BLACK
	if node == node.parent.left && getNodeColor(siblingNode.right) == color_RED {
		siblingNode.right.color = color_BLACK
		tree.rotateLeft(node.parent)
	} else if getNodeColor(siblingNode.left) == color_RED {
		siblingNode.left.color = color_BLACK
		tree.rotateRight(node.parent)
	}
}

// Remove an item from the tree.
//
// Returns a dsa_error.ErrorItemNotFound if the item is not in the tree.
func (tree *RedBlackTree[T]) Remove(item T) error {
	currentNode, err := tree.Find(item)

	// If item not in tree, return that as error
	if err != nil {
		return err
	}

	// If we have two children, replace with successor
	if currentNode.left != nil && currentNode.right != nil {
		successor := currentNode.Successor()
		currentNode.item, successor.item = successor.item, currentNode.item
		currentNode = successor
	}

	// Get the child (if it exists)
	// If node has NO children, childNode remains nil (and that's okay!)
	var childNode *RedBlackTreeNode[T]
	if currentNode.left != nil && currentNode.right == nil {
		childNode = currentNode.left
	} else if currentNode.left == nil && currentNode.right != nil {
		childNode = currentNode.right
	} else {
		childNode = nil
	}

	parentNode := currentNode.parent

	if currentNode.color == color_BLACK {
		currentNode.color = getNodeColor(childNode)
		tree.removeCase1(currentNode)
	}

	tree.replaceNode(currentNode, childNode)
	if currentNode.parent == nil && childNode != nil {
		childNode.color = color_BLACK
	}

	for parentNode != nil {
		parentNode.fixSize()
		parentNode.fixHeight()
		parentNode = parentNode.parent
	}

	return nil
}
