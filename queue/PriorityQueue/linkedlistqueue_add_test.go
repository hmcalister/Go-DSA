package priorityqueue_test

import (
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	priorityqueue "github.com/hmcalister/Go-DSA/queue/PriorityQueue"
)

func TestPriorityQueueAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	for _, item := range items {
		queue.Add(item)
	}
}

