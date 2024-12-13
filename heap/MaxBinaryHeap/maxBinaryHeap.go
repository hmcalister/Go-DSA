package maxbinaryheap

import (
	"iter"

	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
	dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"
)

type MaxBinaryHeap[T any] struct {
	heapData           []T
	comparatorFunction comparator.ComparatorFunction[T]
}

// Create a new max-BinaryHeap, with comparator given by the comparatorFunction.
//
// See `github.com/hmcalister/Go-DSA/Comparator` for more information on the comparator.
func New[T any](comparatorFunction comparator.ComparatorFunction[T]) *MaxBinaryHeap[T] {
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

// Max-Heapify the heap.
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
// Get methods

// Peek at the max-element of this heap. The item is not removed from the heap.
//
// If the heap is empty, a dsa_error.ErrorDataStructureEmpty is returned.
func (heap *MaxBinaryHeap[T]) PeekMax() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	return heap.heapData[0], nil
}

// Find the first item in a heap matching a predicate.
//
// Returns (item, nil) if the item is present, or (*new(T), dsa_error.ErrorItemNotFound) if the item is not present.
func (heap *MaxBinaryHeap[T]) Find(predicate func(item T) bool) (T, error) {
	for _, item := range heap.heapData {
		if predicate(item) {
			return item, nil
		}
	}
	return *new(T), dsa_error.ErrorItemNotFound
}

// Find the first item in a heap matching a predicate.
//
// Returns all items from the heap that match the predicate.
func (heap *MaxBinaryHeap[T]) FindAll(predicate func(item T) bool) []T {
	foundItems := make([]T, 0)
	for _, item := range heap.heapData {
		if predicate(item) {
			foundItems = append(foundItems, item)
		}
	}
	return foundItems
}

// Get all items from the heap. This method allocates an array of length equal to the number of items.
func (heap *MaxBinaryHeap[T]) Items() []T {
	items := make([]T, heap.Size())
	copy(items, heap.heapData)
	return items
}

// Get the size of this heap.
func (heap *MaxBinaryHeap[T]) Size() int {
	return len(heap.heapData)
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
// Remove methods

// Remove (and return) the top (maximal) item from this Heap.
//
// If the heap is empty, a dsa_error.ErrorDataStructureEmpty is returned.
func (heap *MaxBinaryHeap[T]) RemoveMax() (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	// Get the root element, so we can return it later
	maxElement := heap.heapData[0]

	// Then, replace the final element with the root element
	// and slice off one element to remove the root
	heapSize := len(heap.heapData) - 1
	heap.heapData[0], heap.heapData[heapSize] = heap.heapData[heapSize], heap.heapData[0]
	heap.heapData = heap.heapData[:heapSize]

	// Finally, heapify the result
	// We only need to heapify the root
	if len(heap.heapData) > 0 {
		heap.maxHeapify(0)
	}

	return maxElement, nil
}

// Remove (and return) an item from the heap.
// If the heap is empty, a dsa_error.ErrorDataStructureEmpty is returned.
// If the item is not present in the tree, a dsa_error.ErrorItemNotFound is returned.
func (heap *MaxBinaryHeap[T]) RemoveItem(item T) (T, error) {
	if len(heap.heapData) == 0 {
		return *new(T), dsa_error.ErrorDataStructureEmpty
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
		return *new(T), dsa_error.ErrorItemNotFound
	}

	// Here's the sneaky trick:
	// Make the target item larger than the root,
	// then heapify to get the node to the top, and finally
	// delete the root using RemoveMax()
	//
	// To do this, we will repeat the heapify algorithm here but enforce largest is always this element
	targetItem := heap.heapData[targetItemIndex]
	currentIndex := targetItemIndex
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		heap.heapData[parentIndex], heap.heapData[currentIndex] = heap.heapData[currentIndex], heap.heapData[parentIndex]
		currentIndex = parentIndex
	}

	heap.RemoveMax()

	return targetItem, nil
}

// ----------------------------------------------------------------------------
// Apply, Map, and Fold methods
//
// Methods to apply a function across ALL items in a heap.

// Iterate over the heap and apply a function to each item.
// In case it matters, the iteration is effectively in "reading order" along the heap.
// Since Apply does not update the heap items, this method does *not* call heapify.
//
// It is expected that Apply does *not* update the heap items.
// To modify the heap items, use Map.
// To accumulate values over the heap, use Fold.
func Apply[T any](heap *MaxBinaryHeap[T], f func(item T)) {
	for index := 0; index < len(heap.heapData); index += 1 {
		f(heap.heapData[index])
	}
}

// Iterate over the heap and apply a function to each item, assigning the result to the item.
// In case it matters, the iteration is effectively in "reading order" along the heap.
// The result of this function is then assigned to the node at each step.
//
// BEWARE: Since this method updates the heap data, this method calls heapify to restore heap order.
// However, since this method may update *all* heap items, this method calls heapify on *all* non-leaf items.
// That is potentially very expensive!
//
// Map can update the node items by returning the update value.
// If you do not need to modify the heap items, use Apply.
// To accumulate values over the heap, use Fold.
func Map[T any](heap *MaxBinaryHeap[T], f func(item T) T) {
	for index := 0; index < len(heap.heapData); index += 1 {
		heap.heapData[index] = f(heap.heapData[index])
	}

	for index := len(heap.heapData) / 2; index >= 0; index -= 1 {
		heap.maxHeapify(index)
	}
}

// Iterate over the heap and apply the function f to it.
// In case it matters, the iteration is effectively in "reading order" along the heap.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on MaxBinaryHeap to allow for generic accumulators.
func Fold[T any, G any](heap *MaxBinaryHeap[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for index := 0; index < len(heap.heapData); index += 1 {
		accumulator = f(heap.heapData[index], accumulator)
	}

	return accumulator
}

// Iterate over the items of the heap.
// In case it matters, the iteration is effectively in "reading order" along the heap.
// This is *not* a sorted order. To iterate in sorted order you may either extract the heap items with Items() and sort,
// or continually pop items from the heap (which will naturally update the heap).
//
// If you are updating items in the heap, please note this method does *not* reheapify.
//
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (heap *MaxBinaryHeap[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for index := 0; index < len(heap.heapData); index += 1 {
			item := heap.heapData[index]
			if !yield(item) {
				break
			}
		}
	}
}
