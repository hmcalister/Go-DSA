package linkedlist

type LinkedList[T any] struct {
	// Head of the list, the first Node
	//
	// nil only when the length is zero
	head *linkedListNode[T]

	// Tail of the list, the last Node
	//
	// nil only when the length is zero
	tail *linkedListNode[T]

	// Length of the list, the total number of Nodes
	length int
}

type linkedListNode[T any] struct {
	item T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}

// Create a new linked list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Iterate over the list in the forward direction and apply a function to each item
func (list *LinkedList[T]) IterateApplyFunction(f func(item T)) {
	currentNode := list.head
	for currentNode != nil {
		f(currentNode.item)
		currentNode = currentNode.next
	}
}

// Iterate over the list in the reverse direction and apply a function to each item
func (list *LinkedList[T]) ReverseIterateApplyFunction(f func(item T)) {
	currentNode := list.tail
	for currentNode != nil {
		f(currentNode.item)
		currentNode = currentNode.prev
	}
}

// Get the length of this linked list
func (list *LinkedList[T]) Length() int {
	return list.length
}

// Get the item at the specified index
//
// Returns an error if the index is out of bounds
func (list *LinkedList[T]) ItemAtIndex(index int) (T, error) {
	if list.length <= index {
		return *new(T), &IndexOutOfBoundsError{
			targetIndex: index,
			listLength:  list.length,
		}
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

// Add a new item to the end of the list
func (list *LinkedList[T]) Add(item T) {
	newNode := &linkedListNode[T]{
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
// # Returns a IndexOutOfBoundsError if the specified index is out of bounds
//
// Example:
//
// ```
// list := linkedlist.New[string]()
// list.Add("hello")				// list = ["hello"]
// list.Add("world")				// list = ["hello", "world"]
// list.AddAtIndex("(linked!)", 1)	// list = ["hello", "(linked!)", "world"]
// ````
func (list *LinkedList[T]) AddAtIndex(item T, index int) error {
	// Note here we allow list.length==index, as we *can* insert at the end of the list
	if list.length < index {
		return &IndexOutOfBoundsError{
			targetIndex: index,
			listLength:  list.length,
		}
	}

	newNode := &linkedListNode[T]{
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
	var beforeSpliceNode *linkedListNode[T]
	var afterSpliceNode *linkedListNode[T]
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

// Remove and return the item from the end of the list
//
// Returns the item removed, or an error if the list is empty.
func (list *LinkedList[T]) Remove() (T, error) {
	if list.length == 0 {
		// Apparently idiomatic "zero-value" of a generic T is *new(T)... feels odd.
		// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type
		return *new(T), &EmptyListError{}
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
// Returns an error if the list is empty, or is the target index is out of range.
//
// Example:
// ```
// list := linkedlist.New[string]()
// list.Add("hello")					// list = ["hello"]
// list.Add("(linked!)")				// list = ["hello", "(linked!)"]
// list.Add("world")					// list = ["hello", "(linked!)", "world"]
//
// item, err := list.RemoveAtIndex(1)	// list = ["hello", "world"]
// fmt.Printf("%v", item)				// (linked!)
// ```
func (list *LinkedList[T]) RemoveAtIndex(index int) (T, error) {
	if list.length == 0 {
		// Apparently idiomatic "zero-value" of a generic T is *new(T)... feels odd.
		// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type
		return *new(T), &EmptyListError{}
	}

	// Note here we do not allow RemoveAtIndex(list.Length()) as this is "out of bounds"
	// and unlike inserting it does not make sense to define it here.
	if list.length <= index {
		return *new(T), &IndexOutOfBoundsError{
			targetIndex: index,
			listLength:  list.length,
		}
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
	var beforeSpliceNode *linkedListNode[T]
	var afterSpliceNode *linkedListNode[T]
	var removedNode *linkedListNode[T]
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
