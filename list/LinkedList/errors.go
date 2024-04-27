package linkedlist

import (
	"errors"
)

var (
	ErrorIndexOutOfBounds = errors.New("index out of bounds for list")
	ErrorEmptyList        = errors.New("cannot perform operation on an empty list")
	ErrorItemNotFound     = errors.New("item not found in list")
)
