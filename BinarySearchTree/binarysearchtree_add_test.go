package binarysearchtree_test

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	binarysearchtree "github.com/hmcalister/Go-DSA/BinarySearchTree"
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

func TestAddItems(t *testing.T) {
	addItemsHelper := func(t *testing.T, items []int) {
		tree := binarysearchtree.New(comparator.DefaultIntegerComparator)
		for _, item := range items {
			err := tree.Add(item)
			if err != nil {
				t.Errorf("error (%v) occurred during insertion of unique item", err)
			}
		}
	}
	t.Run("add increasing item", func(t *testing.T) {
		addItemsHelper(t, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("add decreasing item", func(t *testing.T) {
		addItemsHelper(t, []int{7, 6, 5, 4, 3, 2, 1})
	})

	t.Run("add alternating item", func(t *testing.T) {
		addItemsHelper(t, []int{4, 5, 3, 6, 2, 7, 1})
	})
