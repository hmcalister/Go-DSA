package linkedlist_test

import (
	"slices"
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

// Generic function to test adding items to list of different data types
func genericTestAdd[T any](t *testing.T, list *linkedlist.LinkedList[T], items []T) {
	for index, item := range items {
		list.Add(item)

		expectedLength := index + 1
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected length %v", list.Length(), expectedLength)
		}
	}
}

// Test adding items of different data types to a list
func TestGenericAdd(t *testing.T) {
	t.Run("testAdd int", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[int](), []int{1, 2, 3})
	})

	t.Run("testAdd float", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[float64](), []float64{1.0, 2.0, 3.0})
	})

	t.Run("testAdd string", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[string](), []string{"a", "b", "c"})
	})
}

// Test adding items to a list.
//
// Ensures list length is correct.
func TestListLengthOnAdd(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for index, item := range items {
		list.Add(item)

		expectedLength := index + 1
		if list.Length() != expectedLength {
			t.Errorf("list length (%v) does not match expected list length (%v)", list.Length(), expectedLength)
		}
	}
}

func TestAddAtIndex(t *testing.T) {
	items := []int{10, 20, 30, 40, 50}
	newItem := 0

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
			t.Errorf("removed item (%v) does not match inserted item (%v)", retrievedItem, newItem)
		}
		if err != nil {
			t.Errorf("error when getting item at new index: %v", err)
		}

		expectedLength := len(items) + 1
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected list length %v", list.Length(), expectedLength)
		}
	}

	t.Run("remove at head index", func(t *testing.T) {
		addHelper(t, 0)
	})

	t.Run("remove at tail index", func(t *testing.T) {
		addHelper(t, len(items)-1)
	})

	t.Run("remove at middle index", func(t *testing.T) {
		addHelper(t, 2)
	})

}
