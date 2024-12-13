package binarysearchtree_test

import (
	"iter"
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

func TestBinarySearchTreeIterators(t *testing.T) {
	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9

	items := []int{5, 3, 7, 1, 4, 6, 9}
	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	testNodeIteratorMethod := func(t *testing.T, iteratorMethod func(*binarysearchtree.BinarySearchTreeNode[int]) iter.Seq[int], iteratorDescriptor string, expectedOrder []int) {
		foundOrder := make([]int, 0)
		for item := range iteratorMethod(tree.Root()) {
			foundOrder = append(foundOrder, item)
		}

		if !slices.Equal(expectedOrder, foundOrder) {
			t.Errorf("%v iterator: expected order %v does not match found order %v", iteratorDescriptor, expectedOrder, foundOrder)
		}
	}

	testTreeIteratorMethod := func(t *testing.T, iteratorMethod func(*binarysearchtree.BinarySearchTree[int]) iter.Seq[int], iteratorDescriptor string, expectedOrder []int) {
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
	testNodeIteratorMethod(t, binarysearchtree.IteratorNodePreorder, "node preorder", expectedOrder)
	testTreeIteratorMethod(t, binarysearchtree.IteratorTreePreorder, "tree preorder", expectedOrder)

	// In-Order
	expectedOrder = []int{1, 3, 4, 5, 6, 7, 9}
	testNodeIteratorMethod(t, binarysearchtree.IteratorNodeInorder, "node inorder", expectedOrder)
	testTreeIteratorMethod(t, binarysearchtree.IteratorTreeInorder, "tree inorder", expectedOrder)

	// Post-Order
	expectedOrder = []int{1, 4, 3, 6, 9, 7, 5}
	testNodeIteratorMethod(t, binarysearchtree.IteratorNodePostorder, "node postorder", expectedOrder)
	testTreeIteratorMethod(t, binarysearchtree.IteratorTreePostorder, "tree postorder", expectedOrder)
}

func TestBinarySearchTreeLargeIterator(t *testing.T) {
	const MAX_ITEM = 4096
	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	for i := 0; i < MAX_ITEM; i += 1 {
		tree.Add(i)
	}

	expectedItem := 0
	for item := range binarysearchtree.IteratorNodeInorder(tree.Root()) {
		if expectedItem != item {
			t.Errorf("expected item %v does not match found item %v", expectedItem, item)
			return
		}
		expectedItem += 1
	}
}
