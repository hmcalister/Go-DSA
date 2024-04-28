package arrayqueue_test

import (
	"testing"

	arrayqueue "github.com/hmcalister/Go-DSA/queue/ArrayQueue"
)

func TestArrayQueueAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := arrayqueue.New[int]()

	for _, item := range items {
		queue.Add(item)
	}
}

