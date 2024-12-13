package arrayqueue

import (
	"iter"

	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

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

// ----------------------------------------------------------------------------
// Apply, Map, and Fold methods
//
// Methods to apply a function across ALL items in a queue.

// Iterate over the queue in the forward direction and apply a function to each item.
//
// Idiomatic Go should likely use ForwardIterator() rather than functional methods.
//
// It is expected that ForwardApply does *not* update the queue items.
// To modify the queue items, use ForwardMap.
// To accumulate values over the queue, use ForwardFold.
func ForwardApply[T any](queue *ArrayQueue[T], f func(item T)) {
	for index := 0; index < len(queue.queueData); index += 1 {
		f(queue.queueData[index])
	}
}

// Iterate over the queue in the forward direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// Idiomatic Go should likely use ForwardIterator() rather than functional methods.
//
// ForwardMap can update the node items by returning the update value.
// If you do not need to modify the queue items, use ForwardApply.
// To accumulate values over the queue, use ForwardFold.
func ForwardMap[T any](queue *ArrayQueue[T], f func(item T) T) {
	for index := 0; index < len(queue.queueData); index += 1 {
		queue.queueData[index] = f(queue.queueData[index])
	}
}

// Iterate over the queue and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// Idiomatic Go should likely use ForwardIterator() rather than functional methods.
//
// This function is not a method on ArrayQueue to allow for generic accumulators.
func ForwardFold[T any, G any](queue *ArrayQueue[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for index := 0; index < len(queue.queueData); index += 1 {
		accumulator = f(queue.queueData[index], accumulator)
	}

	return accumulator
}

// Iterate over the queue in the reverse direction and apply a function to each item.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// It is expected that ReverseApply does *not* update the queue items.
// To modify the queue items, use ReverseMap.
// To accumulate values over the queue, use ReverseFold.
func ReverseApply[T any](queue *ArrayQueue[T], f func(item T)) {
	for index := len(queue.queueData) - 1; index >= 0; index -= 1 {
		f(queue.queueData[index])
	}
}

// Iterate over the queue in the reverse direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// ReverseMap can update the node items by returning the update value.
// If you do not need to modify the queue items, use ReverseApply.
// To accumulate values over the queue, use ReverseFold.
func ReverseMap[T any](queue *ArrayQueue[T], f func(item T) T) {
	for index := len(queue.queueData) - 1; index >= 0; index -= 1 {
		queue.queueData[index] = f(queue.queueData[index])
	}
}

// Iterate over the queue and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// This function is not a method on ArrayQueue to allow for generic accumulators.
func ReverseFold[T any, G any](queue *ArrayQueue[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for index := len(queue.queueData) - 1; index >= 0; index -= 1 {
		accumulator = f(queue.queueData[index], accumulator)
	}

	return accumulator
}

// Iterate over the items of the queue in the forward direction (front to back).
// Returns both the index (as counted from the front of the queue) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (queue *ArrayQueue[T]) ForwardIterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for index := 0; index < len(queue.queueData); index += 1 {
			item := queue.queueData[index]
			if !yield(index, item) {
				break
			}
		}
	}
}

// Iterate over the items of the queue in the reverse direction (back to front).
// Returns both the index (as counted from the back of the queue) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (queue *ArrayQueue[T]) ReverseIterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for index := 0; index < len(queue.queueData); index += 1 {
			item := queue.queueData[len(queue.queueData)-index-1]
			if !yield(index, item) {
				break
			}
		}
	}
}
