package linkedlistqueue

import linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"

// Implement a queue using a linked list.
type LinkedListQueue[T any] struct {
	queueData *linkedlist.LinkedList[T]
}

