package linkedliststack_test

import (
	"slices"
	"testing"

	linkedliststack "github.com/hmcalister/Go-DSA/stack/LinkedListStack"
)

func TestLinkedListStackInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		linkedliststack.New[int]()
	})
	t.Run("linked list float", func(t *testing.T) {
		linkedliststack.New[float64]()
	})
	t.Run("linked list string", func(t *testing.T) {
		linkedliststack.New[string]()
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		linkedliststack.New[S]()
	})
}

func TestCheckPeekOfEmptyLinkedListStack(t *testing.T) {
	stack := linkedliststack.New[int]()

	_, err := stack.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty stack", err)
	}
}

func TestFindFromEmptyLinkedListStack(t *testing.T) {
	stack := linkedliststack.New[int]()

	_, err := stack.Find(func(item int) bool { return item == 1 })
	if err == nil {
		t.Errorf("found nil error after finding from empty stack")
	}
}

func TestFindAllFromEmptyLinkedListStack(t *testing.T) {
	stack := linkedliststack.New[int]()

	items := stack.FindAll(func(item int) bool { return item == 1 })
	if len(items) != 0 {
		t.Errorf("found a non-zero number of items from an empty stack")
	}
}

func TestLinkedListStackItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stack := linkedliststack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}

	retrievedItems := stack.Items()
	for _, item := range items {
		if !slices.Contains(retrievedItems, item) {
			t.Errorf("retrieved items %v does not contain expected item %v", retrievedItems, item)
		}
	}
}
