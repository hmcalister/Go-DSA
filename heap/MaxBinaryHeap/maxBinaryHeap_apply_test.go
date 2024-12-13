package maxbinaryheap_test

import (
	"fmt"
	"slices"
	"testing"

	maxbinaryheap "github.com/hmcalister/Go-DSA/heap/MaxBinaryHeap"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

func TestMaxBinaryHeapApply(t *testing.T) {
	heap := maxbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	concatString := ""
	maxbinaryheap.Apply(heap, func(item string) { concatString += item })
}
func TestMaxBinaryHeapMap(t *testing.T) {
	heap := maxbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	globalCounter := 0
	maxbinaryheap.Map(heap, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestMaxBinaryHeapMapUpdateHeap(t *testing.T) {
	heap := maxbinaryheap.New[int](comparator.DefaultIntegerComparator)
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		heap.Add(item)
	}

	maxbinaryheap.Map(heap, func(item int) int {
		newItem := (9 - item) * (9 - item)
		return newItem
	})
	expectedMaxItem := 64
	maxItem, _ := heap.PeekMax()
	if maxItem != expectedMaxItem {
		t.Errorf("expected maximum item %v does not match found maximum item %v", expectedMaxItem, expectedMaxItem)
	}
}

func TestMaxBinaryHeapFold(t *testing.T) {
	heap := maxbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	itemsContained := maxbinaryheap.Fold(heap, true, func(item string, accumulator bool) bool {
		if !accumulator {
			return accumulator
		}
		return slices.Contains(items, item)
	})

	if !itemsContained {
		t.Errorf("expected all items to be contained in ground truth array")
	}
}

func TestMaxBinaryIterator(t *testing.T) {
	heap := maxbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	concatString := ""
	for item := range heap.Iterator() {
		concatString += item
	}
}
