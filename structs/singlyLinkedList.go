package structs

// SinglyLinkedList represents a singly linked list data structure.
type SinglyLinkedList struct {
	fst *listElem
}

type listElem struct {
	next *listElem
}

// NewSinglyLinkedList creates an empty SinglyLinkedList.
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		fst: nil,
	}
}
