package arraystack_test

import (
	"testing"

	arraystack "github.com/hmcalister/Go-DSA/stack/ArrayStack"
)

func TestArrayStackInit(t *testing.T) {
	t.Run("array stack int", func(t *testing.T) {
		arraystack.New[int]()
	})
	t.Run("array stack float", func(t *testing.T) {
		arraystack.New[float64]()
	})
	t.Run("array stack string", func(t *testing.T) {
		arraystack.New[string]()
	})
	t.Run("array stack struct", func(t *testing.T) {
		type S struct {
			_ int
			_ float64
			_ string
		}
		arraystack.New[S]()
	})
}

