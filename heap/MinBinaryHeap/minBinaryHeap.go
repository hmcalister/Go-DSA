package minbinaryheap

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

type MinBinaryHeap[T any] struct {
	heapData           []T
	comparatorFunction comparator.ComparatorFunction[T]
}

// Create a new Min-BinaryHeap, with comparator given by the comparatorFunction.
//
// See `github.com/hmcalister/Go-DSA/Comparator` for more information on the comparator.
func New[T any](comparatorFunction comparator.ComparatorFunction[T]) *MinBinaryHeap[T] {
	return &MinBinaryHeap[T]{
		// Store the heap as an array.
		// The root is stored in heapData[0], then recursively the
		// node at index `i` has left child at `2i+1` and right child at `2i+2`.
		// Therefore, the parent of a node is given by floor( (i-1) / 2 ).
		//
		// Note that slices are backed by arrays, which is doubled in size when needed
		// making the append calls very cheap (amortized).
		// See https://go.dev/blog/slices
		heapData:           make([]T, 0),
		comparatorFunction: comparatorFunction,
	}
}

// ----------------------------------------------------------------------------
// Heap Helper Methods

// Min-Heapify the heap
func (heap *MinBinaryHeap[T]) minHeapify(targetIndex int) {
	leftIndex := 2*targetIndex + 1
	rightIndex := 2*targetIndex + 2

	smallestIndex := targetIndex
	if leftIndex < len(heap.heapData) && heap.comparatorFunction(heap.heapData[leftIndex], heap.heapData[smallestIndex]) < 0 {
		smallestIndex = leftIndex
	}
	if rightIndex < len(heap.heapData) && heap.comparatorFunction(heap.heapData[rightIndex], heap.heapData[smallestIndex]) < 0 {
		smallestIndex = rightIndex
	}
	if smallestIndex != targetIndex {
		heap.heapData[targetIndex], heap.heapData[smallestIndex] = heap.heapData[smallestIndex], heap.heapData[targetIndex]
		heap.minHeapify(smallestIndex)
	}
}

// ----------------------------------------------------------------------------
// Get methods

// Get the Min-element of this heap. The item is not removed from the heap.
//
// If the heap is empty, a EmptyHeapError is returned.
func (heap *MinBinaryHeap[T]) PeekMin() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), ErrorEmptyHeap
	}

	return heap.heapData[0], nil
}

// Find the first item in a heap matching a predicate.
//
// Returns (item, nil) if the item is present, or (*new(T), ErrorItemNotPresent) if the item is not present.
func (heap *MinBinaryHeap[T]) Find(predicate func(item T) bool) (T, error) {
	for _, item := range heap.heapData {
		if predicate(item) {
			return item, nil
		}
	}
	return *new(T), ErrorItemNotFound
}

// Find the first item in a heap matching a predicate.
//
// Returns all items from the heap that match the predicate.
func (heap *MinBinaryHeap[T]) FindAll(predicate func(item T) bool) []T {
	foundItems := make([]T, 0)
	for _, item := range heap.heapData {
		if predicate(item) {
			foundItems = append(foundItems, item)
		}
	}
	return foundItems
}

// Get the size of this heap
func (heap *MinBinaryHeap[T]) Size() int {
	return len(heap.heapData)
}

// ----------------------------------------------------------------------------
// Add methods

// Add a new element to the heap.
//
// Heaps are allowed to have duplicate values.
func (heap *MinBinaryHeap[T]) Add(item T) {
	// Simply add the new item to the end of the heap
	heap.heapData = append(heap.heapData, item)

	// Then heapify
	//
	// If we have no items or only one item, we are already a heap
	if len(heap.heapData) <= 1 {
		return
	}

	// Start from the lowest leaf node, given by index (n/2 - 1), and walk up the tree to the root,
	// calling the heapify function as we go
	for i := len(heap.heapData)/2 - 1; i >= 0; i -= 1 {
		heap.minHeapify(i)
	}
}

// ----------------------------------------------------------------------------
// Remove methods

// Remove (and return) the top (Minimal) item from this Heap.
//
// If the heap is empty, a EmptyHeapError is returned.
func (heap *MinBinaryHeap[T]) RemoveMin() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), ErrorEmptyHeap
	}

	// Get the root element, so we can return it later
	minElement := heap.heapData[0]

	// Then, replace the final element with the root element
	// and slice off one element to remove the root
	heapSize := len(heap.heapData) - 1
	heap.heapData[0], heap.heapData[heapSize] = heap.heapData[heapSize], heap.heapData[0]
	heap.heapData = heap.heapData[:heapSize]

	// Finally, heapify the result
	// We only need to heapify the root
	if len(heap.heapData) > 0 {
		heap.minHeapify(0)
	}

	return minElement, nil
}

// Remove (and return) an item from the heap.
// If the heap is empty, a ErrorItemNotPresent is returned.
// If the item is not present in the tree, a ErrorItemNotPresent is returned.
func (heap *MinBinaryHeap[T]) RemoveItem(item T) (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), ErrorEmptyHeap
	}

	// First, see if the element exists

	targetItemIndex := -1
	for i, currItem := range heap.heapData {
		if heap.comparatorFunction(currItem, item) == 0 {
			targetItemIndex = i
			break
		}
	}
	// If we did not set the index, we did not find the item
	if targetItemIndex == -1 {
		return *new(T), ErrorItemNotFound
	}

	// Here's the sneaky trick:
	// Make the target item smaller than the root,
	// then heapify to get the node to the top, and finally
	// delete the root using RemoveMin()
	//
	// To do this, we will repeat the heapify algorithm here but enforce smallest is always this element
	targetItem := heap.heapData[targetItemIndex]
	currentIndex := targetItemIndex
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		heap.heapData[parentIndex], heap.heapData[currentIndex] = heap.heapData[currentIndex], heap.heapData[parentIndex]
		currentIndex = parentIndex
	}

	heap.RemoveMin()

	return targetItem, nil
}
