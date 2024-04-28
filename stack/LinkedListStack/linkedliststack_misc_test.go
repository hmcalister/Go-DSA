package linkedliststack_test

import (
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
