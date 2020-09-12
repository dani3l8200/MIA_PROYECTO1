package datastructure

import (
	"MIA-PROYECTO1/structs_lwh"
)

type NodeU struct {
	value structs_lwh.User
	next  *NodeU
}

// Getters

// Value method to return the current node value
func (n *NodeU) Value() structs_lwh.User {
	return n.value
}

// Next method to return the next node if exists
func (n *NodeU) Next() *NodeU {
	return n.next
}

// Setters

// SetNext method to set the next node
func (n *NodeU) SetNext(next *NodeU) {
	n.next = next
}
