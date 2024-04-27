package linkedlistqueue

import linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"

// Implement a queue using a linked list.
type LinkedListQueue[T any] struct {
	queueData *linkedlist.LinkedList[T]
}

// Create a new LinkedListQueue using github.com/hmcalister/Go-DSA/list/LinkedList as a backing data structure.
func New[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{
		queueData: linkedlist.New[T](),
	}
}

// ----------------------------------------------------------------------------
// Add Methods

// Enqueue an item, adding it to the end of the queue.
func (queue *LinkedListQueue[T]) Add(item T) {
	queue.queueData.Add(item)
}

// ----------------------------------------------------------------------------
// Get Methods

// Peek at the front item in the queue.
//
// Returns an error if the queue is empty.
func (queue *LinkedListQueue[T]) Peek() (T, error) {
	if queue.queueData.Length() == 0 {
		return *new(T), ErrorQueueEmpty
	}

	item, err := queue.queueData.ItemAtIndex(0)
	if err != nil {
		return *new(T), err
	}

	return item, nil
}

