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

