package minbinaryheap

import "errors"

var (
	ErrorEmptyHeap    = errors.New("heap is empty")
	ErrorItemNotFound = errors.New("item not found in heap")
)
