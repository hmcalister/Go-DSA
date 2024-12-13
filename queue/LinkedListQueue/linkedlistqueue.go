package linkedlistqueue

import (
	"iter"

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

// Get all items from the queue. This method allocates an array of length equal to the number of items.
func (queue *LinkedListQueue[T]) Items() []T {
	items := queue.queueData.Items()
	return items
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
//
// Internally, this method calls linkedlist.ForwardApply
func ForwardApply[T any](queue *LinkedListQueue[T], f func(item T)) {
	linkedlist.ForwardApply(queue.queueData, f)
}

// Iterate over the queue in the forward direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// Idiomatic Go should likely use ForwardIterator() rather than functional methods.
//
// ForwardMap can update the node items by returning the update value.
// If you do not need to modify the queue items, use ForwardApply.
// To accumulate values over the queue, use ForwardFold.
//
// Internally, this method calls linkedlist.ForwardMap
func ForwardMap[T any](queue *LinkedListQueue[T], f func(item T) T) {
	linkedlist.ForwardMap(queue.queueData, f)
}

// Iterate over the queue and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// Idiomatic Go should likely use ForwardIterator() rather than functional methods.
//
// This function is not a method on LinkedListQueue to allow for generic accumulators.
//
// Internally, this method calls linkedlist.ForwardFold
func ForwardFold[T any, G any](queue *LinkedListQueue[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return linkedlist.ForwardFold(queue.queueData, initialAccumulator, f)
}

// Iterate over the queue in the reverse direction and apply a function to each item.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// It is expected that ReverseApply does *not* update the queue items.
// To modify the queue items, use ReverseMap.
// To accumulate values over the queue, use ReverseFold.
//
// Internally, this method calls linkedlist.ReverseApply
func ReverseApply[T any](queue *LinkedListQueue[T], f func(item T)) {
	linkedlist.ReverseApply(queue.queueData, f)
}

// Iterate over the queue in the reverse direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// ReverseMap can update the node items by returning the update value.
// If you do not need to modify the queue items, use ReverseApply.
// To accumulate values over the queue, use ReverseFold.
//
// Internally, this method calls linkedlist.ReverseMap
func ReverseMap[T any](queue *LinkedListQueue[T], f func(item T) T) {
	linkedlist.ReverseMap(queue.queueData, f)
}

// Iterate over the queue and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// Idiomatic Go should likely use ReverseIterator() rather than functional methods.
//
// This function is not a method on LinkedListQueue to allow for generic accumulators.
//
// Internally, this method calls linkedlist.ReverseFold
func ReverseFold[T any, G any](queue *LinkedListQueue[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return linkedlist.ReverseFold(queue.queueData, initialAccumulator, f)
}

// Iterate over the items of the queue in the forward direction (front to back).
// Returns both the index (from the front of the queue) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (queue *LinkedListQueue[T]) ForwardIterator() iter.Seq2[int, T] {
	return queue.queueData.ForwardIterator()
}

// Iterate over the items of the queue in the reverse direction (top to bottom).
// Returns both the index from the back of the queue) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (queue *LinkedListQueue[T]) ReverseIterator() iter.Seq2[int, T] {
	return queue.queueData.ReverseIterator()
}
