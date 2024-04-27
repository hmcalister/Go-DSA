package priorityqueue

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
	minbinaryheap "github.com/hmcalister/Go-DSA/heap/MinBinaryHeap"
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
// Returns an error if the queue is empty.
func (queue *PriorityQueue[T]) Peek() (T, error) {
	if queue.queueData.Size() == 0 {
		return *new(T), ErrorQueueEmpty
	}

	item, err := queue.queueData.PeekMin()
	if err != nil {
		return *new(T), err
	}

	return item, nil
}

