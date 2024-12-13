package linkedlist_test

import (
	"fmt"
	"slices"
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/list/LinkedList"
)

func TestForwardApply(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	concatString := ""
	linkedlist.ForwardApply(list, func(item string) { concatString += item })
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
	linkedlist.ReverseApply(list, func(item string) { concatString += item })

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
	linkedlist.ForwardMap(list, func(item string) string {
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

func TestReverseMap(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	globalCounter := 0
	linkedlist.ReverseMap(list, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})

	for index := range items {
		item, _ := list.ItemAtIndex(index)
		expectedItem := fmt.Sprintf("%v, %v", items[index], len(items)-index-1)
		if item != expectedItem {
			t.Errorf("found item (%v) does not match expected item (%v)", item, expectedItem)
		}
	}
}

func TestForwardFold(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	concatString := linkedlist.ForwardFold(list, "", func(item string, accumulator string) string {
		return accumulator + item
	})
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestReverseFold(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		list.Add(item)
	}

	concatString := linkedlist.ReverseFold(list, "", func(item string, accumulator string) string {
		return accumulator + item
	})

	slices.Reverse(items)
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestForwardIterator(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		list.Add(item)
	}

	sum := 0
	for index, item := range list.ForwardIterator() {
		sum += index * item
	}

	expectedSum := 0
	for index, item := range items {
		expectedSum += index * item
	}

	if sum != expectedSum {
		t.Errorf("result (%v) does not match expected result (%v)", sum, expectedSum)
	}
}

func TestReverseIterator(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		list.Add(item)
	}

	sum := 0
	for index, item := range list.ReverseIterator() {
		sum += index * item
	}

	slices.Reverse(items)
	expectedSum := 0
	for index, item := range items {
		expectedSum += index * item
	}

	if sum != expectedSum {
		t.Errorf("result (%v) does not match expected result (%v)", sum, expectedSum)
	}
}
