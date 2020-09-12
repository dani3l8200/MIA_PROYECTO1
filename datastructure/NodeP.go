package datastructure

import (
	"MIA-PROYECTO1/structs_lwh"
)

//Node struc
type NodeP struct {
	value structs_lwh.Pointer
	next  *NodeP
}

// Getters

// Value method to return the current node value
func (n *NodeP) Value() structs_lwh.Pointer {
	return n.value
}

// Next method to return the next node if exists
func (n *NodeP) Next() *NodeP {
	return n.next
}

// Setters

// SetNext method to set the next node
func (n *NodeP) SetNext(next *NodeP) {
	n.next = next
}
