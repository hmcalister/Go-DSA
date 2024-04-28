package linkedlistqueue_test

import (
	"testing"

	linkedlistqueue "github.com/hmcalister/Go-DSA/queue/LinkedListQueue"
)

func TestLinkedListQueueRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

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

func TestLinkedListQueueRemoveFromEmptyQueue(t *testing.T) {
	queue := linkedlistqueue.New[int]()

	_, err := queue.Remove()
	if err == nil {
		t.Errorf("did not encounter error (%v) when removing from empty queue", err)
	}
}

func TestLinkedListQueueCheckPeekAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

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

func TestLinkedListQueueCheckRemovedItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

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

func TestLinkedListQueueCheckSizeAfterRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

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

func TestLinkedListQueueCheckFindAfterRemove(t *testing.T) {
	queue := linkedlistqueue.New[int]()

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
