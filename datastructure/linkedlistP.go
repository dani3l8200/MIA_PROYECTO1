package datastructure

import (
	"MIA-PROYECTO1/structs_lwh"
)

type LinkedListP struct {
	Head   *NodeP
	Tail   *NodeP
	length int
}

// Getters

// Length method to return the list length
func (l *LinkedListP) Length() int {
	return l.length
}

// Linked list structure methods

// Insert new node at the end of the linked list
func (l *LinkedListP) Insert(val structs_lwh.Pointer) {
	n := &NodeP{value: val}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.SetNext(n)
	}
	l.Tail = n
	l.length = l.length + 1
}

// InsertAt method adds a new value at the given position
func (l *LinkedListP) InsertAt(pos int, val structs_lwh.Pointer) {
	n := &NodeP{value: val}
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.length < pos:
		l.Insert(val)
	case pos == 1:
		n.SetNext(l.Head)
		l.Head = n
	default:
		node := l.Head
		// Position - 2 since we want the element replacing the given position
		for i := 1; i < (pos - 1); i++ {
			node = node.Next()
		}
		n.SetNext(node.Next())
		node.SetNext(n)
	}

	l.length = l.length + 1
}

// Get value in the given position
func (l *LinkedListP) Get(pos int) int64 {
	if pos > l.length {
		return -1
	}

	node := l.Head
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		node = node.Next()
	}

	return node.Value().PPointer
}

// Delete value at the given position
func (l *LinkedListP) Delete() {
	l.Head = nil
	l.Tail = nil
	l.length = 0
}
