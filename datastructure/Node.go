package datastructure

import (
	"MIA-PROYECTO1/structs_lwh"
)

//Node struc
type Node struct {
	value structs_lwh.MountDisk
	next  *Node
}

// Getters

// Value method to return the current node value
func (n *Node) Value() structs_lwh.MountDisk {
	return n.value
}

// Next method to return the next node if exists
func (n *Node) Next() *Node {
	return n.next
}

// Setters

// SetNext method to set the next node
func (n *Node) SetNext(next *Node) {
	n.next = next
}
