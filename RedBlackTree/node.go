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
