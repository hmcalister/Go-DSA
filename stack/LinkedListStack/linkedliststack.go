package linkedliststack

import (
	"iter"

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

// ----------------------------------------------------------------------------
// Apply, Map, and Fold methods
//
// Methods to apply a function across ALL items in a stack.

// Iterate over the stack in the forward direction (bottom to top) and apply a function to each item.
//
// It is expected that ForwardApply does *not* update the stack items.
// To modify the stack items, use ForwardMap.
// To accumulate values over the stack, use ForwardFold.
//
// Internally, this method calls linkedlist.ForwardApply
func ForwardApply[T any](stack *LinkedListStack[T], f func(item T)) {
	linkedlist.ForwardApply(stack.stackData, f)
}

// Iterate over the stack in the forward direction (bottom to top) and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ForwardMap can update the node items by returning the update value.
// If you do not need to modify the stack items, use ForwardApply.
// To accumulate values over the stack, use ForwardFold.
//
// Internally, this method calls linkedlist.ForwardMap
func ForwardMap[T any](stack *LinkedListStack[T], f func(item T) T) {
	linkedlist.ForwardMap(stack.stackData, f)
}

// Iterate over the stack (bottom to top) and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on LinkedListStack to allow for generic accumulators.
//
// Internally, this method calls linkedlist.ForwardFold
func ForwardFold[T any, G any](stack *LinkedListStack[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return linkedlist.ForwardFold(stack.stackData, initialAccumulator, f)
}

// Iterate over the stack in the reverse direction (top to bottom) and apply a function to each item.
//
// It is expected that ReverseApply does *not* update the stack items.
// To modify the stack items, use ReverseMap.
// To accumulate values over the stack, use ReverseFold.
//
// Internally, this method calls linkedlist.ReverseApply
func ReverseApply[T any](stack *LinkedListStack[T], f func(item T)) {
	linkedlist.ReverseApply(stack.stackData, f)
}

// Iterate over the stack in the reverse direction (top to bottom) and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ReverseMap can update the node items by returning the update value.
// If you do not need to modify the stack items, use ReverseApply.
// To accumulate values over the stack, use ReverseFold.
//
// Internally, this method calls linkedlist.ReverseMap
func ReverseMap[T any](stack *LinkedListStack[T], f func(item T) T) {
	linkedlist.ReverseMap(stack.stackData, f)
}

// Iterate over the stack (top to bottom) and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on LinkedListStack to allow for generic accumulators.
//
// Internally, this method calls linkedlist.ReverseFold
func ReverseFold[T any, G any](stack *LinkedListStack[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	return linkedlist.ReverseFold(stack.stackData, initialAccumulator, f)
}

// Iterate over the items of the stack in the forward direction (bottom to top).
// Returns both the index (as counted from the bottom of the stack) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (stack *LinkedListStack[T]) ForwardIterator() iter.Seq2[int, T] {
	return stack.stackData.ForwardIterator()
}

// Iterate over the items of the stack in the reverse direction (top to bottom).
// Returns both the index (as counted from the top of the stack) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (stack *LinkedListStack[T]) ReverseIterator() iter.Seq2[int, T] {
	return stack.stackData.ReverseIterator()
}
