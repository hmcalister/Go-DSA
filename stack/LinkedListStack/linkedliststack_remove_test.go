package linkedliststack_test

import (
	"slices"
	"testing"

	linkedliststack "github.com/hmcalister/Go-DSA/stack/LinkedListStack"
)

func TestLinkedListStackRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := linkedliststack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	for range items {
		_, err := stack.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty stack", err)
		}
	}
}

func TestLinkedListStackRemoveFromEmptyStack(t *testing.T) {
	stack := linkedliststack.New[int]()

	_, err := stack.Remove()
	if err == nil {
		t.Errorf("did not encounter error (%v) when removing from empty stack", err)
	}
}

