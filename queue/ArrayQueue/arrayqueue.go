package arrayqueue

import dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"

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
// Returns a dsa_error.ErrorDataStructureEmpty if the queue is empty.
func (queue *ArrayQueue[T]) Peek() (T, error) {
	if len(queue.queueData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item := queue.queueData[0]
	return item, nil
}

// Find the first item in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (queue *ArrayQueue[T]) Find(predicate func(item T) bool) (T, error) {
	for _, item := range queue.queueData {
		if predicate(item) {
			return item, nil
		}
	}
	return *new(T), dsa_error.ErrorItemNotFound
}

// Find all items in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns all items from the queue that match the predicate.
func (queue *ArrayQueue[T]) FindAll(predicate func(item T) bool) []T {
	foundItems := make([]T, 0)
	for _, item := range queue.queueData {
		if predicate(item) {
			foundItems = append(foundItems, item)
		}
	}
	return foundItems
}

// Get all items from the queue. This method allocates an array of length equal to the number of items.
func (queue *ArrayQueue[T]) Items() []T {
	items := make([]T, queue.Size())
	copy(items, queue.queueData)
	return items
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
// Returns a dsa_error.ErrorDataStructureEmpty error if the queue is empty.
func (queue *ArrayQueue[T]) Remove() (T, error) {
	if len(queue.queueData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item := queue.queueData[0]
	queue.queueData = queue.queueData[1:]
	return item, nil
}
