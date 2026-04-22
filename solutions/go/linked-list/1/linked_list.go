package linkedlist

import "errors"

// Node represents a node in the doubly linked list.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

// Next returns a pointer to the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// List represents a doubly linked list.
type List struct {
	first *Node
	last  *Node
}

// NewList creates a new linked list, preserving the order of the given values.
func NewList(args ...any) *List {
	l := &List{}
	for _, v := range args {
		l.Push(v)
	}
	return l
}

// First returns a pointer to the first node (head).
func (l *List) First() *Node {
	return l.first
}

// Last returns a pointer to the last node (tail).
func (l *List) Last() *Node {
	return l.last
}

// Push inserts a value at the back of the list.
func (l *List) Push(v any) {
	node := &Node{Value: v, prev: l.last}
	if l.last != nil {
		l.last.next = node
	} else {
		// List was empty — new node is also the head.
		l.first = node
	}
	l.last = node
}

// Pop removes a value from the back of the list.
func (l *List) Pop() (any, error) {
	if l.last == nil {
		return nil, errors.New("pop from empty list")
	}
	node := l.last
	l.last = node.prev
	if l.last != nil {
		l.last.next = nil
	} else {
		// List is now empty.
		l.first = nil
	}
	node.prev = nil
	return node.Value, nil
}

// Unshift inserts a value at the front of the list.
func (l *List) Unshift(v any) {
	node := &Node{Value: v, next: l.first}
	if l.first != nil {
		l.first.prev = node
	} else {
		// List was empty — new node is also the tail.
		l.last = node
	}
	l.first = node
}

// Shift removes a value from the front of the list.
func (l *List) Shift() (any, error) {
	if l.first == nil {
		return nil, errors.New("shift from empty list")
	}
	node := l.first
	l.first = node.next
	if l.first != nil {
		l.first.prev = nil
	} else {
		// List is now empty.
		l.last = nil
	}
	node.next = nil
	return node.Value, nil
}

// Reverse reverses the linked list in place.
func (l *List) Reverse() {
	curr := l.first
	for curr != nil {
		// Swap next and prev for this node.
		curr.next, curr.prev = curr.prev, curr.next
		// Move to the original "next", which is now stored in prev.
		curr = curr.prev
	}
	// Swap head and tail.
	l.first, l.last = l.last, l.first
}

func (l *List) Count() int {
	n := 0
	for curr := l.first; curr != nil; curr = curr.next {
		n++
	}
	return n
}

func (l *List) Delete(v any) error {
	for curr := l.first; curr != nil; curr = curr.next {
		if curr.Value == v {
			// Rewire previous (or update head).
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				l.first = curr.next
			}

			// Rewire next (or update tail).
			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				l.last = curr.prev
			}

			// Detach.
			curr.next = nil
			curr.prev = nil
			return nil
		}
	}
	return errors.New("value not found")
}