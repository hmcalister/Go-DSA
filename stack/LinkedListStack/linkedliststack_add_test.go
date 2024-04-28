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

func TestLinkedListStackCheckPeekAfterAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := linkedliststack.New[int]()

	for _, item := range items {
		stack.Add(item)

		peekItem, err := stack.Peek()
		if err != nil {
			t.Errorf("encountered error (%v) after peeking at non-empty stack", err)
		}

		expectedItem := item
		if peekItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", peekItem, expectedItem)
		}
	}
}

func TestLinkedListStackCheckSizeAfterAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := linkedliststack.New[int]()

	for index, item := range items {
		stack.Add(item)

		stackSize := stack.Size()
		expectedSize := index + 1
		if stackSize != expectedSize {
			t.Errorf("found stack size (%v) does not match the expected size (%v)", stackSize, expectedSize)
		}
	}
}
