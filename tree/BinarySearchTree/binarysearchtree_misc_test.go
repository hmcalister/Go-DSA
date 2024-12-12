package binarysearchtree_test

import (
	"slices"
	"testing"

	binarysearchtree "github.com/hmcalister/Go-DSA/tree/BinarySearchTree"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
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

func TestBinarySearchTreeItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		tree.Add(item)
	}

	retrievedItems := tree.Items()
	for _, item := range items {
		if !slices.Contains(retrievedItems, item) {
			t.Errorf("retrieved items %v does not contain expected item %v", retrievedItems, item)
		}
	}
}
