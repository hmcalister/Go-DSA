package linkedlist

type LinkedListNode[T any] struct {
	item T
	next *LinkedListNode[T]
	prev *LinkedListNode[T]
}

