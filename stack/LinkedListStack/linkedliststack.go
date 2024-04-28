package linkedliststack

import linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"

// Implement a stack using a linked list.
//
// Stacks are a last in, first out data structure. Items added to the stack are removed in the reverse order they are added.
type LinkedListStack[T any] struct {
	stackData *linkedlist.LinkedList[T]
}

// Create a new LinkedListStack using github.com/hmcalister/Go-DSA/list/LinkedList as a backing data structure.
func New[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		stackData: linkedlist.New[T](),
	}
}

// ----------------------------------------------------------------------------
// Get Methods

// Peek at the top item in the stack.
//
// Returns an error if the stack is empty.
func (stack *LinkedListStack[T]) Peek() (T, error) {
	if stack.stackData.Length() == 0 {
		return *new(T), ErrorStackEmpty
	}

	item, err := stack.stackData.ItemAtIndex(stack.Size() - 1)
	if err != nil {
		return *new(T), err
	}

	return item, nil
}

// Get the size of the stack, the number of items in the stack.
func (stack *LinkedListStack[T]) Size() int {
	return stack.stackData.Length()
}

// ----------------------------------------------------------------------------
// Add Methods

// Add an item to the top of the stack.
func (stack *LinkedListStack[T]) Add(item T) {
	stack.stackData.Add(item)
}

