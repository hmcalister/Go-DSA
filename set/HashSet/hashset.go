package hashset

import "iter"

// An implementation of a set using maps as the underlying data structure.
type HashSet[T comparable] struct {
	setData map[T]interface{}
}

// Create a new HashSet.
func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		setData: make(map[T]interface{}),
	}
}

// Return the size of the set, the number of items contained.
func (set *HashSet[T]) Size() int {
	return len(set.setData)
}

// Add an item to the set. Returns true if the item was *not* already present.
func (set *HashSet[T]) Add(item T) bool {
	ok := set.Contains(item)
	if !ok {
		set.setData[item] = struct{}{}
	}
	return !ok
}

// Checks if an item is already present in the set.
func (set *HashSet[T]) Contains(item T) bool {
	_, ok := set.setData[item]
	return ok
}

// Remove an item from the set. Returns an error if the item is not contained in the set.
func (set *HashSet[T]) Remove(item T) error {
	ok := set.Contains(item)
	if !ok {
		return ErrorItemNotContained
	}
	delete(set.setData, item)
	return nil
}

// Get all items from the hashset. This method allocates an array of length equal to the number of items.
// The items are not guaranteed to be in the order they were inserted into the hashset.
func (set *HashSet[T]) Items() []T {
	items := make([]T, set.Size())
	itemIndex := 0
	for item := range set.setData {
		items[itemIndex] = item
		itemIndex += 1
	}
	return items
}

// Iterate over the items of the hash set and apply a function to each item.
//
// Idiomatic Go should likely use Iterator() rather than functional methods.
//
// BEWARE: Iteration over a hashset does not guarantee a specific order ---
// you may find elements in any order, not the order they were inserted!
// Ensure your function accounts for this.
//
// To accumulate values over items, use Fold.
func Apply[T comparable](set *HashSet[T], f func(item T)) {
	for item := range set.setData {
		f(item)
	}
}

// Iterate over set items and apply the function f.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// Idiomatic Go should likely use Iterator() rather than functional methods.
//
// BEWARE: Iteration over a hashset does not guarantee a specific order ---
// you may find elements in any order, not the order they were inserted!
// Ensure your function accounts for this. This is especially important for
// a fold!
//
// This function is not a method on HashSet to allow for generic accumulators.
func Fold[T comparable, G any](set *HashSet[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	accumulator := initialAccumulator
	for item := range set.setData {
		accumulator = f(item, accumulator)
	}

	return accumulator
}

// Iterate over the items of the hashset. Note the iteration order may not be the insertion order.
// This method is not concurrency safe. For concurrent applications, consider using a mutex, or pull the data out using Items().
func (set *HashSet[T]) Iterator() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range set.setData {
			if !yield(item) {
				return
			}
		}
	}
}
