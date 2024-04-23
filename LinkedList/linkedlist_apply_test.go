package linkedlist_test

import (
	"fmt"
	"slices"
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

func TestForwardApply(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	concatString := ""
	list.ForwardApply(func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestReverseApply(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	concatString := ""
	list.ReverseApply(func(item string) { concatString += item })

	slices.Reverse(items)
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestForwardMap(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	globalCounter := 0
	list.ForwardMap(func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})

	for index := range items {
		item, _ := list.ItemAtIndex(index)
		expectedItem := fmt.Sprintf("%v, %v", items[index], index)
		if item != expectedItem {
			t.Errorf("found item (%v) does not match expected item (%v)", item, expectedItem)
		}
	}
}

