package binarysearchtree

type BinarySearchTree[T any] struct {
	root *binarySearchTreeNode[T]
}

type binarySearchTreeNode[T any] struct {
	item  T
	left  *binarySearchTreeNode[T]
	right *binarySearchTreeNode[T]
}
