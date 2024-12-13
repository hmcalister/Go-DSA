package arrayqueue_test

import (
	"fmt"
	"slices"
	"testing"

	arrayqueue "github.com/hmcalister/Go-DSA/queue/ArrayQueue"
)

func TestForwardApply(t *testing.T) {
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	arrayqueue.ForwardApply(queue, func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}

func TestReverseApply(t *testing.T) {
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	arrayqueue.ReverseApply(queue, func(item string) { concatString += item })

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
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	arrayqueue.ForwardMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestReverseMap(t *testing.T) {
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	arrayqueue.ReverseMap(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})

}

func TestForwardFold(t *testing.T) {
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := arrayqueue.ForwardFold(queue, "", func(item string, accumulator string) string {
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
	queue := arrayqueue.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := arrayqueue.ReverseFold(queue, "", func(item string, accumulator string) string {
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
	queue := arrayqueue.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		queue.Add(item)
	}

	sum := 0
	for index, item := range queue.ForwardIterator() {
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
	queue := arrayqueue.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		queue.Add(item)
	}

	sum := 0
	for index, item := range queue.ReverseIterator() {
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
