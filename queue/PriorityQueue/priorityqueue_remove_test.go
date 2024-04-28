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

func TestPriorityQueueCheckPeekAfterRemoveNonpriorityOrder(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	sortedItems := make([]int, len(items))
	copy(sortedItems, items)
	slices.Sort(sortedItems)

	for _, item := range items {
		queue.Add(item)
	}

	for _, item := range sortedItems {
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

func TestPriorityQueueCheckRemovedItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}

	for _, item := range items {
		removedItem, err := queue.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty queue", err)
		}

		expectedItem := item
		if removedItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", removedItem, expectedItem)
		}

	}
}

func TestPriorityQueueCheckRemovedItemNonpriorityOrder(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	sortedItems := make([]int, len(items))
	copy(sortedItems, items)
	slices.Sort(sortedItems)

	for _, item := range items {
		queue.Add(item)
	}

	for _, item := range sortedItems {
		removedItem, err := queue.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty queue", err)
		}

		expectedItem := item
		if removedItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", removedItem, expectedItem)
		}

	}
}

func TestPriorityQueueCheckSizeAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}

	for index := range items {
		queueSize := queue.Size()
		expectedSize := len(items) - index
		if queueSize != expectedSize {
			t.Errorf("found queue size (%v) does not match the expected size (%v)", queueSize, expectedSize)
		}

		_, err := queue.Remove()
		if err != nil {
			t.Errorf("encountered error (%v) when removing from non-empty queue", err)
		}
	}
}

func TestPriorityQueueCheckFindAfterRemove(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	targetItem := 1
	queue.Add(targetItem)
	item, err := queue.Find(func(item int) bool { return item == targetItem })
	if err != nil {
		t.Errorf("found error (%v) after finding from queue that should have item", err)
	}
	if item != targetItem {
		t.Errorf("found item (%v) does not match expected item (%v)", item, targetItem)
	}

	queue.Remove()
	_, err = queue.Find(func(item int) bool { return item == targetItem })
	if err == nil {
		t.Errorf("found nil error after finding from queue without item")
	}
}

func TestPriorityQueueCheckFindAllAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}
	for range len(items) - 1 {
		queue.Remove()
	}

	foundItems := queue.FindAll(func(item int) bool { return item%2 == 0 })
	if len(foundItems) != 0 {
		t.Errorf("found a non-zero number of items from a queue with expected zero number of matches")
	}
}
