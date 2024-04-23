package binarysearchtree

import "fmt"

// Error for if item already exists in the tree

type ItemAlreadyPresentError[T any] struct {
	item T
}

func (e *ItemAlreadyPresentError[T]) Error() string {
	return fmt.Sprintf("item %#v already present in binary tree", e.item)
}

