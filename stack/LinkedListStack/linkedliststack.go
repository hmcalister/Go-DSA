package linkedliststack

import linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"

// Implement a stack using a linked list.
//
// Stacks are a last in, first out data structure. Items added to the stack are removed in the reverse order they are added.
type LinkedListStack[T any] struct {
	stackData *linkedlist.LinkedList[T]
}

// Create a new LinkedListStack using github.com/hmcalister/Go-DSA/list/LinkedList as a backing data structure.
func New[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		stackData: linkedlist.New[T](),
	}
}

