package linkedlist_test

import (
	"slices"
	"testing"

	linkedlist "github.com/hmcalister/Go-DSA/LinkedList"
)

// Generic function to test adding items to list of different data types
func genericTestAdd[T any](t *testing.T, list *linkedlist.LinkedList[T], items []T) {
	for index, item := range items {
		list.Add(item)

		expectedLength := index + 1
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected length %v", list.Length(), expectedLength)
		}
	}
}

// Test adding items of different data types to a list
func TestGenericAdd(t *testing.T) {
	t.Run("testAdd int", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[int](), []int{1, 2, 3})
	})

	t.Run("testAdd float", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[float64](), []float64{1.0, 2.0, 3.0})
	})

	t.Run("testAdd string", func(t *testing.T) {
		genericTestAdd(t, linkedlist.New[string](), []string{"a", "b", "c"})
	})
}

// Test adding items to a list.
//
// Ensures list length is correct.
func TestListLengthOnAdd(t *testing.T) {
	list := linkedlist.New[int]()
	items := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for index, item := range items {
		list.Add(item)

		expectedLength := index + 1
		if list.Length() != expectedLength {
			t.Errorf("list length (%v) does not match expected list length (%v)", list.Length(), expectedLength)
		}
	}
}

// Tests adding items to a list at an index.
func TestAddAtIndex(t *testing.T) {
	items := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	newItem := 0

	// Define a small helper method that creates a new list,
	// inserts the items from the items array, then tries to insert an item
	// at the specified index.
	//
	// Calls t.Errorf if the insert fails or if the list length does not match the expected
	addAtIndexHelper := func(t *testing.T, addIndex int) {
		list := linkedlist.New[int]()
		for _, item := range items {
			list.Add(item)
		}

		err := list.AddAtIndex(newItem, addIndex)
		if err != nil {
			t.Errorf("error when adding item to list: %v", err)
		}

		expectedLength := len(items) + 1
		if list.Length() != expectedLength {
			t.Errorf("list length %v does not match expected list length %v", list.Length(), expectedLength)
		}
	}

	t.Run("add at head index", func(t *testing.T) {
		addAtIndexHelper(t, 0)
	})

	t.Run("add at non-head first-half index", func(t *testing.T) {
		addAtIndexHelper(t, 2)
	})

	t.Run("add at non-head second-half index", func(t *testing.T) {
		addAtIndexHelper(t, len(items)-3)
	})

	t.Run("add at tail index", func(t *testing.T) {
		addAtIndexHelper(t, len(items)-1)
	})
}

// Tests that the list pointers are correct after multiple additions
func TestPointerCorrectnessAfterAdd(t *testing.T) {
	list := linkedlist.New[string]()
	items := []string{"a", "b", "c", "d", "e", "f", "g"}
	for index, item := range items {
		list.Add(item)
		currentItems := slices.Clone(items)[:index+1]

		expectedConcatString := ""
		for _, item := range currentItems {
			expectedConcatString += item
		}
		concatStr := ""
		linkedlist.ForwardApply(list, func(item string) {
			concatStr += item
		})
		if expectedConcatString != concatStr {
			t.Errorf("forward concatenated string (%v) does not match expected concatenated string (%v)", concatStr, expectedConcatString)
		}

		// Reverse the items to test back concat
		slices.Reverse(currentItems)

		expectedConcatString = ""
		for _, item := range currentItems {
			expectedConcatString += item
		}
		concatStr = ""
		linkedlist.ReverseApply(list, func(item string) {
			concatStr += item
		})
		if expectedConcatString != concatStr {
			t.Errorf("backwards concatenated string (%v) does not match expected concatenated string (%v)", concatStr, expectedConcatString)
		}
	}
}

// Tests that the list pointers are correct after addition at an index
func TestPointerCorrectnessAfterAddAtIndex(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e", "f", "g"}
	newItem := "z"

	addAtIndexHelper := func(t *testing.T, targetIndex int) {
		list := linkedlist.New[string]()
		for _, item := range items {
			list.Add(item)
		}
		list.AddAtIndex(newItem, targetIndex)

		// slices.Insert will not affect the original items array
		effectiveItems := slices.Insert(items, targetIndex, newItem)
		expectedConcatString := ""
		for _, item := range effectiveItems {
			expectedConcatString += item
		}
		concatStr := ""
		linkedlist.ForwardApply(list, func(item string) {
			concatStr += item
		})
		if expectedConcatString != concatStr {
			t.Errorf("forward concatenated string (%v) does not match expected concatenated string (%v) for deletion at index %v", concatStr, expectedConcatString, targetIndex)
		}

		// Reverse the items to test back concat
		slices.Reverse(effectiveItems)
		expectedConcatString = ""
		for _, item := range effectiveItems {
			expectedConcatString += item
		}
		concatStr = ""
		linkedlist.ReverseApply(list, func(item string) {
			concatStr += item
		})
		if expectedConcatString != concatStr {
			t.Errorf("backwards concatenated string (%v) does not match expected concatenated string (%v) for deletion at index %v", concatStr, expectedConcatString, targetIndex)
		}
	}

	t.Run("add at head index", func(t *testing.T) {
		addAtIndexHelper(t, 0)
	})

	t.Run("add at non-head first-half index", func(t *testing.T) {
		addAtIndexHelper(t, 2)
	})

	t.Run("add at non-tail second-half index", func(t *testing.T) {
		addAtIndexHelper(t, len(items)-3)
	})

	t.Run("add at tail index", func(t *testing.T) {
		addAtIndexHelper(t, len(items)-1)
	})

}
