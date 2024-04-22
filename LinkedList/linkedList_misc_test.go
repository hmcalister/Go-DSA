package linkedlist_test

import (
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

// define a helper that ensures function f panics
func shouldPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not find one")
		}
	}()
	f()
}

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
