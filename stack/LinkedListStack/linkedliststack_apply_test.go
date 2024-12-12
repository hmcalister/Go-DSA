package linkedliststack_test

import (
	"fmt"
	"slices"
	"testing"

	linkedliststack "github.com/hmcalister/Go-DSA/stack/LinkedListStack"
)

func TestForwardApply(t *testing.T) {
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	linkedliststack.ForwardApply(queue, func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestReverseApply(t *testing.T) {
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	linkedliststack.ReverseApply(queue, func(item string) { concatString += item })

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
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	linkedliststack.ForwardMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestReverseMap(t *testing.T) {
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	linkedliststack.ReverseMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})

}

func TestForwardFold(t *testing.T) {
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := linkedliststack.ForwardFold(queue, "", func(item string, accumulator string) string {
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
	queue := linkedliststack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := linkedliststack.ReverseFold(queue, "", func(item string, accumulator string) string {
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
