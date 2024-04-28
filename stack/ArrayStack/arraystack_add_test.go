package arraystack_test

import (
	"testing"

	arraystack "github.com/hmcalister/Go-DSA/stack/ArrayStack"
)

func TestArrayStackAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	stack := arraystack.New[int]()

	for _, item := range items {
		stack.Add(item)
	}
}

