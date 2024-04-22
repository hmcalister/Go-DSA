package linkedlist

type LinkedList[T any] struct {
	// Head of the list, the first Node
	//
	// nil only when the length is zero
	head *linkedListNode[T]

	// Tail of the list, the last Node
	//
	// nil only when the length is zero
	tail *LinkedListNode[T]

	// Length of the list, the total number of Nodes
	length int
}

type LinkedListNode[T any] struct {
	Item T
	next *LinkedListNode[T]
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Add a new item to the end of the list
func (list *LinkedList[T]) Add(item T) {
	newNode := &LinkedListNode[T]{
		Item: item,
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

	newNode := &LinkedListNode[T]{
		Item: item,
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

// Remove the item from the end of the list
//
// Returns the item removed, or an error if the list is empty.
func (list *LinkedList[T]) Remove() (T, error) {
	if list.length == 0 {
		// Apparently idiomatic "zero-value" of a generic T is *new(T)... feels odd.
		// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type
		return *new(T), &EmptyListError{}
	}

}
