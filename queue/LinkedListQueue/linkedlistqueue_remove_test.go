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

