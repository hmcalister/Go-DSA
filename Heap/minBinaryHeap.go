package heap

import (
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

type MinBinaryHeap[T any] struct {
	heapData           []T
	comparatorFunction comparator.ComparatorFunction[T]
}

