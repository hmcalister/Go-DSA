package minbinaryheap_test

import (
	"fmt"
	"slices"
	"testing"

	minbinaryheap "github.com/hmcalister/Go-DSA/heap/MinBinaryHeap"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

func TestMinBinaryHeapApply(t *testing.T) {
	heap := minbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	concatString := ""
	minbinaryheap.Apply(heap, func(item string) { concatString += item })
	expectedConcatString := ""
	for _, item := range items {
		expectedConcatString += item
	}

	if concatString != expectedConcatString {
		t.Errorf("result (%v) does not match expected result (%v)", concatString, expectedConcatString)
	}
}
func TestMinBinaryHeapMap(t *testing.T) {
	heap := minbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	globalCounter := 0
	minbinaryheap.Map(heap, func(item string) string {
		newItem := fmt.Sprintf("%v, %v", item, globalCounter)
		globalCounter += 1
		return newItem
	})
}

func TestMinBinaryHeapMapUpdateHeap(t *testing.T) {
	heap := minbinaryheap.New[int](comparator.DefaultIntegerComparator)
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, item := range items {
		heap.Add(item)
	}

	minbinaryheap.Map(heap, func(item int) int {
		newItem := 9 - item
		return newItem
	})
	expectedMinItem := 0
	minItem, _ := heap.PeekMin()
	if minItem != expectedMinItem {
		t.Errorf("expected minimum item %v does not match found minimum item %v", expectedMinItem, minItem)
	}
}

func TestMinBinaryHeapFold(t *testing.T) {
	heap := minbinaryheap.New[string](comparator.DefaultStringComparator)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, item := range items {
		heap.Add(item)
	}

	itemsContained := minbinaryheap.Fold(heap, true, func(item string, accumulator bool) bool {
		if !accumulator {
			return accumulator
		}
		return slices.Contains(items, item)
	})

	if !itemsContained {
		t.Errorf("expected all items to be contained in ground truth array")
	}
}
