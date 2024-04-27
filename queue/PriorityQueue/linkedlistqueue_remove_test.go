package priorityqueue_test

import (
	"slices"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	priorityqueue "github.com/hmcalister/Go-DSA/queue/PriorityQueue"
)

func TestPriorityQueueRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}

	for range items {
		_, err := queue.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty queue", err)
		}
	}
}

func TestPriorityQueueRemoveFromEmptyQueue(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	_, err := queue.Remove()
	if err == nil {
		t.Errorf("did not encounter error (%v) when removing from empty queue", err)
	}
}

func TestPriorityQueueCheckPeekAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}

	for _, item := range items {
		peekItem, err := queue.Peek()
		if err != nil {
			t.Errorf("encountered error (%v) after peeking at non-empty queue", err)
		}

		expectedItem := item
		if peekItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", peekItem, expectedItem)
		}

		_, err = queue.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty queue", err)
		}
	}
}

