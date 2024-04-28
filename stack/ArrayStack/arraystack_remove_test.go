package arraystack_test

import (
	"slices"
	"testing"

	arraystack "github.com/hmcalister/Go-DSA/stack/ArrayStack"
)

func TestArrayStackRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

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

func TestArrayStackRemoveFromEmptyStack(t *testing.T) {
	stack := arraystack.New[int]()

	_, err := stack.Remove()
	if err == nil {
		t.Errorf("did not encounter error (%v) when removing from empty stack", err)
	}
}

func TestArrayStackCheckPeekAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	slices.Reverse(items)

	for _, item := range items {
		peekItem, err := stack.Peek()
		if err != nil {
			t.Errorf("encountered error (%v) after peeking at non-empty stack", err)
		}

		expectedItem := item
		if peekItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", peekItem, expectedItem)
		}

		_, err = stack.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty stack", err)
		}
	}
}

func TestArrayStackCheckRemovedItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	slices.Reverse(items)

	for _, item := range items {
		removedItem, err := stack.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty stack", err)
		}

		expectedItem := item
		if removedItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", removedItem, expectedItem)
		}

	}
}

func TestArrayStackCheckSizeAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	for index := range items {
		stackSize := stack.Size()
		expectedSize := len(items) - index
		if stackSize != expectedSize {
			t.Errorf("found stack size (%v) does not match the expected size (%v)", stackSize, expectedSize)
		}

		_, err := stack.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty stack", err)
		}
	}
}
func TestArrayStackCheckFindAfterRemove(t *testing.T) {
	stack := arraystack.New[int]()

	targetItem := 1
	stack.Add(targetItem)
	item, err := stack.Find(func(item int) bool { return item == targetItem })
	if err != nil {
		t.Errorf("found error (%v) after finding from stack that should have item", err)
	}
	if item != targetItem {
		t.Errorf("found item (%v) does not match expected item (%v)", item, targetItem)
	}

	stack.Remove()
	_, err = stack.Find(func(item int) bool { return item == targetItem })
	if err == nil {
		t.Errorf("found nil error after finding from stack without item")
	}
}

func TestArrayStackCheckFindAllAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}
	for range len(items) - 1 {
		stack.Remove()
	}

	foundItems := stack.FindAll(func(item int) bool { return item%2 == 0 })
	if len(foundItems) != 0 {
		t.Errorf("found a non-zero number of items from a stack with expected zero number of matches")
	}
}
