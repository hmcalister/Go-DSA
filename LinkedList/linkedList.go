package linkedlist

type LinkedList[T any] struct {
	head *LinkedListNode[T]
}

type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
}
