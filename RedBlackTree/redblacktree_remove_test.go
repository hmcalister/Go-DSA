package redblacktree_test

import (
	"math/rand"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	redblacktree "github.com/hmcalister/Go-DSA/RedBlackTree"
)

func TestRemoveRootAsOnlynode(t *testing.T) {
	items := []int{1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(1)
	if err != nil {
		t.Errorf("encountered error (%v) when removing root node", err)
	}

	node, err := tree.Find(1)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting root")
	}
}

func TestRemoveRoot(t *testing.T) {
	items := []int{3, 4, 2, 5, 1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing root node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting root")
	}
}

func TestRemoveTwoChildNode(t *testing.T) {
	items := []int{1, 3, 2, 4}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing two child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting two child node")
	}
}

func TestRemoveNodeWithOnlyLeftChild(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing one child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting one child node")
	}
}

func TestRemoveNodeWithOnlyRightChild(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing one child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting one child node")
	}
}

func TestRemoveLeafNodeAsLeftChild(t *testing.T) {
	items := []int{3, 2, 1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(1)
	if err != nil {
		t.Errorf("encountered error (%v) when removing leaf node", err)
	}

	node, err := tree.Find(1)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting leaf node")
	}
}

func TestRemoveLeafNodeAsRightChild(t *testing.T) {
	items := []int{1, 2, 3}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing leaf node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting leaf node")
	}
}

// ----------------------------------------------------------------------------
// Size Tests

func TestSizeAfterRemoval(t *testing.T) {
	// Define a helper function that creates a new tree, removes an item, and tests each remaining item against the expected size map
	testSizeAfterRemovalHelper := func(t *testing.T, items []int, removalItem int, expectedSizeMap map[int]int) {
		tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
		for _, item := range items {
			tree.Add(item)
		}

		err := tree.Remove(removalItem)
		if err != nil {
			t.Errorf("error (%v) encountered when removing item", err)
		}

		for item, expectedSize := range expectedSizeMap {
			node, err := tree.Find(item)
			if err != nil {
				t.Errorf("error (%v) encountered when finding item that was inserted into tree", err)
			}

			if node.Size() != expectedSize {
				t.Errorf("found size (%v) does not match expected size (%v) for item %v", node.Size(), expectedSize, item)
			}
		}
	}

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9
	//
	// And then remove the root, resulting in
	// 				6
	// 			/		\
	// 		  3			  7
	// 		/	\			\
	// 	   1	 4	   		  9
	t.Run("remove root node", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6, 9}
		removalItem := 5
		expectedSizeMap := map[int]int{
			6: 6,
			3: 3,
			7: 2,
			1: 1,
			4: 1,
			9: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/
	// 	   1	 4	   6
	//
	// And then remove 7, resulting in
	// 				3
	// 			/		\
	// 		  1			 5
	// 				   /   \
	// 	       		  4     6
	t.Run("remove node with only left child", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6}
		removalItem := 7
		expectedSizeMap := map[int]int{
			3: 5,
			5: 3,
			1: 1,
			4: 1,
			6: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\			\
	// 	   1	 4	         9
	//
	// And then remove 7, resulting in
	// 				3
	// 			/		\
	// 		  1			 5
	// 				   /   \
	// 	       		  4     9
	t.Run("remove node with only right child", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 9}
		removalItem := 7
		expectedSizeMap := map[int]int{
			3: 5,
			5: 3,
			9: 1,
			1: 1,
			4: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	 9
	//
	// And then remove 7, resulting in
	// 				5
	// 			/		\
	// 		  3			 9
	// 		/	\		/
	// 	   1	 4	   6
	t.Run("remove non-root with two children", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6, 9}
		removalItem := 7
		expectedSizeMap := map[int]int{
			5: 6,
			3: 3,
			9: 2,
			1: 1,
			4: 1,
			6: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	 9
	//
	// And then remove 1, resulting in
	// 				5
	// 			/		\
	// 		  3			  7
	// 			\		/	\
	// 	   		 4	   6	 9
	t.Run("remove leaf as left child", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6, 9}
		removalItem := 1
		expectedSizeMap := map[int]int{
			5: 6,
			3: 2,
			7: 3,
			4: 1,
			6: 1,
			9: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	 9
	//
	// And then remove 4, resulting in
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/			/	\
	// 	   1		   6	 9
	t.Run("remove leaf as right child", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6, 9}
		removalItem := 4
		expectedSizeMap := map[int]int{
			5: 6,
			3: 2,
			7: 3,
			1: 1,
			6: 1,
			9: 1,
		}

		testSizeAfterRemovalHelper(t, items, removalItem, expectedSizeMap)
	})
}

// ----------------------------------------------------------------------------
// Height Tests

func TestHeightAfterRemoval(t *testing.T) {
	// Define a helper function that creates a new tree, removes an item, and tests each remaining item against the expected size map
	testHeightAfterRemovalHelper := func(t *testing.T, items []int, removalItem int, expectedHeightMap map[int]int) {
		tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
		for _, item := range items {
			tree.Add(item)
		}

		err := tree.Remove(removalItem)
		if err != nil {
			t.Errorf("error (%v) encountered when removing item", err)
		}

		for item, expectedHeight := range expectedHeightMap {
			node, err := tree.Find(item)
			if err != nil {
				t.Errorf("error (%v) encountered when finding item that was inserted into tree", err)
			}

			if node.Height() != expectedHeight {
				t.Errorf("found height (%v) does not match expected height (%v) for item %v", node.Size(), expectedHeight, item)
			}
		}
	}

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/	\
	// 	   1	 4	   6	  9
	//
	// And then remove the root, resulting in
	// 				6
	// 			/		\
	// 		  3			  7
	// 		/	\			\
	// 	   1	 4	   		  9
	t.Run("remove root node", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6, 9}
		removalItem := 5
		expectedHeightMap := map[int]int{
			6: 2,
			3: 1,
			7: 1,
			1: 0,
			4: 0,
			9: 0,
		}

		testHeightAfterRemovalHelper(t, items, removalItem, expectedHeightMap)
	})

	// We will construct this tree
	// 				5
	// 			/		\
	// 		  3			  7
	// 		/	\		/
	// 	   1	 4	   6
	//
	// And then remove 7, resulting in
	// 				3
	// 			/		\
	// 		  1			 5
	// 				   /   \
	// 	       		  4     6
	t.Run("remove node with only left child", func(t *testing.T) {
		items := []int{5, 3, 7, 1, 4, 6}
		removalItem := 7
		expectedHeightMap := map[int]int{
			3: 2,
			5: 1,
			1: 0,
			6: 0,
			4: 0,
		}

		testHeightAfterRemovalHelper(t, items, removalItem, expectedHeightMap)
	})

