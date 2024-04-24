package redblacktree

type colorEnum int

const (
	BLACK colorEnum = iota
	RED   colorEnum = iota
)

type RedBlackTreeNode[T any] struct {
	// The item of this node
	item T

	// The size of this node, the number of nodes in this subtree
	// (count of this node and all children)
	size int

	// The height of this node, the number of steps to the furthest leaf node
	height int

	// The color of this node, either RED or BLACK
	color colorEnum

	// The parent of this node
	parent *RedBlackTreeNode[T]

	// The left child of this node
	left *RedBlackTreeNode[T]

	// The right child of this node
	right *RedBlackTreeNode[T]
}

// Create a new node from an item.
func newNode[T any](item T) *RedBlackTreeNode[T] {
	return &RedBlackTreeNode[T]{
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
func (node *RedBlackTreeNode[T]) Item() T {
	return node.item
}

// Get the size of this Node, the number of items in the subtree rooted at this node
//
// A leaf node has size 1.
func (node *RedBlackTreeNode[T]) Size() int {
	return node.size
}

// Get the height of this node, the number of steps from this node to the furthest leaf node.
//
// A leaf node has height 0.
func (node *RedBlackTreeNode[T]) Height() int {
	return node.height
}

