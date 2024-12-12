package priorityqueue_test

import (
	"fmt"
	"slices"
	"testing"

	priorityqueue "github.com/hmcalister/Go-DSA/queue/PriorityQueue"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

func TestPriorityQueueApply(t *testing.T) {
	queue := priorityqueue.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	concatString := ""
	priorityqueue.Apply(queue, func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}
func TestPriorityQueueMap(t *testing.T) {
	queue := priorityqueue.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	globalCounter := 0
	priorityqueue.Map(queue, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestPriorityQueueMapUpdateHeap(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, item := range items {
		queue.Add(item)
	}

	priorityqueue.Map(queue, func(item int) int {
		newItem := 9 - item
		return newItem
	})
	expectedMinItem := 0
	minItem, _ := queue.Peek()
	if minItem != expectedMinItem {
		t.Errorf("expected minimum item %v does not match found minimum item %v", expectedMinItem, minItem)
	}
}

func TestPriorityQueueFold(t *testing.T) {
	queue := priorityqueue.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		queue.Add(item)
	}

	itemsContained := priorityqueue.Fold(queue, true, func(item string, accumulator bool) bool {
		if !accumulator {
			return accumulator
		}
		return slices.Contains(items, item)
	})

	if !itemsContained {
		t.Errorf("expected all items to be contained in ground truth array")
	}
}
