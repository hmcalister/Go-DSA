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

