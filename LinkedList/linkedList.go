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
}

// Create a new linked list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
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

	// If we are requesting the final item, just access it directly
	if index == list.length-1 {
		return list.tail.item, nil
	}

	// Otherwise we perform list traversal
	currentNode := list.head
	for _ = range index {
		currentNode = currentNode.next
	}
	return currentNode.item, nil
}

// Add a new item to the end of the list
func (list *LinkedList[T]) Add(item T) {
	newNode := &linkedListNode[T]{
		item: item,
		next: nil,
	}

	// If length is zero, tail is nil as we have an empty list
	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
		return
	}

	// Otherwise, we add the node to the end of the list
	//
	// This is defined as tail.next
	list.tail.next = newNode

	// We also update the tail to be the newly inserted node
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

	// If we are inserting at the tail, we can simplify the call by just calling list.Add
	if index == list.length {
		list.Add(item)
		return nil
	}

	// We now know we have to handle the insert logic ourselves, so let's make the Node

	newNode := &linkedListNode[T]{
		item: item,
		next: nil,
	}

	// If we are inserting at the head of the list (index=0)
	// We have a special case, as we splice into
	// list.head rather than node.next
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
		return nil
	}

	// If we are inserting anywhere except the head or tail, we walk along the list
	//
	// Starting from the head, walk (index-1) nodes along the list.
	// This gives us the node *before* the splice position, i.e. the node to update next of
	currentNode := list.head
	for _ = range index - 1 {
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = newNode

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

	// We must traverse the list to get the second-to-last node
	// This could be improved with a doubly-linked-list,
	// i.e. keep a pointer to next *and* previous

	currentNode := list.head
	for _ = range list.length - 2 {
		currentNode = currentNode.next
	}

	// We can now set the list tail to the new, previously penultimate node
	removedNode := list.tail
	list.tail = currentNode
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

	if list.length <= index {
		return *new(T), &IndexOutOfBoundsError{
			targetIndex: index,
			listLength:  list.length,
		}
	}

	// If we are trying to remove from the end of the list, just call the regular Remove method
	// This also handles the case of removing the head of a list of length 1
	if index == list.length-1 {
		return list.Remove()
	}

	// If we are trying to remove the head of the list,
	// splice the head node out with list.head pointer
	if index == 0 {
		prevHead := list.head
		list.head = list.head.next
		list.length -= 1
		return prevHead.item, nil
	}

	// Otherwise...
	// We do the same traversal as in list.Remove but we stop some way along
	// and splice the node out

	currentNode := list.head
	for _ = range index - 1 {
		currentNode = currentNode.next
	}

	// currentNode is now one before the node in question,
	// and we know it is not the last node (by above) if index==list.length-1
	removedNode := currentNode.next
	currentNode.next = removedNode.next
	list.length -= 1
	return removedNode.item, nil
}
