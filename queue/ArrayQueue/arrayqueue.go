package arrayqueue

import comparator "github.com/hmcalister/Go-DSA/Comparator"

// Implement a queue using a array / slice.
//
// Queues are a first in, first out data structure. Items added to the queue are removed in the order they were added.
type ArrayQueue[T any] struct {
	queueData []T
}

// Create a new ArrayQueue using github.com/hmcalister/Go-DSA/list/ArrayQueue as a backing data structure.
func New[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		// Slices are backed by arrays which grow with a growth factor of 2.
		//
		// This will be fine for our purposes.
		queueData: make([]T, 0),
	}
}

// ----------------------------------------------------------------------------
// Get Methods

// Peek at the front item in the queue.
//
// Returns an error if the queue is empty.
func (queue *ArrayQueue[T]) Peek() (T, error) {
	if len(queue.queueData) == 0 {
		return *new(T), ErrorQueueEmpty
	}

	item := queue.queueData[0]
	return item, nil
}

// Find an item in a queue.
//
// Requires a comparator that can compare items of type T. The comparator must return 0 when two items are equal.
// See github.com/hmcalister/Go-DSA/Comparator for more information.
//
// Returns (item, nil) if the item is present, or (*new(T), ErrorItemNotFound) if the item is not present.
func (queue *ArrayQueue[T]) Find(targetItem T, comparatorFunction comparator.ComparatorFunction[T]) (T, error) {
	for _, item := range queue.queueData {
		if comparatorFunction(targetItem, item) == 0 {
			return item, nil
		}
	}

	return *new(T), ErrorItemNotFound
}

// Get the size of the queue, the number of items in the queue.
func (queue *ArrayQueue[T]) Size() int {
	return len(queue.queueData)
}

// ----------------------------------------------------------------------------
// Add Methods

// Enqueue an item, adding it to the end of the queue.
func (queue *ArrayQueue[T]) Add(item T) {
	queue.queueData = append(queue.queueData, item)
}

// ----------------------------------------------------------------------------
// Remove methods

// Dequeue an item, removing from the front of the queue.
//
// Returns an error if the queue is empty.
func (queue *ArrayQueue[T]) Remove() (T, error) {
	if len(queue.queueData) == 0 {
		return *new(T), ErrorQueueEmpty
	}

	item := queue.queueData[0]
	queue.queueData = queue.queueData[1:]
	return item, nil
}
