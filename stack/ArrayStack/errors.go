package arraystack

import "errors"

var (
	ErrorStackEmpty   = errors.New("stack is empty")
	ErrorItemNotFound = errors.New("item not found in stack")
)
