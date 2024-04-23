package linkedlist

type LinkedListNode[T any] struct {
	item T
	next *LinkedListNode[T]
	prev *LinkedListNode[T]
}

// Get the item of this node.
func (node *LinkedListNode[T]) Item() T {
	return node.item
}

// Return the next node from this node. May be null if this node is the tail.
func (node *LinkedListNode[T]) Next() *LinkedListNode[T] {
	return node.next
}

