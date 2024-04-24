package redblacktree

import "fmt"

// Error for if item already exists in the tree

type ItemAlreadyPresentError[T any] struct {
	item T
}

func (e *ItemAlreadyPresentError[T]) Error() string {
	return fmt.Sprintf("item %#v already present in binary tree", e.item)
}

// Error for if item does not exists in the tree

type ItemNotFoundError[T any] struct {
	item T
}

func (e *ItemNotFoundError[T]) Error() string {
	return fmt.Sprintf("item %#v not present in binary tree", e.item)
}

// Error for rotation is impossible due to, for example, the node not having the correct children present

type RotationNotPossible[T any] struct {
}

func (e *RotationNotPossible[T]) Error() string {
	return fmt.Sprintf("rotation is not possible around node")
}
