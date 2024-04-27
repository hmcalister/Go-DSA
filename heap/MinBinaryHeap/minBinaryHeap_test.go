package minbinaryheap_test

import (
	"math/rand"
	"slices"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	minbinaryheap "github.com/hmcalister/Go-DSA/heap/MinBinaryHeap"
)

// ----------------------------------------------------------------------------
// Initialization Tests

func TestMinHeapIntInit(t *testing.T) {
	minbinaryheap.New[int](comparator.DefaultIntegerComparator)
}

func TestMinHeapFloatInit(t *testing.T) {
	minbinaryheap.New[float64](comparator.DefaultFloat64Comparator)
}

func TestMinHeapStringInit(t *testing.T) {
	minbinaryheap.New[string](comparator.DefaultStringComparator)
}

func TestMinHeapStructInit(t *testing.T) {
	type S struct {
		_ int
		f float64
	}
	minbinaryheap.New[S](func(a, b S) int {
		if a.f > b.f {
			return 1
		} else if a.f < b.f {
			return -1
		} else {
			return 0
		}
	})
}

// ----------------------------------------------------------------------------
// Misc Tests

func TestMinHeapRemoveFromEmpty(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	_, err := heap.RemoveMin()
	if err == nil {
		t.Errorf("got nil error when removing from empty heap")
	}
}

func TestMinHeapEmptySize(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMinHeapSingleItemSize(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)
	heap.Add(1)

	expectedSize := 1
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMinHeapAfterRemoveSize(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)
	heap.Add(1)
	heap.RemoveMin()

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

// ----------------------------------------------------------------------------
// Add Tests

func TestMinHeapAdd(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}
}

func TestMinHeapAddRandomOrder(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	numItems := 100
	items := make([]int, numItems)
	for i := range numItems {
		items[i] = i
	}
	rand.Shuffle(numItems, func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	for _, item := range items {
		heap.Add(item)
	}
}

// ----------------------------------------------------------------------------
// Remove Tests

func TestMinHeapRemoveMin(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	for range items {
		_, err := heap.RemoveMin()
		if err != nil {
			t.Errorf("failed to remove min item from a heap of size %v", heap.Size())
		}
	}
}

func TestMinHeapRemoveMinItem(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	slices.Reverse(items)
	for _, expectedItem := range items {
		removedItem, err := heap.RemoveMin()
		if err != nil {
			t.Errorf("failed to remove min item from a heap of size %v", heap.Size())
		}
		if removedItem != expectedItem {
			t.Errorf("removed min item (%v) does not match expected min item (%v)", removedItem, expectedItem)
		}
	}
}

func TestMinHeapRemoveItem(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	targetItem := 5
	removedItem, err := heap.RemoveItem(targetItem)
	if err != nil {
		t.Errorf("failed to remove item from a heap of size %v", heap.Size())
	}
	if removedItem != targetItem {
		t.Errorf("removed item item (%v) does not match expected min item (%v)", removedItem, targetItem)
	}
}

func TestMinHeapManyRemoveItem(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	for _, targetItem := range items {
		removedItem, err := heap.RemoveItem(targetItem)
		if err != nil {
			t.Errorf("failed to remove item from a heap of size %v", heap.Size())
		}
		if removedItem != targetItem {
			t.Errorf("removed item item (%v) does not match expected min item (%v)", removedItem, targetItem)
		}
	}
}

// ----------------------------------------------------------------------------
// Get min Tests

func TestMinHeapAddPeekMin(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	numItems := 100
	items := make([]int, numItems)
	for i := range numItems {
		items[i] = i
	}
	rand.Shuffle(numItems, func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	for _, item := range items {
		heap.Add(item)
	}

	heapMinItem, err := heap.PeekMin()
	expectedMinItem := 0
	if err != nil {
		t.Errorf("found error when getting min item from a non-empty heap: %v", err)
	}
	if heapMinItem != expectedMinItem {
		t.Errorf("heap min item (%v) does not match expected min item (%v)", heapMinItem, expectedMinItem)
	}
}

func TestMinHeapRemovePeekMin(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)

	numItems := 100
	items := make([]int, numItems)
	for i := range numItems {
		items[i] = i
	}
	rand.Shuffle(numItems, func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	for _, item := range items {
		heap.Add(item)
	}

	numRemove := 50
	for range numRemove {
		heap.RemoveMin()
	}

	heapMinItem, err := heap.PeekMin()
	expectedMinItem := numRemove
	if err != nil {
		t.Errorf("found error when getting min item from a non-empty heap: %v", err)
	}
	if heapMinItem != expectedMinItem {
		t.Errorf("heap min item (%v) does not match expected min item (%v)", heapMinItem, expectedMinItem)
	}
}
