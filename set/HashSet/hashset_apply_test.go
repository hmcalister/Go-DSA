package hashset_test

import (
	"testing"

	hashset "github.com/hmcalister/Go-DSA/set/HashSet"
)

func TestApply(t *testing.T) {
	set := hashset.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		set.Add(item)
	}

	sum := 0
	hashset.Apply(set, func(item int) { sum += item })
	expectedSum := 0
	for _, item := range items {
		expectedSum += item
	}

	if sum != expectedSum {
		t.Errorf("result (%v) does not match expected result (%v)", sum, expectedSum)
	}
}

func TestFold(t *testing.T) {
	set := hashset.New[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, item := range items {
		set.Add(item)
	}

	sum := hashset.Fold(set, 0, func(item int, accumulator int) int {
		return accumulator + item
	})

	expectedSum := 0
	for _, item := range items {
		expectedSum += item
	}

	if sum != expectedSum {
		t.Errorf("result (%v) does not match expected result (%v)", sum, expectedSum)
	}
}
