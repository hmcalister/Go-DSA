package linkedlist

import dsa_error "github.com/hmcalister/Go-DSA/utils/DSA_Error"

// An implementation of a doubly linked list.
type LinkedList[T any] struct {
	// Head of the list, the first Node.
	//
	// nil only when the length is zero.
	head *LinkedListNode[T]

	// Tail of the list, the last Node.
	//
	// nil only when the length is zero.
	tail *LinkedListNode[T]

	// Length of the list, the total number of Nodes.
	length int
}

// Create a new linked list.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Get the length of this linked list.
func (list *LinkedList[T]) Length() int {
	return list.length
}

// ----------------------------------------------------------------------------
// Get and Find methods

// Find and return the first item in the list satisfying a predicate function.
// If no item satisfies the predicate, a dsa_error.ErrorItemNotFound is returned instead.
//
// The list is walked forward during this search.
func (list *LinkedList[T]) Find(predicate func(item T) bool) (T, error) {
	currentNode := list.head
	for currentNode != nil {
		if predicate(currentNode.item) {
			return currentNode.item, nil
		}
		currentNode = currentNode.next
	}

	return *new(T), dsa_error.ErrorItemNotFound
}

// Find and return the last item in the list satisfying a predicate function.
// If no item satisfies the predicate, a dsa_error.ErrorItemNotFound is returned instead.
//
// The list is walked backward during this search.
func (list *LinkedList[T]) ReverseFind(predicate func(item T) bool) (T, error) {
	currentNode := list.tail
	for currentNode != nil {
		if predicate(currentNode.item) {
			return currentNode.item, nil
		}
		currentNode = currentNode.prev
	}

	return *new(T), dsa_error.ErrorItemNotFound
}

// Find ALL of the items in the list satisfying a predicate.
// If no item satisfies the predicate, the list will be empty.
// The list is walked forward during this search.
func (list *LinkedList[T]) FindAll(predicate func(item T) bool) []T {
	satisfyingItems := make([]T, 0)
	currentNode := list.head
	for currentNode != nil {
		if predicate(currentNode.item) {
			satisfyingItems = append(satisfyingItems, currentNode.item)
		}
		currentNode = currentNode.next
	}

	return satisfyingItems
}

// Find ALL of the items in the list satisfying a predicate.
// If no item satisfies the predicate, the list will be empty.
// The list is walked backwards during this search.
func (list *LinkedList[T]) ReverseFindAll(predicate func(item T) bool) []T {
	satisfyingItems := make([]T, 0)
	currentNode := list.tail
	for currentNode != nil {
		if predicate(currentNode.item) {
			satisfyingItems = append(satisfyingItems, currentNode.item)
		}
		currentNode = currentNode.prev
	}

	return satisfyingItems
}

// Get the item at the specified index.
//
// Returns a dsa_error.ErrorIndexOutOfBounds if the index is out of bounds.
func (list *LinkedList[T]) ItemAtIndex(index int) (T, error) {
	if list.length <= index {
		return *new(T), dsa_error.ErrorIndexOutOfBounds
	}

	// If the target index is after the halfway point
	// we can traverse backwards to find the node
	if index > list.length/2 {
		currentNode := list.tail
		for range list.length - index - 1 {
			currentNode = currentNode.prev
		}
		return currentNode.item, nil
	} else {
		currentNode := list.head
		for range index {
			currentNode = currentNode.next
		}
		return currentNode.item, nil
	}
}

// Get all items from the list. This method allocates an array of length equal to the number of items.
func (list *LinkedList[T]) Items() []T {
	items := make([]T, list.Length())
	itemIndex := 0
	currentNode := list.head
	for currentNode != nil {
		items[itemIndex] = currentNode.item
		currentNode = currentNode.next
		itemIndex += 1
	}
	return items
}

// ----------------------------------------------------------------------------
// Apply, Map, and Fold methods
//
// Methods to apply a function across ALL nodes in a list.

// Iterate over the list in the forward direction and apply a function to each item.
//
// It is expected that ForwardApply does *not* update the list items.
// To modify the list items, use ForwardMap.
// To accumulate values over the list, use ForwardFold.
func ForwardApply[T any](list *LinkedList[T], f func(item T)) {
	currentNode := list.head
	for currentNode != nil {
		f(currentNode.item)
		currentNode = currentNode.next
	}
}

// Iterate over the list in the forward direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ForwardMap can update the node items by returning the update value.
// If you do not need to modify the list items, use ForwardApply.
// To accumulate values over the list, use ForwardFold.
func ForwardMap[T any](list *LinkedList[T], f func(item T) T) {
	currentNode := list.head
	for currentNode != nil {
		currentNode.item = f(currentNode.item)
		currentNode = currentNode.next
	}
}

// Iterate over the list and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on LinkedList to allow for generic accumulators.
func ForwardFold[T any, G any](list *LinkedList[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentNode := list.head
	acc := initialAccumulator
	for currentNode != nil {
		acc = f(currentNode.item, acc)
		currentNode = currentNode.next
	}

	return acc
}

// Iterate over the list in the reverse direction and apply a function to each item.
//
// It is expected that ReverseApply does *not* update the list items.
// To modify the list items, use ReverseMap.
// To accumulate values over the list, use ReverseFold.
func ReverseApply[T any](list *LinkedList[T], f func(item T)) {
	currentNode := list.tail
	for currentNode != nil {
		f(currentNode.item)
		currentNode = currentNode.prev
	}
}

// Iterate over the list in the reverse direction and apply a function to each item
// The result of this function is then assigned to the node at each step.
//
// ReverseMap can update the node items by returning the update value.
// If you do not need to modify the list items, use ReverseApply.
// To accumulate values over the list, use ReverseFold.
func ReverseMap[T any](list *LinkedList[T], f func(item T) T) {
	currentNode := list.tail
	for currentNode != nil {
		currentNode.item = f(currentNode.item)
		currentNode = currentNode.prev
	}
}

// Iterate over the list and apply the function f to it.
// The function f also takes the current value of the accumulator.
// The results of f become the new value of the accumulator at each step.
//
// This function returns the final accumulator.
//
// This function is not a method on LinkedList to allow for generic accumulators.
func ReverseFold[T any, G any](list *LinkedList[T], initialAccumulator G, f func(item T, accumulator G) G) G {
	currentNode := list.tail
	acc := initialAccumulator
	for currentNode != nil {
		acc = f(currentNode.item, acc)
		currentNode = currentNode.prev
	}

	return acc
}

// ----------------------------------------------------------------------------
// Add methods

// Add a new item to the end of the list.
func (list *LinkedList[T]) Add(item T) {
	newNode := &LinkedListNode[T]{
		item: item,
	}

	// If length is zero, tail is nil as we have an empty list
	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
		list.length += 1
		return
	}

	// Otherwise, we add the node to the end of the list
	//
	// This is defined as tail.next
	// we also ensure the previous is set correctly
	list.tail.next = newNode
	newNode.prev = list.tail

	// Update the tail to be the newly inserted node
	list.tail = newNode

	// Finally, update the length
	list.length += 1
}

// Add a new item to the list in the specified position.
//
// Returns a dsa_error.ErrorIndexOutOfBounds if the specified index is out of bounds.
//
// Example:
//
// ```
//
// list := linkedlist.New[string]()
//
// list.Add("hello")				// list = ["hello"]
//
// list.Add("world")				// list = ["hello", "world"]
//
// list.AddAtIndex("(linked!)", 1)	// list = ["hello", "(linked!)", "world"]
//
// ````
func (list *LinkedList[T]) AddAtIndex(item T, index int) error {
	// Note here we allow list.length==index, as we *can* insert at the end of the list
	if list.length < index {
		return dsa_error.ErrorIndexOutOfBounds
	}

	newNode := &LinkedListNode[T]{
		item: item,
	}

	// If we are inserting at the head of the list (index=0)
	// we have a special case, as we splice into
	// list.head rather than node.next
	if index == 0 {
		// Set the list pointers correctly
		newNode.next = list.head
		list.head.prev = newNode

		list.head = newNode
		list.length += 1
		return nil
	}

	// If we are inserting at the tail of the list,
	// we have another special case. Here can can just call
	// list.Add()
	if index == list.length {
		list.Add(item)
		return nil
	}

	// Otherwise we have to do some traversal.
	// Let's find the before and after splice nodes
	var beforeSpliceNode *LinkedListNode[T]
	var afterSpliceNode *LinkedListNode[T]
	if index > list.length/2 {
		currentNode := list.tail
		for range list.length - index - 1 {
			currentNode = currentNode.prev
		}
		beforeSpliceNode = currentNode.prev
		afterSpliceNode = currentNode
	} else {
		currentNode := list.head
		for range index - 1 {
			currentNode = currentNode.next
		}
		beforeSpliceNode = currentNode
		afterSpliceNode = currentNode.next
	}

	// And set the pointers correctly
	beforeSpliceNode.next = newNode
	newNode.prev = beforeSpliceNode
	afterSpliceNode.prev = newNode
	newNode.next = afterSpliceNode

	list.length += 1

	return nil
}

// ----------------------------------------------------------------------------
// Remove methods

// Remove and return the item from the end of the list.
//
// Returns the item removed, or a dsa_error.ErrorDataStructureEmpty if the list is empty.
func (list *LinkedList[T]) Remove() (T, error) {
	if list.length == 0 {
		// Apparently idiomatic "zero-value" of a generic T is *new(T)... feels odd.
		// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	// If we have only one element, we must remove both head *and* tail
	if list.length == 1 {
		removedNode := list.head
		list.head = nil
		list.tail = nil
		list.length -= 1
		return removedNode.item, nil
	}

	removedNode := list.tail
	newTail := list.tail.prev

	list.tail = newTail
	removedNode.next = nil
	removedNode.prev = nil
	newTail.next = nil

	list.length -= 1
	return removedNode.item, nil
}

// Remove and return the item from a particular index.
//
// Returns a dsa_error.ErrorDataStructureEmpty if the list is empty,
// or a dsa_error.ErrorIndexOutOfBounds if the target index is out of range.
//
// Example:
//
// ```
//
// list := linkedlist.New[string]()
//
// list.Add("hello")					// list = ["hello"]
//
// list.Add("(linked!)")				// list = ["hello", "(linked!)"]
//
// list.Add("world")					// list = ["hello", "(linked!)", "world"]
//
// item, err := list.RemoveAtIndex(1)	// list = ["hello", "world"]
//
// fmt.Printf("%v", item)				// (linked!)
//
// ```
func (list *LinkedList[T]) RemoveAtIndex(index int) (T, error) {
	if list.length == 0 {
		// Apparently idiomatic "zero-value" of a generic T is *new(T)... feels odd.
		// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type
		return *new(T), dsa_error.ErrorDataStructureEmpty
	}

	// Note here we do not allow RemoveAtIndex(list.Length()) as this is "out of bounds"
	// and unlike inserting it does not make sense to define it here.
	if list.length <= index {
		return *new(T), dsa_error.ErrorIndexOutOfBounds
	}

	// If we are removing at the tail of the list,
	// we have another special case. Here can can just call
	// list.Remove()
	if index == list.length-1 {
		return list.Remove()
	}

	// If we are removing at the head of the list (index=0)
	// we have a special case, as we splice into
	// list.head rather than node.next
	if index == 0 {
		removedNode := list.head
		// Set the list pointers correctly
		list.head = removedNode.next
		list.head.prev = nil

		list.length -= 1
		return removedNode.item, nil
	}

	// Otherwise we have to do some traversal.
	// Let's find the before and after splice nodes
	var beforeSpliceNode *LinkedListNode[T]
	var afterSpliceNode *LinkedListNode[T]
	var removedNode *LinkedListNode[T]
	if index > list.length/2 {
		currentNode := list.tail
		for range list.length - index - 2 {
			currentNode = currentNode.prev
		}
		afterSpliceNode = currentNode
		removedNode = currentNode.prev
		beforeSpliceNode = removedNode.prev
	} else {
		currentNode := list.head
		for range index - 1 {
			currentNode = currentNode.next
		}
		beforeSpliceNode = currentNode
		removedNode = beforeSpliceNode.next
		afterSpliceNode = removedNode.next
	}

	// And set the pointers correctly
	beforeSpliceNode.next = afterSpliceNode
	afterSpliceNode.prev = beforeSpliceNode
	removedNode.next = nil
	removedNode.prev = nil

	list.length -= 1
	return removedNode.item, nil
}
