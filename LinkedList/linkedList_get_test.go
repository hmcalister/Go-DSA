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

