package maxbinaryheap_test

import (
	"math/rand"
	"slices"
	"testing"

	maxbinaryheap "github.com/hmcalister/Go-DSA/heap/MaxBinaryHeap"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

// ----------------------------------------------------------------------------
// Initialization Tests

func TestMaxHeapIntInit(t *testing.T) {
	maxbinaryheap.New[int](comparator.DefaultIntegerComparator)
}

func TestMaxHeapFloatInit(t *testing.T) {
	maxbinaryheap.New[float64](comparator.DefaultFloat64Comparator)
}

func TestMaxHeapStringInit(t *testing.T) {
	maxbinaryheap.New[string](comparator.DefaultStringComparator)
}

func TestMaxHeapStructInit(t *testing.T) {
	type S struct {
		_ int
		f float64
	}
	maxbinaryheap.New[S](func(a, b S) int {
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

func TestMaxHeapRemoveFromEmpty(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	_, err := heap.RemoveMax()
	if err == nil {
		t.Errorf("got nil error when removing from empty heap")
	}
}

func TestMaxHeapEmptySize(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMaxHeapSingleItemSize(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)
	heap.Add(1)

	expectedSize := 1
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMaxHeapAfterRemoveSize(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)
	heap.Add(1)
	heap.RemoveMax()

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

// ----------------------------------------------------------------------------
// Add Tests

func TestMaxHeapAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}
}

func TestMaxHeapAddRandomOrder(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

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

func TestMaxHeapRemoveMax(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	for range items {
		_, err := heap.RemoveMax()
		if err != nil {
			t.Errorf("failed to remove max item from a heap of size %v", heap.Size())
		}
	}
}

func TestMaxHeapRemoveMaxItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	slices.Reverse(items)
	for _, expectedItem := range items {
		removedItem, err := heap.RemoveMax()
		if err != nil {
			t.Errorf("failed to remove max item from a heap of size %v", heap.Size())
		}
		if removedItem != expectedItem {
			t.Errorf("removed max item (%v) does not match expected max item (%v)", removedItem, expectedItem)
		}
	}
}

func TestMaxHeapRemoveItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	targetItem := 5
	removedItem, err := heap.RemoveItem(targetItem)
	if err != nil {
		t.Errorf("failed to remove item from a heap of size %v", heap.Size())
	}
	if removedItem != targetItem {
		t.Errorf("removed item item (%v) does not match expected max item (%v)", removedItem, targetItem)
	}
}

func TestMaxHeapManyRemoveItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	for _, targetItem := range items {
		removedItem, err := heap.RemoveItem(targetItem)
		if err != nil {
			t.Errorf("failed to remove item from a heap of size %v", heap.Size())
		}
		if removedItem != targetItem {
			t.Errorf("removed item item (%v) does not match expected max item (%v)", removedItem, targetItem)
		}
	}
}

// ----------------------------------------------------------------------------
// Get max Tests

func TestMaxHeapAddPeekMax(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

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

	heapMaxItem, err := heap.PeekMax()
	expectedMaxItem := numItems - 1
	if err != nil {
		t.Errorf("found error when getting max item from a non-empty heap: %v", err)
	}
	if heapMaxItem != expectedMaxItem {
		t.Errorf("heap max item (%v) does not match expected max item (%v)", heapMaxItem, expectedMaxItem)
	}
}

func TestMaxHeapRemovePeekMax(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

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
		heap.RemoveMax()
	}

	heapMaxItem, err := heap.PeekMax()
	expectedMaxItem := numItems - numRemove - 1
	if err != nil {
		t.Errorf("found error when getting max item from a non-empty heap: %v", err)
	}
	if heapMaxItem != expectedMaxItem {
		t.Errorf("heap max item (%v) does not match expected max item (%v)", heapMaxItem, expectedMaxItem)
	}
}

func TestMaxBinaryHeapItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		heap.Add(item)
	}

	retrievedItems := heap.Items()
	for _, item := range items {
		if !slices.Contains(retrievedItems, item) {
			t.Errorf("retrieved items %v does not contain expected item %v", retrievedItems, item)
		}
	}
}
