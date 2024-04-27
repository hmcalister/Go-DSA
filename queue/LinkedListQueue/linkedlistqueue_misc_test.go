package linkedlistqueue_test

import (
	"testing"

	linkedlistqueue "github.com/hmcalister/Go-DSA/queue/LinkedListQueue"
)

func TestLinkedListQueueInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		linkedlistqueue.New[int]()
	})
	t.Run("linked list float", func(t *testing.T) {
		linkedlistqueue.New[float64]()
	})
	t.Run("linked list string", func(t *testing.T) {
		linkedlistqueue.New[string]()
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		linkedlistqueue.New[S]()
	})
}

func TestCheckPeekOfEmptyLinkedListQueue(t *testing.T) {
	queue := linkedlistqueue.New[int]()

	_, err := queue.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty queue", err)
	}
}
