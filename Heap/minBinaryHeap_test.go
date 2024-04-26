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

