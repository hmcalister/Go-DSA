package priorityqueue

import (
	"iter"

	minbinaryheap "github.com/hmcalister/Go-DSA/heap/MinBinaryHeap"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

// Implement a priority queue.
//
// A priority queue will accept items and ensure those items are retrievable in priority order.
//
// This implementation uses a min-heap (github.com/hmcalister/Go-DSA/heap/MinBinaryHeap) and hence lower priority values are put at the front of the queue.
// If you require the opposite behavior, simply flip the logic in the comparator passed to the constructor.
type PriorityQueue[T any] struct {
	queueData          *minbinaryheap.MinBinaryHeap[T]
	comparatorFunction comparator.ComparatorFunction[T]
}

// Create a new priority queue.
//
// The comparatorFunction allows for items in the queue to be compared based on priority.
// Remember that lower priority values are pushed to the front of the queue.
func New[T any](comparatorFunction comparator.ComparatorFunction[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		queueData:          minbinaryheap.New[T](comparatorFunction),
		comparatorFunction: comparatorFunction,
	}
}

// ----------------------------------------------------------------------------
// Get Methods

// Peek at the front item in the queue.
//
// Returns a dsa_error.ErrorDataStructureEmpty error if the queue is empty.
func (queue *PriorityQueue[T]) Peek() (T, error) {
	if queue.queueData.Size() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item, err := queue.queueData.PeekMin()
	if err != nil {
		return *new(T), err
	}

	return item, nil
}

// Find the first item in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (queue *PriorityQueue[T]) Find(predicate func(item T) bool) (T, error) {
	item, err := queue.queueData.Find(predicate)
	if err != nil {
		return *new(T), dsa_error.ErrorItemNotFound
	}
	return item, nil
}

// Find all items in a queue matching a predicate.
// The queue is traversed from front to back.
//
// Returns all items from the queue that match the predicate.
func (queue *PriorityQueue[T]) FindAll(predicate func(item T) bool) []T {
	return queue.queueData.FindAll(predicate)
}

// Get all items from the queue. This method allocates an array of length equal to the number of items.
func (queue *PriorityQueue[T]) Items() []T {
	items := queue.queueData.Items()
	return items
}

// Get the size of the queue, the number of items in the queue.
func (queue *PriorityQueue[T]) Size() int {
	return queue.queueData.Size()
}

// ----------------------------------------------------------------------------
// Add Methods

// Enqueue an item, adding it to the end of the queue.
//
// This method automatically updates the priority queue to ensure the head item has the lowest priority value.
func (queue *PriorityQueue[T]) Add(item T) {
	queue.queueData.Add(item)
}

// ----------------------------------------------------------------------------
// Remove Methods

// Dequeue an item, removing from the front of the queue.
//
// Returns a dsa_error.ErrorDataStructureEmpty if the queue is empty.
func (queue *PriorityQueue[T]) Remove() (T, error) {
	if queue.queueData.Size() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item, err := queue.queueData.RemoveMin()
	if err != nil {
		return *new(T), err
	}

	return item, nil
}

// ----------------------------------------------------------------------------
// Apply, Map, and Fold methods
//
// Methods to apply a function across ALL items in a queue.

// Iterate over the queue and apply a function to each item.
//
// BEWARE: Iteration order is not the same as priority order!
// To iterate in priority order, use Items() and sort by priority.
//
// Since Apply does not update the queue items, this method does *not* call heapify (reorganize the queue).
//
// Internally this method calls minbinaryheap.Apply, as the backing data structure is a heap.
//
// It is expected that Apply does *not* update the queue items.
// To modify the queue items, use Map.
// To accumulate values over the queue, use Fold.
func Apply[T any](queue *PriorityQueue[T], f func(item T)) {
	minbinaryheap.Apply(queue.queueData, f)
}

// Iterate over the queue apply a function to each item.
//
// BEWARE: Iteration order is not the same as priority order!
// To iterate in priority order, use Items() and sort by priority.
//
// BEWARE: Since this method updates the queue data, this method calls heapify to restore queue order.
// However, since this method may update *all* queue items, this method calls heapify on *all* items.
// That is potentially very expensive!
//
// Internally this method calls minbinaryheap.Map, as the backing data structure is a heap.
//
// Map can update the node items by returning the update value.
// If you do not need to modify the queue items, use Apply.
// To accumulate values over the queue, use Fold.
func Map[T any](queue *PriorityQueue[T], f func(item T) T) {
	minbinaryheap.Map(queue.queueData, f)
}

// Iterate over the queue and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// BEWARE: Iteration order is not the same as priority order!
// To iterate in priority order, use Items() and sort by priority.
//
// This function returns the final accumulator.
//
// Internally this method calls minbinaryheap.Fold, as the backing data structure is a heap.
//
// This function is not a method on PriorityQueue to allow for generic accumulators.
func Fold[T any, G any](queue *PriorityQueue[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return minbinaryheap.Fold(queue.queueData, initialAccumulator, f)
}

// Iterate over the items of the queue.
//
// BEWARE: Iteration order is not the same as priority order!
// To iterate in priority order, use Items() and sort by priority.
//
// If you are updating items in the queue, please note this method does *not* reheapify.
//
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (queue *PriorityQueue[T]) Iterator() iter.Seq[T] {
	return queue.queueData.Iterator()
}
