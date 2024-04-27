package linkedlistqueue_test

import (
	"testing"

	linkedlistqueue "github.com/hmcalister/Go-DSA/queue/LinkedListQueue"
)

func TestLinkedListQueueAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

	for _, item := range items {
		queue.Add(item)
	}
}

func TestLinkedListQueueCheckPeekAfterAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

	for _, item := range items {
		queue.Add(item)

		peekItem, err := queue.Peek()
		if err != nil {
			t.Errorf("encountered error (%v) after peeking at non-empty queue", err)
		}

		expectedItem := items[0]
		if peekItem != expectedItem {
			t.Errorf("found peek item (%v) does not match the expected item (%v)", peekItem, expectedItem)
		}
	}
}

func TestLinkedListQueueCheckSizeAfterAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := linkedlistqueue.New[int]()

	for index, item := range items {
		queue.Add(item)

		queueSize := queue.Size()
		expectedSize := index + 1
		if queueSize != expectedSize {
			t.Errorf("found queue size (%v) does not match the expected size (%v)", queueSize, expectedSize)
		}
	}
}
