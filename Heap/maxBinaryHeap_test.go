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

func TestMaxHeapIntInit(t *testing.T) {
	heap.NewMaxBinaryHeap[int](comparator.DefaultIntegerComparator)
}

func TestMaxHeapFloatInit(t *testing.T) {
	heap.NewMaxBinaryHeap[float64](comparator.DefaultFloat64Comparator)
}

func TestMaxHeapStringInit(t *testing.T) {
	heap.NewMaxBinaryHeap[string](comparator.DefaultStringComparator)
}

func TestMaxHeapStructInit(t *testing.T) {
	type S struct {
		_ int
		f float64
	}
	heap.NewMaxBinaryHeap[S](func(a, b S) int {
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
	heap := heap.NewMaxBinaryHeap[int](comparator.DefaultIntegerComparator)

	_, err := heap.RemoveMax()
	if err == nil {
		t.Errorf("got nil error when removing from empty heap")
	}
}

func TestMaxHeapEmptySize(t *testing.T) {
	heap := heap.NewMaxBinaryHeap[int](comparator.DefaultIntegerComparator)

	expectedSize := 0
	heapSize := heap.Size()
	if heapSize != expectedSize {
		t.Errorf("heap size (%v) does not match expected size (%v)", heapSize, expectedSize)
	}
}

