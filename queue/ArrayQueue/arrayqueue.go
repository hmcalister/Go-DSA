package arrayqueue

// Implement a queue using a array / slice.
//
// Queues are a first in, first out data structure. Items added to the queue are removed in the order they were added.
type ArrayQueue[T any] struct {
	queueData []T
}

