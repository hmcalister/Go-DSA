package linkedlistqueue

import (
	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

// Implement a queue using a linked list.
//
// Queues are a first in, first out data structure. Items added to the queue are removed in the order they were added.
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
// Get Methods

// Peek at the front item in the queue.
//
// Returns a dsa_error.DataStructureEmpty error if the queue is empty.
func (queue *LinkedListQueue[T]) Peek() (T, error) {
	if queue.queueData.Length() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	return queue.queueData.ItemAtIndex(0)
}

// Find the first item in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (queue *LinkedListQueue[T]) Find(predicate func(item T) bool) (T, error) {
	return queue.queueData.Find(predicate)
}

// Find all items in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns all items from the queue that match the predicate.
func (queue *LinkedListQueue[T]) FindAll(predicate func(item T) bool) []T {
	return queue.queueData.FindAll(predicate)
}

// Get the size of the queue, the number of items in the queue.
func (queue *LinkedListQueue[T]) Size() int {
	return queue.queueData.Length()
}

// ----------------------------------------------------------------------------
// Add Methods

// Enqueue an item, adding it to the end of the queue.
func (queue *LinkedListQueue[T]) Add(item T) {
	queue.queueData.Add(item)
}

// ----------------------------------------------------------------------------
// Remove methods

// Dequeue an item, removing from the front of the queue.
//
// Returns a dsa_error.ErrorDataStructureEmpty error if the queue is empty.
func (queue *LinkedListQueue[T]) Remove() (T, error) {
	if queue.queueData.Length() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	return queue.queueData.RemoveAtIndex(0)
}
