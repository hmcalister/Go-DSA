package arrayqueue

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

