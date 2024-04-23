package linkedlist

import "fmt"

// ----------------------------------------------------------------------------
type IndexOutOfBoundsError struct {
	targetIndex int
	listLength  int
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %d out of bound for list of length %d", e.targetIndex, e.listLength)
}

// ----------------------------------------------------------------------------
type EmptyListError struct {
}

func (e *EmptyListError) Error() string {
	return "cannot perform operation on an empty list"
}

// ----------------------------------------------------------------------------
