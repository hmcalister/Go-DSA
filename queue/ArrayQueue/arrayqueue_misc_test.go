package arrayqueue_test

import (
	"slices"
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

func TestFindFromEmptyArrayQueue(t *testing.T) {
	queue := arrayqueue.New[int]()

	_, err := queue.Find(func(item int) bool { return item == 1 })
	if err == nil {
		t.Errorf("found nil error after finding from empty queue")
	}
}

func TestFindAllFromEmptyArrayQueue(t *testing.T) {
	queue := arrayqueue.New[int]()

	items := queue.FindAll(func(item int) bool { return item == 1 })
	if len(items) != 0 {
		t.Errorf("found a non-zero number of items from an empty queue")
	}
}

func TestArrayQueueItems(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	queue := arrayqueue.New[int]()

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
