package datastructure

import "MIA-PROYECTO1/structs_lwh"

type NodeG struct {
	value structs_lwh.Group
	next  *NodeG
}

// Getters

// Value method to return the current node value
func (n *NodeG) Value() structs_lwh.Group {
	return n.value
}

// Next method to return the next node if exists
func (n *NodeG) Next() *NodeG {
	return n.next
}

// Setters

// SetNext method to set the next node
func (n *NodeG) SetNext(next *NodeG) {
	n.next = next
}
