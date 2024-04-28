package arraystack_test

import (
	"testing"

	arraystack "github.com/hmcalister/Go-DSA/stack/ArrayStack"
)

func TestArrayStackInit(t *testing.T) {
	t.Run("array stack int", func(t *testing.T) {
		arraystack.New[int]()
	})
	t.Run("array stack float", func(t *testing.T) {
		arraystack.New[float64]()
	})
	t.Run("array stack string", func(t *testing.T) {
		arraystack.New[string]()
	})
	t.Run("array stack struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		arraystack.New[S]()
	})
}

func TestCheckPeekOfEmptyArrayStack(t *testing.T) {
	stack := arraystack.New[int]()

	_, err := stack.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty stack", err)
	}
}

func TestFindFromEmptyArrayStack(t *testing.T) {
	stack := arraystack.New[int]()

	_, err := stack.Find(func(item int) bool { return item == 1 })
	if err == nil {
		t.Errorf("found nil error after finding from empty stack")
	}
}

func TestFindAllFromEmptyArrayStack(t *testing.T) {
	stack := arraystack.New[int]()

	items := stack.FindAll(func(item int) bool { return item == 1 })
	if len(items) != 0 {
		t.Errorf("found a non-zero number of items from an empty stack")
	}
}
