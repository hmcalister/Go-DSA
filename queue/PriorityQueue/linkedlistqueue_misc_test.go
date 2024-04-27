package priorityqueue_test

import (
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	priorityqueue "github.com/hmcalister/Go-DSA/queue/PriorityQueue"
)

func TestPriorityQueueInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		priorityqueue.New[int](comparator.DefaultIntegerComparator)
	})
	t.Run("linked list float", func(t *testing.T) {
		priorityqueue.New[float64](comparator.DefaultFloat64Comparator)
	})
	t.Run("linked list string", func(t *testing.T) {
		priorityqueue.New[string](comparator.DefaultStringComparator)
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			f float64
			_ string
		}
		priorityqueue.New[S](func(a, b S) int {
			if a.f < b.f {
				return -1
			}

			if a.f > b.f {
				return +1
			}

			return 0
		})
	})
}

func TestCheckPeekOfEmptyPriorityQueue(t *testing.T) {
	queue := priorityqueue.New[int](comparator.DefaultIntegerComparator)

	_, err := queue.Peek()
	if err == nil {
		t.Errorf("did not encounter error (%v) when peeking at empty queue", err)
	}
}
