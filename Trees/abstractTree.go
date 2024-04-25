package abstracttree

type Tree[T any] interface {
	// Get the root node of this tree
	Root() *TreeNode[T]

	// Find the node holding this item in the tree
	// or an error if the item does not exist in the tree
	Find(item T) (*TreeNode[T], error)

	// Add an item to the tree,
	// returning an error if the item is already in the tree
	Add(item T) error

	// Remove an item from the tree,
	// returning an error if the item is not in the tree
	Remove(item T) error
}

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for PreorderTraversalFold(tree.root, initialAccumulator, f)
func ApplyTreePreorder[T any](tree *Tree[T], f func(item T)) {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return
	}
	ApplyNodePreorder(treeRoot, f)
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for InorderTraversalFold(tree.root, initialAccumulator, f)
func ApplyTreeInorder[T any](tree *Tree[T], f func(item T)) {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return
	}
	ApplyNodeInorder(treeRoot, f)
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
// This method is a wrapper for PostorderTraversalFold(tree.root, initialAccumulator, f)
func ApplyTreePostorder[T any](tree *Tree[T], f func(item T)) {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return
	}
	ApplyNodePostorder(treeRoot, f)
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f over the tree preorder.
//
// This method is a wrapper for FoldPreorder(tree.root, initialAccumulator, f)
func FoldTreePreorder[T, G any](tree *Tree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return initialAccumulator
	}
	return FoldNodePreorder(treeRoot, initialAccumulator, f)
}

// Fold a function f over the tree Inorder.
//
// This method is a wrapper for FoldInorder(tree.root, initialAccumulator, f)
func FoldTreeInorder[T, G any](tree *Tree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return initialAccumulator
	}
	return FoldNodeInorder(treeRoot, initialAccumulator, f)
}

// Fold a function f over the tree Postorder.
//
// This method is a wrapper for FoldPostorder(tree.root, initialAccumulator, f)
func FoldTreePostorder[T, G any](tree *Tree[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	treeRoot := (*tree).Root()
	if treeRoot == nil {
		return initialAccumulator
	}
	return FoldNodePostorder(treeRoot, initialAccumulator, f)
}
