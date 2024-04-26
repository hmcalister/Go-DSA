package heap_test

import (
	"math/rand"
	"slices"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	heap "github.com/hmcalister/Go-DSA/Heap"
)

// ----------------------------------------------------------------------------
// Initialization Tests

func TestMinHeapIntInit(t *testing.T) {
	heap.NewMinBinaryHeap[int](comparator.DefaultIntegerComparator)
}

func TestMinHeapFloatInit(t *testing.T) {
	heap.NewMinBinaryHeap[float64](comparator.DefaultFloat64Comparator)
}

func TestMinHeapStringInit(t *testing.T) {
	heap.NewMinBinaryHeap[string](comparator.DefaultStringComparator)
}

func TestMinHeapStructInit(t *testing.T) {
	type S struct {
		_ int
		f float64
	}
	heap.NewMinBinaryHeap[S](func(a, b S) int {
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
	heap := heap.NewMinBinaryHeap[int](comparator.DefaultIntegerComparator)

	_, err := heap.RemoveMin()
	if err == nil {
		t.Errorf("got nil error when removing from empty heap")
	}
}

func TestMinHeapEmptySize(t *testing.T) {
	heap := heap.NewMinBinaryHeap[int](comparator.DefaultIntegerComparator)

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMinHeapSingleItemSize(t *testing.T) {
	heap := heap.NewMinBinaryHeap[int](comparator.DefaultIntegerComparator)
	heap.Add(1)

	expectedSize := 1
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

func TestMinHeapAfterRemoveSize(t *testing.T) {
	heap := heap.NewMinBinaryHeap[int](comparator.DefaultIntegerComparator)
	heap.Add(1)
	heap.RemoveMin()

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

