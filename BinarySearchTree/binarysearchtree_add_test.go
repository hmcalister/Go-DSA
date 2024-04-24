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

	t.Run("add many items random order", func(t *testing.T) {
		numItems := 100
		items := make([]int, numItems)
		for i := range numItems {
			items[i] = i
		}
		rand.Shuffle(numItems, func(i, j int) {
			items[i], items[j] = items[j], items[i]
		})
		addItemsHelper(t, items)
	})
}

func TestAddItemsCheckOrdering(t *testing.T) {
	addAndCheckOrderingHelper := func(t *testing.T, items []int) {
		tree := binarysearchtree.New(comparator.DefaultIntegerComparator)

		for _, item := range items {
			tree.Add(item)
		}

		// An in-order traversal should give the items in a sorted order
		slices.Sort(items)
		expectedInorderTraversal := ""
		for _, item := range items {
			expectedInorderTraversal += fmt.Sprintf("%d,", item)
		}
		inorderTraversal := ""
		tree.ApplyTreeInorder(func(item int) {
			inorderTraversal += fmt.Sprintf("%d,", item)
		})

		if expectedInorderTraversal != inorderTraversal {
			t.Errorf("inorder traversal (%v) does not match expected inorder traversal (%v)", inorderTraversal, expectedInorderTraversal)
		}
	}

