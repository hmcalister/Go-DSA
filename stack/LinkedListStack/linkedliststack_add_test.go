package linkedliststack_test

import (
	"testing"

	linkedliststack "github.com/hmcalister/Go-DSA/stack/LinkedListStack"
)

func TestLinkedListStackAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := linkedliststack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}
}

