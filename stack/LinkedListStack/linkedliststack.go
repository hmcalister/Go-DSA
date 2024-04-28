package linkedliststack

import linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"

// Implement a stack using a linked list.
//
// Stacks are a last in, first out data structure. Items added to the stack are removed in the reverse order they are added.
type LinkedListStack[T any] struct {
	stackData *linkedlist.LinkedList[T]
}

