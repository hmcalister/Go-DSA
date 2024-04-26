package heap

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
func NewMinBinaryHeap[T any](comparatorFunction comparator.ComparatorFunction[T]) *MinBinaryHeap[T] {
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
// Get methods

// Get the Min-element of this heap
//
// If the heap is empty, a EmptyHeapError is returned
func (heap *MinBinaryHeap[T]) GetMin() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), ErrorEmptyHeap
	}

	return heap.heapData[0], nil
}

// Get the size of this heap
func (heap *MinBinaryHeap[T]) Size() int {
	return len(heap.heapData)
}

// ----------------------------------------------------------------------------
// Remove methods

// Remove (and return) the top (Minimal) item from this Heap.
//
// If the heap is empty, a EmptyHeapError is returned
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

