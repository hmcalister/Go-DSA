package redblacktree_test

import (
	"testing"

	redblacktree "github.com/hmcalister/Go-DSA/tree/RedBlackTree"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

func TestInitializeTreeGenericTypes(t *testing.T) {
	t.Run("rbt int", func(t *testing.T) {
		redblacktree.New[int](comparator.DefaultIntegerComparator)
	})

	t.Run("rbt float", func(t *testing.T) {
		redblacktree.New[float64](comparator.DefaultFloat64Comparator)
	})

	t.Run("rbt string", func(t *testing.T) {
		redblacktree.New[string](comparator.DefaultStringComparator)
	})

	type S struct {
		i int
		_ float64
		_ string
	}
	t.Run("rbt struct", func(t *testing.T) {
		redblacktree.New[S](func(a, b S) int {
			if a.i < b.i {
				return -1
			} else if a.i > b.i {
				return 1
			}
			return 0
		})
	})
}
