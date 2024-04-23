package linkedlist_test

import (
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

func TestItemAtIndex(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	for index, item := range items {
		retrievedItem, err := list.ItemAtIndex(index)

		if err != nil {
			t.Errorf("failed to get item at index %v: %v", index, err)
		}
		if item != retrievedItem {
			t.Errorf("item retrieved (%v) is not item stored (%v) for index %v", retrievedItem, item, index)
		}
	}
}

// Tests adding items to a list at an index, then getting from that index.
func TestGetAtIndexAfterAdd(t *testing.T) {
	items := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	newItem := 0

	// Define a small helper method that creates a new list,
	// inserts the items from the items array, then tries to insert an item
	// at the specified index.
	//
	// Calls t.Errorf if the insert fails or if the list length does not match the expected
	addHelper := func(t *testing.T, addIndex int) {
		list := linkedlist.New[int]()
		for _, item := range items {
			list.Add(item)
		}

		err := list.AddAtIndex(newItem, addIndex)
		if err != nil {
			t.Errorf("error when adding item to list: %v", err)
		}

		retrievedItem, err := list.ItemAtIndex(addIndex)

		if retrievedItem != newItem {
			t.Errorf("retrieved item (%v) does not match inserted item (%v)", retrievedItem, newItem)
		}
		if err != nil {
			t.Errorf("error when getting item at new index: %v", err)
		}
	}

	t.Run("get at head index", func(t *testing.T) {
		addHelper(t, 0)
	})

	t.Run("get at non-head first-half index", func(t *testing.T) {
		addHelper(t, 2)
	})

	t.Run("get at non-head second-half index", func(t *testing.T) {
		addHelper(t, len(items)-3)
	})

	t.Run("get at tail index", func(t *testing.T) {
		addHelper(t, len(items)-1)
	})
}

// Tests adding items to a list at an index.
func TestGetAtIndexAfterRemove(t *testing.T) {
	items := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	// Define a small helper method that creates a new list,
	// inserts the items from the items array, then removes an item from an index.
	//
	// Calls t.Errorf if the remove fails or if the list length does not match the expected
	removeHelper := func(t *testing.T, removeIndex int) {
		list := linkedlist.New[int]()
		for _, item := range items {
			list.Add(item)
		}

		removedItem, err := list.RemoveAtIndex(removeIndex)
		if err != nil {
			t.Errorf("error when adding item to list: %v", err)
		}
		if removedItem != items[removeIndex] {
			t.Errorf("removed item %v does not match the expected removed item %v at index %v", removedItem, items[removeIndex], removeIndex)
		}

		retrievedItem, err := list.ItemAtIndex(removeIndex)

		if retrievedItem != items[removeIndex+1] {
			t.Errorf("retrieved item (%v) does not match inserted item (%v)", retrievedItem, items[removeIndex+1])
		}
		if err != nil {
			t.Errorf("error when getting item at new index: %v", err)
		}
	}

	t.Run("remove at head index", func(t *testing.T) {
		removeHelper(t, 0)
	})

	t.Run("remove at non-head first-half index", func(t *testing.T) {
		removeHelper(t, 2)
	})

	t.Run("remove at non-head second-half index", func(t *testing.T) {
		removeHelper(t, len(items)-3)
	})

	// We cannot remove and test getting at the tail, as the list will shrink
	// 	t.Run("remove at tail index", func(t *testing.T) {
	// 		removeHelper(t, len(items)-1)
	// 	})
}

// Test Find
func TestFind(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	item, err := list.Find(func(item int) bool { return item > 20 })
	if err != nil {
		t.Errorf("error encountered when finding item")
	}
	if item != 30 {
		t.Errorf("found item (%v) does not match expected item (%v)", item, 30)
	}
}

func TestFindFail(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	_, err := list.Find(func(item int) bool { return item < 0 })
	if err == nil {
		t.Errorf("expected error during finding item, no error found")
	}
}

func TestReverseFind(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	item, err := list.ReverseFind(func(item int) bool { return item > 20 })
	if err != nil {
		t.Errorf("error encountered when finding item")
	}
	if item != 80 {
		t.Errorf("found item (%v) does not match expected item (%v)", item, 80)
	}
}

func TestReverseFindFail(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	_, err := list.ReverseFind(func(item int) bool { return item < 0 })
	if err == nil {
		t.Errorf("expected error during finding item, no error found")
	}
}

func TestFindAll(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, item := range items {
		list.Add(item)
	}

	foundItems := list.FindAll(func(item int) bool { return item > 20 })
	if len(foundItems) != 6 {
		t.Errorf("length of found items (%v) does not match expected length %v", len(foundItems), 6)
	}
	for _, item := range foundItems {
		if item <= 20 {
			t.Errorf("found item (%v) does not match predicate", item)
		}
	}
}

