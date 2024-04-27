package redblacktree

import (
	"errors"
)

var (
	ErrorItemAlreadyPresent  = errors.New("item already present in binary tree")
	ErrorItemNotFound        = errors.New("item not present in binary tree")
	ErrorRotationNotPossible = errors.New("rotation is not possible around node")
)
