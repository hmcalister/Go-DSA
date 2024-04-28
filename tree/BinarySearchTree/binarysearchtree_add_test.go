package binarysearchtree_test

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	binarysearchtree "github.com/hmcalister/Go-DSA/tree/BinarySearchTree"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
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
		binarysearchtree.ApplyTreeInorder(tree, func(item int) {
			inorderTraversal += fmt.Sprintf("%d,", item)
		})

		if expectedInorderTraversal != inorderTraversal {
			t.Errorf("inorder traversal (%v) does not match expected inorder traversal (%v)", inorderTraversal, expectedInorderTraversal)
		}
	}

	t.Run("check ordering increasing item", func(t *testing.T) {
		addAndCheckOrderingHelper(t, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("check ordering decreasing item", func(t *testing.T) {
		addAndCheckOrderingHelper(t, []int{7, 6, 5, 4, 3, 2, 1})
	})

	t.Run("check ordering alternating item", func(t *testing.T) {
		addAndCheckOrderingHelper(t, []int{4, 5, 3, 6, 2, 7, 1})
	})

	t.Run("check ordering many items random order", func(t *testing.T) {
		numItems := 100
		items := make([]int, numItems)
		for i := range numItems {
			items[i] = i
		}
		rand.Shuffle(numItems, func(i, j int) {
			items[i], items[j] = items[j], items[i]
		})
		addAndCheckOrderingHelper(t, items)
	})
}

func TestSizeAfterAdd(t *testing.T) {
	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9

	items := []int{5, 3, 7, 1, 4, 6, 9}
	itemSizeMap := map[int]int{
		5: 7,
		3: 3,
		7: 3,
		1: 1,
		4: 1,
		6: 1,
		9: 1,
	}

	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	for item, expectedSize := range itemSizeMap {
		node, err := tree.Find(item)
		if err != nil {
			t.Errorf("error (%v) encountered when finding item that was inserted into tree", err)
		}

		if node.Size() != expectedSize {
			t.Errorf("found size (%v) does not match expected size (%v) for item %v", node.Size(), expectedSize, item)
		}
	}
}

func TestHeightAfterAdd(t *testing.T) {
	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9

	items := []int{5, 3, 7, 1, 4, 6, 9}
	itemHeightMap := map[int]int{
		5: 2,
		3: 1,
		7: 1,
		1: 0,
		4: 0,
		6: 0,
		9: 0,
	}

	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	for item, expectedHeight := range itemHeightMap {
		node, err := tree.Find(item)
		if err != nil {
			t.Errorf("error (%v) encountered when finding item that was inserted into tree", err)
		}

		if node.Height() != expectedHeight {
			t.Errorf("found height (%v) does not match expected height (%v) for item %v", node.Height(), expectedHeight, item)
		}
	}
}
