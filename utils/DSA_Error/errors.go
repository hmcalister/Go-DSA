package dsa_error

import (
	"errors"
)

var (
	ErrorItemAlreadyPresent = errors.New("item already present in data structure")
	ErrorItemNotFound       = errors.New("item not present in data structure")
	ErrorDataStructureEmpty = errors.New("data structure is empty")
	ErrorIndexOutOfBounds   = errors.New("index out of bounds")
)
