package redblacktree_test

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	redblacktree "github.com/hmcalister/Go-DSA/RedBlackTree"
)

func TestAddItems(t *testing.T) {
	addItemsHelper := func(t *testing.T, items []int) {
		tree := redblacktree.New(comparator.DefaultIntegerComparator)
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

func TestRootAfterAddItems(t *testing.T) {
	rootAfterAddHelper := func(t *testing.T, items []int, rootAfterItemInsertMap map[int]int) {
		tree := redblacktree.New(comparator.DefaultIntegerComparator)

		for _, item := range items {
			err := tree.Add(item)
			if err != nil {
				t.Errorf("error (%v) occurred during insertion of unique item", err)
			}

			expectedRoot := rootAfterItemInsertMap[item]
			foundRoot := tree.Root().Item()
			if expectedRoot != foundRoot {
				t.Errorf("found root item (%v) does not match expected root item (%v)", foundRoot, expectedRoot)
			}
		}
	}
	t.Run("add increasing item", func(t *testing.T) {
		rootAfterAddHelper(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, map[int]int{
			1: 1,
			2: 1,
			3: 2,
			4: 2,
			5: 2,
			6: 2,
			7: 2,
			8: 4,
		})
	})

	t.Run("add decreasing item", func(t *testing.T) {
		rootAfterAddHelper(t, []int{8, 7, 6, 5, 4, 3, 2, 1}, map[int]int{
			8: 8,
			7: 8,
			6: 7,
			5: 7,
			4: 7,
			3: 7,
			2: 7,
			1: 5,
		})
	})

	t.Run("add alternating item", func(t *testing.T) {
		rootAfterAddHelper(t, []int{4, 3, 5, 2, 6, 1, 7, 8}, map[int]int{
			4: 4,
			3: 4,
			5: 4,
			2: 4,
			6: 4,
			1: 4,
			7: 4,
			8: 4,
		})
	})
}

func TestAddItemsCheckOrdering(t *testing.T) {
	addAndCheckOrderingHelper := func(t *testing.T, items []int) {
		tree := redblacktree.New(comparator.DefaultIntegerComparator)

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

func TestManyRandomInsertsSize(t *testing.T) {
	NUM_TRIALS := 10
	numItems := 1

	for trialIndex := range NUM_TRIALS {
		// Increase number of items by two each trial
		numItems = numItems * 2
		items := make([]int, numItems)
		for i := range numItems {
			items[i] = i
		}
		rand.Shuffle(numItems, func(i, j int) {
			items[i], items[j] = items[j], items[i]
		})

		tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
		for _, item := range items {
			tree.Add(item)
		}

		treeSize := tree.Root().Size()
		if treeSize != numItems {
			t.Errorf("tree root size (%v) does not match the expected size (%v) for trial %v", treeSize, numItems, trialIndex)
		}
	}
}

func TestManyRandomInsertsHeight(t *testing.T) {
	NUM_TRIALS := 10
	numItems := 0

	for trialIndex := range NUM_TRIALS {
		// Increase number of items by two each trial
		numItems = numItems*2 + 1
		items := make([]int, numItems)
		for i := range numItems {
			items[i] = i
		}
		rand.Shuffle(numItems, func(i, j int) {
			items[i], items[j] = items[j], items[i]
		})

		tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
		for _, item := range items {
			tree.Add(item)
		}

		treeHeight := tree.Root().Height()
		maxHeight := 2 * (trialIndex + 1)
		if treeHeight > maxHeight {
			t.Errorf("tree root height (%v) is larger than the expected max height (%v) for trial %v", treeHeight, maxHeight, trialIndex)
		}
	}
}
