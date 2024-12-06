package hashset

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

// Remove an item from the set. Returns an error if the item is not contained in the set
func (set *HashSet[T]) Remove(item T) error {
	ok := set.Contains(item)
	if !ok {
		return ErrorItemNotContained
	}
	delete(set.setData, item)
	return nil
}
