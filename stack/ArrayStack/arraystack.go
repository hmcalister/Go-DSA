package arraystack

// Implement a stack using a array (slice) as the backing data structure.
//
// Stacks are a last in, first out data structure. Items added to the stack are removed in the reverse order they are added.
type ArrayStack[T any] struct {
	stackData []T
}

// Create a new ArrayStack using an array as a backing data structure.
func New[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		// Slices are backed by arrays which grow with a growth factor of 2.
		//
		// This will be fine for our purposes.
		stackData: make([]T, 0),
	}
}

