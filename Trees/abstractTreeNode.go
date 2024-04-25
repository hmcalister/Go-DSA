package abstracttree

type TreeNode[T any] interface {
	// Get the item of this tree node
	//
	// BEWARE: Mutating this item (e.g. if this item is a struct, array, etc...) may break the tree structure!
	// Only mutate the result of node.Item() if:
	// i) The type of T is a primitive, such as int, float... in which case the result is copied anyway
	// ii) You can ensure your mutation will not change the ordering based on the tree's ComparatorFunction
	Item() T

	// Get this node's left child (may be nil)
	Left() *TreeNode[T]

	// Get this node's right child (may be nil)
	Right() *TreeNode[T]

	// Get the successor of this node (the next largest item in this tree).
	// Is nil if the current item is the largest in the tree.
	Successor() *TreeNode[T]

	// Get the predecessor of this node (the next smallest item in this tree).
	// Is nil if the current item is the smallest in the tree.
	Predecessor() *TreeNode[T]

	// Get the size of this Node, the number of items in the subtree rooted at this node
	Size() int

	// Get the height of this node, the number of steps from this node to the furthest leaf node.
	Height() int
}

// ----------------------------------------------------------------------------
// Apply Methods

// Apply a function f to each node in a tree Preorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodePreorder[T any](node *TreeNode[T], f func(item T)) {
	f((*node).Item())
	if (*node).Left() != nil {
		ApplyNodePreorder((*node).Left(), f)
	}
	if (*node).Right() != nil {
		ApplyNodePreorder((*node).Right(), f)
	}
}

// Apply a function f to each node in a tree Inorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodeInorder[T any](node *TreeNode[T], f func(item T)) {
	if (*node).Left() != nil {
		ApplyNodeInorder((*node).Left(), f)
	}
	f((*node).Item())
	if (*node).Right() != nil {
		ApplyNodeInorder((*node).Right(), f)
	}
}

// Apply a function f to each node in a tree Postorder.
//
// Apply should not change the item in a Node, as this could affect the binary tree structure.
func ApplyNodePostorder[T any](node *TreeNode[T], f func(item T)) {
	if (*node).Left() != nil {
		ApplyNodePostorder((*node).Left(), f)
	}
	if (*node).Right() != nil {
		ApplyNodePostorder((*node).Right(), f)
	}
	f((*node).Item())
}

// ----------------------------------------------------------------------------
// Fold Methods

// Fold a function f (taking the current node item and the accumulator value) across the tree Preorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodePreorder[T, G any](node *TreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	currentAccumulator = f((*node).Item(), currentAccumulator)
	if (*node).Left() != nil {
		currentAccumulator = FoldNodePreorder((*node).Left(), currentAccumulator, f)
	}
	if (*node).Right() != nil {
		currentAccumulator = FoldNodePreorder((*node).Right(), currentAccumulator, f)
	}

	return currentAccumulator
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Inorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodeInorder[T, G any](node *TreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if (*node).Left() != nil {
		currentAccumulator = FoldNodeInorder((*node).Left(), currentAccumulator, f)
	}
	currentAccumulator = f((*node).Item(), currentAccumulator)
	if (*node).Right() != nil {
		currentAccumulator = FoldNodeInorder((*node).Right(), currentAccumulator, f)
	}

	return currentAccumulator
}

// Fold a function f (taking the current node item and the accumulator value) across the tree Postorder.
// f must return the next value of the accumulator.
//
// Returns the final accumulator value
func FoldNodePostorder[T, G any](node *TreeNode[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentAccumulator := initialAccumulator

	if (*node).Left() != nil {
		currentAccumulator = FoldNodePostorder((*node).Left(), currentAccumulator, f)
	}
	if (*node).Right() != nil {
		currentAccumulator = FoldNodePostorder((*node).Right(), currentAccumulator, f)
	}
	currentAccumulator = f((*node).Item(), currentAccumulator)

	return currentAccumulator
}
