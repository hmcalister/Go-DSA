package linkedlist

type LinkedList[T any] struct {
	// Head of the list, the first Node
	//
	// nil only when the length is zero
	head *LinkedListNode[T]
	// Tail of the list, the last Node
	//
	// nil only when the length is zero
	tail *LinkedListNode[T]

	// Length of the list, the total number of Nodes
	length int
}

type LinkedListNode[T any] struct {
	Item T
	next *LinkedListNode[T]
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

