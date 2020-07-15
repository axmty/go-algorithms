package singlelist

// List represents a singly linked list data structure.
type List struct {
	front *Elem
}

// Elem represents a list element in a singly linked list.
type Elem struct {
	Next  *Elem
	Value interface{}
}

// New creates an empty SinglyLinkedList.
func New() *List {
	return &List{
		front: nil,
	}
}
