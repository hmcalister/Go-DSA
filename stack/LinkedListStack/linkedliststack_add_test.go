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

func TestLinkedListStackCheckFindAfterAdd(t *testing.T) {
	stack := linkedliststack.New[int]()

	targetItem := 1
	stack.Add(targetItem)
	item, err := stack.Find(func(item int) bool { return item == targetItem })
	if err != nil {
		t.Errorf("found error (%v) after finding from stack that should have item", err)
	}
	if item != targetItem {
		t.Errorf("found item (%v) does not match expected item (%v)", item, targetItem)
	}
}

func TestLinkedListStackCheckFindOfNotPresentItem(t *testing.T) {
	stack := linkedliststack.New[int]()
	stack.Add(1)

	targetItem := 10
	_, err := stack.Find(func(item int) bool { return item == targetItem })
	if err == nil {
		t.Errorf("found nil error after finding from stack without item")
	}
}

func TestLinkedListStackCheckMultipleFindAfterAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := linkedliststack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	expectedItems := []int{4, 2}
	foundItems := stack.FindAll(func(item int) bool { return item%2 == 0 })
	for index := range expectedItems {
		if foundItems[index] != expectedItems[index] {
			t.Errorf("found item (%v) does not match expected item (%v)", foundItems[index], expectedItems[index])
		}
	}
}
