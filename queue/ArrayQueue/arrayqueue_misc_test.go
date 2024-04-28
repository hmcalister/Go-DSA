package arrayqueue_test

import (
	"testing"

	arrayqueue "github.com/hmcalister/Go-DSA/queue/ArrayQueue"
)

func TestArrayQueueInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		arrayqueue.New[int]()
	})
	t.Run("linked list float", func(t *testing.T) {
		arrayqueue.New[float64]()
	})
	t.Run("linked list string", func(t *testing.T) {
		arrayqueue.New[string]()
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		arrayqueue.New[S]()
	})
}

func TestCheckPeekOfEmptyArrayQueue(t *testing.T) {
	queue := arrayqueue.New[int]()

	_, err := queue.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty queue", err)
	}
}
