package hashset_test

import (
	"math/rand/v2"
	"testing"

	hashset "github.com/hmcalister/Go-DSA/set/HashSet"
)

// ----------------------------------------------------------------------------
// Initialization Tests

func TestHashSetIntInit(t *testing.T) {
	hashset.New[int]()
}

func TestHashSetFloatInit(t *testing.T) {
	hashset.New[float64]()
}

func TestHashSetStringInit(t *testing.T) {
	hashset.New[string]()
}

func TestHashSetStructInit(t *testing.T) {
	type S struct {
		_ int
		_ float64
		_ string
	}
	hashset.New[S]()
}

// ----------------------------------------------------------------------------
// Misc Tests

func TestHashSetSize(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for i, item := range items {
		set.Add(item)
		if set.Size() != i+1 {
			t.Errorf("hash set has size %v, expected size %v", set.Size(), i+1)
		}
	}
}

func TestHashSetContains(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		set.Add(item)
	}

	for _, item := range items {
		if !set.Contains(item) {
			t.Errorf("hash set claims to not contain expected item %v", item)
		}
	}

	if set.Contains(0) {
		t.Errorf("hash set claims to contain unexpected item")
	}
}

// ----------------------------------------------------------------------------
// Add Tests

func TestHashSetAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		set.Add(item)
	}
}

func TestHashSetContainsDuringAdd(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		if set.Contains(item) {
			t.Errorf("hash set claims to contain unexpected item %v before addition", item)
		}
		set.Add(item)
		if !set.Contains(item) {
			t.Errorf("hash set claims to not contain expected item %v after addition", item)
		}
	}
}

func TestHashSetAddRandomOrder(t *testing.T) {
	set := hashset.New[int]()

	numItems := 100
	items := make([]int, numItems)
	for i := range numItems {
		items[i] = i
	}
	rand.Shuffle(numItems, func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	for _, item := range items {
		set.Add(item)
	}
}

// ----------------------------------------------------------------------------
// Remove Tests

func TestHashSetRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		set.Add(item)
	}

	for _, item := range items {
		err := set.Remove(item)
		if err != nil {
			t.Error("failed to remove item from hash set")
		}
	}
}

func TestHashSetRemoveUnexpectedItem(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		set.Add(item)
	}

	err := set.Remove(0)
	if err != hashset.ErrorItemNotContained {
		t.Errorf("removal of unexpected item did not result in expected error, got error %v", err)
	}
}

func TestHashSetContainsDuringRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := hashset.New[int]()

	for _, item := range items {
		set.Add(item)
	}

	for _, item := range items {
		if !set.Contains(item) {
			t.Errorf("hash set claims to not contain expected item %v before removal", item)
		}
		err := set.Remove(item)
		if err != nil {
			t.Error("failed to remove item from hash set")
		}
		if set.Contains(item) {
			t.Errorf("hash set claims to contain unexpected item %v after removal", item)
		}
	}
}
