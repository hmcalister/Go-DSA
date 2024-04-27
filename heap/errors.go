package heap

import "errors"

var (
	ErrorEmptyHeap      = errors.New("heap is empty")
	ErrorItemNotPresent = errors.New("item not present in heap")
)
