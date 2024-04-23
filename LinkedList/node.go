package linkedlist

type linkedListNode[T any] struct {
	item T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}
