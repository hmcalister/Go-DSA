package linkedliststack

import (
	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

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
// Returns a dsa_error.ErrorDataStructureEmpty error if the stack is empty.
func (stack *LinkedListStack[T]) Peek() (T, error) {
	if stack.stackData.Length() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	return stack.stackData.ItemAtIndex(stack.Size() - 1)
}

// Find the first item in a stack matching a predicate.
// The stack is traversed from top to bottom.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (stack *LinkedListStack[T]) Find(predicate func(item T) bool) (T, error) {
	return stack.stackData.ReverseFind(predicate)
}

// Find all items in a stack matching a predicate.
// The stack is traversed from top to bottom.
//
// Returns all items from the stack that match the predicate.
func (stack *LinkedListStack[T]) FindAll(predicate func(item T) bool) []T {
	items := stack.stackData.ReverseFindAll(predicate)
	return items
}

// Get all items from the stack. This method allocates an array of length equal to the number of items.
func (stack *LinkedListStack[T]) Items() []T {
	items := stack.stackData.Items()
	return items
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

// ----------------------------------------------------------------------------
// Remove methods

// Remove an item from the top of the stack.
//
// Returns a dsa_error.ErrorDataStructureEmpty error if the stack is empty.
func (stack *LinkedListStack[T]) Remove() (T, error) {
	if stack.stackData.Length() == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	return stack.stackData.Remove()
}
