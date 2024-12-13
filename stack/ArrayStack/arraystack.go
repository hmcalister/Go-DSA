package arraystack

import (
	"iter"

	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

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
// Returns a dsa_error.ErrorDataStructureEmpty if stack is empty.
func (stack *ArrayStack[T]) Peek() (T, error) {
	if len(stack.stackData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item := stack.stackData[len(stack.stackData)-1]
	return item, nil
}

// Find the first item in a stack matching a predicate.
// The stack is traversed from top to bottom.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (stack *ArrayStack[T]) Find(predicate func(item T) bool) (T, error) {
	for index := len(stack.stackData) - 1; index >= 0; index -= 1 {
		item := stack.stackData[index]
		if predicate(item) {
			return item, nil
		}
	}
	return *new(T), dsa_error.ErrorItemNotFound
}

// Find all items in a stack matching a predicate.
// The stack is traversed from top to bottom.
//
// Returns all items from the stack that match the predicate.
func (stack *ArrayStack[T]) FindAll(predicate func(item T) bool) []T {
	items := make([]T, 0)
	for index := len(stack.stackData) - 1; index >= 0; index -= 1 {
		item := stack.stackData[index]
		if predicate(item) {
			items = append(items, item)
		}
	}
	return items
}

// Get all items from the stack. This method allocates an array of length equal to the number of items..
func (stack *ArrayStack[T]) Items() []T {
	items := make([]T, stack.Size())
	copy(items, stack.stackData)
	return items
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
// Returns a dsa_error.ErrorDataStructureEmpty if the stack is empty.
func (stack *ArrayStack[T]) Remove() (T, error) {
	if len(stack.stackData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	item := stack.stackData[len(stack.stackData)-1]
	stack.stackData = stack.stackData[:len(stack.stackData)-1]
	return item, nil
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
func ForwardApply[T any](stack *ArrayStack[T], f func(item T)) {
	for index := 0; index < len(stack.stackData); index += 1 {
		f(stack.stackData[index])
	}
}

// Iterate over the stack in the forward direction (bottom to top) and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ForwardMap can update the node items by returning the update value.
// If you do not need to modify the stack items, use ForwardApply.
// To accumulate values over the stack, use ForwardFold.
func ForwardMap[T any](stack *ArrayStack[T], f func(item T) T) {
	for index := 0; index < len(stack.stackData); index += 1 {
		stack.stackData[index] = f(stack.stackData[index])
	}
}

// Iterate over the stack (bottom to top) and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on ArrayStack to allow for generic accumulators.
func ForwardFold[T any, G any](stack *ArrayStack[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for index := 0; index < len(stack.stackData); index += 1 {
		accumulator = f(stack.stackData[index], accumulator)
	}

	return accumulator
}

// Iterate over the stack in the reverse direction (top to bottom) and apply a function to each item.
//
// It is expected that ReverseApply does *not* update the stack items.
// To modify the stack items, use ReverseMap.
// To accumulate values over the stack, use ReverseFold.
func ReverseApply[T any](stack *ArrayStack[T], f func(item T)) {
	for index := len(stack.stackData) - 1; index >= 0; index -= 1 {
		f(stack.stackData[index])
	}
}

// Iterate over the stack in the reverse direction (top to bottom) and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ReverseMap can update the node items by returning the update value.
// If you do not need to modify the stack items, use ReverseApply.
// To accumulate values over the stack, use ReverseFold.
func ReverseMap[T any](stack *ArrayStack[T], f func(item T) T) {
	for index := len(stack.stackData) - 1; index >= 0; index -= 1 {
		stack.stackData[index] = f(stack.stackData[index])
	}
}

// Iterate over the stack (top to bottom) and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on ArrayStack to allow for generic accumulators.
func ReverseFold[T any, G any](stack *ArrayStack[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for index := len(stack.stackData) - 1; index >= 0; index -= 1 {
		accumulator = f(stack.stackData[index], accumulator)
	}

	return accumulator
}

// Iterate over the items of the stack in the forward direction (bottom to top).
// Returns both the index (as counted from the bottom of the stack) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (stack *ArrayStack[T]) ForwardIterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for index := 0; index < len(stack.stackData); index += 1 {
			item := stack.stackData[index]
			if !yield(index, item) {
				break
			}
		}
	}
}

// Iterate over the items of the stack in the reverse direction (top to bottom).
// Returns both the index (as counted from the top of the stack) and item.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (stack *ArrayStack[T]) ReverseIterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for index := 0; index < len(stack.stackData); index += 1 {
			item := stack.stackData[len(stack.stackData)-index-1]
			if !yield(index, item) {
				break
			}
		}
	}
}
