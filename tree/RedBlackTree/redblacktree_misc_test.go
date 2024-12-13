package redblacktree_test

import (
	"iter"
	"slices"
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

func TestRedBlackTreeItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)

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

func TestRedBlackTreeIterators(t *testing.T) {
	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9

	items := []int{5, 3, 7, 1, 4, 6, 9}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	testNodeIteratorMethod := func(t *testing.T, iteratorMethod func(*redblacktree.RedBlackTreeNode[int]) iter.Seq[int], iteratorDescriptor string, expectedOrder []int) {
		foundOrder := make([]int, 0)
		for item := range iteratorMethod(tree.Root()) {
			foundOrder = append(foundOrder, item)
		}

		if !slices.Equal(expectedOrder, foundOrder) {
			t.Errorf("%v iterator: expected order %v does not match found order %v", iteratorDescriptor, expectedOrder, foundOrder)
		}
	}

	testTreeIteratorMethod := func(t *testing.T, iteratorMethod func(*redblacktree.RedBlackTree[int]) iter.Seq[int], iteratorDescriptor string, expectedOrder []int) {
		foundOrder := make([]int, 0)
		for item := range iteratorMethod(tree) {
			foundOrder = append(foundOrder, item)
		}

		if !slices.Equal(expectedOrder, foundOrder) {
			t.Errorf("%v iterator: expected order %v does not match found order %v", iteratorDescriptor, expectedOrder, foundOrder)
		}
	}

	var expectedOrder []int

	// Pre-Order
	expectedOrder = []int{5, 3, 1, 4, 7, 6, 9}
	testNodeIteratorMethod(t, redblacktree.IteratorNodePreorder, "node preorder", expectedOrder)
	testTreeIteratorMethod(t, redblacktree.IteratorTreePreorder, "tree preorder", expectedOrder)

	// In-Order
	expectedOrder = []int{1, 3, 4, 5, 6, 7, 9}
	testNodeIteratorMethod(t, redblacktree.IteratorNodeInorder, "node inorder", expectedOrder)
	testTreeIteratorMethod(t, redblacktree.IteratorTreeInorder, "tree inorder", expectedOrder)

	// Post-Order
	expectedOrder = []int{1, 4, 3, 6, 9, 7, 5}
	testNodeIteratorMethod(t, redblacktree.IteratorNodePostorder, "node postorder", expectedOrder)
	testTreeIteratorMethod(t, redblacktree.IteratorTreePostorder, "tree postorder", expectedOrder)
}

func TestRedBlackTreeLargeIterator(t *testing.T) {
	const MAX_ITEM = 4096
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for i := 0; i < MAX_ITEM; i += 1 {
		tree.Add(i)
	}

	expectedItem := 0
	for item := range redblacktree.IteratorNodeInorder(tree.Root()) {
		if expectedItem != item {
			t.Errorf("expected item %v does not match found item %v", expectedItem, item)
			return
		}
		expectedItem += 1
	}
}
