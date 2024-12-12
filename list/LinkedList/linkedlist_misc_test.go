package linkedlist_test

import (
	"slices"
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
)

func TestLinkedListInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		linkedlist.New[int]()
	})
	t.Run("linked list float", func(t *testing.T) {
		linkedlist.New[float64]()
	})
	t.Run("linked list string", func(t *testing.T) {
		linkedlist.New[string]()
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		linkedlist.New[S]()
	})
}

func TestAddAtIndexDocumentationExample(t *testing.T) {
	list := linkedlist.New[string]()
	list.Add("hello")               // list = ["hello"]
	list.Add("world")               // list = ["hello", "world"]
	list.AddAtIndex("(linked!)", 1) // list = ["hello", "(linked!)", "world"]

	if list.Length() != 3 {
		t.Errorf("list length does not match expected")
	}

	item, err := list.ItemAtIndex(1)
	if err != nil {
		t.Errorf("error during remove at index")
	}
	if item != "(linked!)" {
		t.Errorf("retrieved item does not match expected item")
	}
}

func TestLinkedListItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := linkedlist.New[int]()

	for _, item := range items {
		set.Add(item)
	}

	retrievedItems := set.Items()
	for _, item := range items {
		if !slices.Contains(retrievedItems, item) {
			t.Errorf("retrieved items %v does not contain expected item %v", retrievedItems, item)
		}
	}
}

func TestRemoveAtIndexDocumentationExample(t *testing.T) {
	list := linkedlist.New[string]()
	list.Add("hello")     // list = ["hello"]
	list.Add("(linked!)") // list = ["hello", "(linked!)"]
	list.Add("world")     // list = ["hello", "(linked!)", "world"]

	item, err := list.RemoveAtIndex(1) // list = ["hello", "world"]
	if err != nil {
		t.Errorf("error during remove at index")
	}
	if item != "(linked!)" {
		t.Errorf("retrieved item does not match expected item")
	}

	if list.Length() != 2 {
		t.Errorf("list length does not match expected")
	}
}
