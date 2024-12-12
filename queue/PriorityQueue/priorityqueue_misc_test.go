package priorityqueue_test

import (
	"slices"
	"testing"

	priorityqueue "github.com/hmcalister/Go-DSA/queue/PriorityQueue"
	comparator "github.com/hmcalister/Go-DSA/utils/Comparator"
)

func TestPriorityQueueInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		priorityqueue.New[int](comparator.DefaultIntegerComparator)
	})
	t.Run("linked list float", func(t *testing.T) {
		priorityqueue.New[float64](comparator.DefaultFloat64Comparator)
	})
	t.Run("linked list string", func(t *testing.T) {
		priorityqueue.New[string](comparator.DefaultStringComparator)
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			f float64
			_ string
		}
		priorityqueue.New[S](func(a, b S) int {
			if a.f < b.f {
				return -1
			}

			if a.f > b.f {
				return +1
			}

			return 0
		})
	})
}

func TestCheckPeekOfEmptyPriorityQueue(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	_, err := queue.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty queue", err)
	}
}

func TestFindFromEmptyPriorityQueue(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	_, err := queue.Find(func(item int) bool { return item == 1 })
	if err == nil {
		t.Errorf("found nil error after finding from empty queue")
	}
}

func TestFindAllFromEmptyPriorityQueue(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	items := queue.FindAll(func(item int) bool { return item == 1 })
	if len(items) != 0 {
		t.Errorf("found a non-zero number of items from an empty queue")
	}
}

func TestPriorityQueueItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}

	retrievedItems := queue.Items()
	for _, item := range items {
		if !slices.Contains(retrievedItems, item) {
			t.Errorf("retrieved items %v does not contain expected item %v", retrievedItems, item)
		}
	}
}
