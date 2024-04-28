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

