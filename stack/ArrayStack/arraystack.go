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

// ----------------------------------------------------------------------------
// Get Methods

// Peek at the top item in the stack.
//
// Returns an error if the stack is empty.
func (stack *ArrayStack[T]) Peek() (T, error) {
	if len(stack.stackData) == 0 {
		return *new(T), ErrorStackEmpty
	}

	item := stack.stackData[len(stack.stackData)-1]
	return item, nil
}

// Get the size of the stack, the number of items in the stack.
func (stack *ArrayStack[T]) Size() int {
	return len(stack.stackData)
}

// ----------------------------------------------------------------------------
// Add Methods

// Add an item to the top of the stack.
func (stack *ArrayStack[T]) Add(item T) {
	stack.stackData = append(stack.stackData, item)
}

// ----------------------------------------------------------------------------
// Remove methods

// Remove an item from the top of the stack.
//
// Returns an error if the stack is empty.
func (stack *ArrayStack[T]) Remove() (T, error) {
	if len(stack.stackData) == 0 {
		return *new(T), ErrorStackEmpty
	}

	item := stack.stackData[len(stack.stackData)-1]
	stack.stackData = stack.stackData[:len(stack.stackData)-1]
	return item, nil
}