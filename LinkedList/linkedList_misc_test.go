package linkedlist_test

import (
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

func TestLinkedListInit(t *testing.T) {
	t.Run("linked list int", func(t *testing.T) {
		linkedlist.New[int]()
	})
	t.Run("linked list float", func(t *testing.T) {
		linkedlist.New[float64]()
	})
	t.Run("linked list string", func(t *testing.T) {
		linkedlist.New[string]()
	})
	t.Run("linked list struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		linkedlist.New[S]()
	})
}
