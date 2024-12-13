package arraystack_test

import (
	"fmt"
	"slices"
	"testing"

	arraystack "github.com/hmcalister/Go-DSA/stack/ArrayStack"
)

func TestForwardApply(t *testing.T) {
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	arraystack.ForwardApply(queue, func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestReverseApply(t *testing.T) {
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	arraystack.ReverseApply(queue, func(item string) { concatString += item })

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
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	arraystack.ForwardMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestReverseMap(t *testing.T) {
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	arraystack.ReverseMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})

}

func TestForwardFold(t *testing.T) {
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := arraystack.ForwardFold(queue, "", func(item string, accumulator string) string {
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
	queue := arraystack.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := arraystack.ReverseFold(queue, "", func(item string, accumulator string) string {
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
	stack := arraystack.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		stack.Add(item)
	}

	sum := 0
	for index, item := range stack.ForwardIterator() {
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
	stack := arraystack.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		stack.Add(item)
	}

	sum := 0
	for index, item := range stack.ReverseIterator() {
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
