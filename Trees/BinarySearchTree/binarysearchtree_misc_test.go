package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/hmcalister/Go-DSA/BinarySearchTree"
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

func TestInitializeTreeGenericTypes(t *testing.T) {
	t.Run("bst int", func(t *testing.T) {
		binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	})

	t.Run("bst float", func(t *testing.T) {
		binarysearchtree.New[float64](comparator.DefaultFloat64Comparator)
	})

	t.Run("bst string", func(t *testing.T) {
		binarysearchtree.New[string](comparator.DefaultStringComparator)
	})

	type S struct {
		i int
		_ float64
		_ string
	}
	t.Run("bst struct", func(t *testing.T) {
		binarysearchtree.New[S](func(a, b S) int {
			if a.i < b.i {
				return -1
			} else if a.i > b.i {
				return 1
			}
			return 0
		})
	})
}
