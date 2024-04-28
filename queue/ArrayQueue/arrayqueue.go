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

