package heap

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

type MaxBinaryHeap[T any] struct {
	heapData           []T
	comparatorFunction comparator.ComparatorFunction[T]
}

// Create a new max-BinaryHeap, with comparator given by the comparatorFunction.
//
// See `github.com/hmcalister/Go-DSA/Comparator` for more information on the comparator.
func NewMaxBinaryHeap[T any](comparatorFunction comparator.ComparatorFunction[T]) *MaxBinaryHeap[T] {
	return &MaxBinaryHeap[T]{
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

// Max-Heapify the heap
func (heap *MaxBinaryHeap[T]) maxHeapify(targetIndex int) {
	leftIndex := 2*targetIndex + 1
	rightIndex := 2*targetIndex + 2

	largestIndex := targetIndex
	if leftIndex < len(heap.heapData) && heap.comparatorFunction(heap.heapData[leftIndex], heap.heapData[largestIndex]) > 0 {
		largestIndex = leftIndex
	}
	if rightIndex < len(heap.heapData) && heap.comparatorFunction(heap.heapData[rightIndex], heap.heapData[largestIndex]) > 0 {
		largestIndex = rightIndex
	}
	if largestIndex != targetIndex {
		heap.heapData[targetIndex], heap.heapData[largestIndex] = heap.heapData[largestIndex], heap.heapData[targetIndex]
		heap.maxHeapify(largestIndex)
	}
}

// ----------------------------------------------------------------------------
// Add methods

// Add a new element to the heap.
//
// Heaps are allowed to have duplicate values.
func (heap *MaxBinaryHeap[T]) Add(item T) {
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
		heap.maxHeapify(i)
	}
}

// ----------------------------------------------------------------------------
// Get methods

// Get the max-element of this heap
//
// If the heap is empty, a EmptyHeapError is returned
func (heap *MaxBinaryHeap[T]) GetMax() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), ErrorEmptyHeap
	}

	return heap.heapData[0], nil
}

// Get the size of this heap
func (heap *MaxBinaryHeap[T]) Size() int {
	return len(heap.heapData)
}

