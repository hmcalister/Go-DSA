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

