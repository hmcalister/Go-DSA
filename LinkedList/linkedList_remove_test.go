package linkedlist_test

import (
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

func genericTestAddThenRemove[T comparable](t *testing.T, list *linkedlist.LinkedList[T], items []T) {
	for _, item := range items {
		list.Add(item)
	}

	// reverse the items
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}

	for index, item := range items {
		retrievedItem, err := list.Remove()
		if err != nil {
			t.Errorf("error when removing item from list: %v", err)
		}
		if retrievedItem != item {
			t.Errorf("removed item (%v) does not match inserted item (%v)", retrievedItem, item)
		}

		expectedLength := len(items) - 1 - index
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected length %v", list.Length(), expectedLength)
		}
	}
}

func TestGenericAddThenRemove(t *testing.T) {
	t.Run("testAddThenRemove int", func(t *testing.T) {
		genericTestAddThenRemove(t, linkedlist.New[int](), []int{1, 2, 3})
	})

	t.Run("testAddThenRemove float", func(t *testing.T) {
		genericTestAddThenRemove(t, linkedlist.New[float64](), []float64{1.0, 2.0, 3.0})
	})

	t.Run("testAddThenRemove string", func(t *testing.T) {
		genericTestAddThenRemove(t, linkedlist.New[string](), []string{"a", "b", "c"})
	})
}

func TestRemoveFromEmptyList(t *testing.T) {
	list := linkedlist.New[int]()
	_, err := list.Remove()
	if err == nil {
		t.Errorf("expected error when removing from empty list")
	}

	list.Add(1)
	list.Remove()
	_, err = list.Remove()
	if err == nil {
		t.Errorf("found error when removing from empty list that previously had items")
	}
}

func TestRemoveFromSingleItemList(t *testing.T) {
	list := linkedlist.New[int]()
	storedItem := 1
	list.Add(storedItem)

	retrievedItem, err := list.Remove()
	if err != nil {
		t.Errorf("error when removing item from list: %v", err)
	}
	if retrievedItem != storedItem {
		t.Errorf("removed item (%v) does not match inserted item (%v)", retrievedItem, storedItem)
	}
}

func TestMultipleAddsAndRemoves(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50}

	for _, item := range items {
		list.Add(item)

		if list.Length() != 1 {
			t.Errorf("list length should be one")
		}

		retrievedItem, err := list.Remove()
		if err != nil {
			t.Errorf("error when removing item from list: %v", err)
		}
		if retrievedItem != item {
			t.Errorf("removed item (%v) does not match inserted item (%v)", retrievedItem, item)
		}
		if list.Length() != 0 {
			t.Errorf("list length should be zero")
		}
	}
}

func TestRemoveAtIndex(t *testing.T) {
	items := []int{10, 20, 30, 40, 50}

	removeHelper := func(t *testing.T, removeIndex int) {
		list := linkedlist.New[int]()

		for _, item := range items {
			list.Add(item)
		}

		retrievedItem, err := list.RemoveAtIndex(removeIndex)
		if err != nil {
			t.Errorf("error when removing item from list: %v", err)
		}
		if retrievedItem != items[removeIndex] {
			t.Errorf("removed item (%v) does not match inserted item (%v)", retrievedItem, items[removeIndex])
		}

		expectedLength := len(items) - 1
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected list length %v", list.Length(), expectedLength)
		}
	}

	t.Run("remove at head index", func(t *testing.T) {
		removeHelper(t, 0)
	})

	t.Run("remove at tail index", func(t *testing.T) {
		removeHelper(t, len(items)-1)
	})

	t.Run("remove at middle index", func(t *testing.T) {
		removeHelper(t, 2)
	})

}
