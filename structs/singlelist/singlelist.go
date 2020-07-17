package singlelist

// Elem is an element of a linked list.
type Elem struct {
	Value interface{}
	next  *Elem
	list  *List
}

// Next returns the next list element or nil.
func (e *Elem) Next() *Elem { return e.next }

// List represents a singly linked list data structure.
type List struct {
	front *Elem
	len   int
}

// New returns an initialized list.
func New() *List {
	return &List{
		front: nil,
		len:   0,
	}

}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.front = nil
	l.len = 0
	return l
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Elem { return l.front }

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Elem {
	if l.len == 0 {
		return nil
	}
	e := l.front
	for e.next != nil {
		e = e.next
	}
	return e
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified. The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Elem) *Elem {
	if mark.list != l {
		return nil
	}
	e := &Elem{Value: v, next: mark.next, list: l}
	mark.next = e
	l.len++
	return e
}

// Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() int { return l.len }
