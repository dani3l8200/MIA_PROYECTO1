package datastructure

import (
	"MIA-PROYECTO1/structs_lwh"
	"fmt"
	"strconv"
)

//LinkedList structura donde se almacenaran los ids, nombres, paths para el mount
type LinkedList struct {
	head       *Node
	tail       *Node
	length     int
	charActual byte
}

// Getters

// Length method to return the list length
func (l *LinkedList) Length() int {
	return l.length
}

// Linked list structure methods

// Insert new node at the end of the linked list
func (l *LinkedList) Insert(val structs_lwh.MountDisk) {
	n := &Node{value: val}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.SetNext(n)
	}
	l.tail = n
	l.length = l.length + 1
}

// InsertAt method adds a new value at the given position
func (l *LinkedList) InsertAt(pos int, val structs_lwh.MountDisk) {
	n := &Node{value: val}
	// If the given position is lower than the list length
	// the element will be inserted at the end of the list
	switch {
	case l.length < pos:
		l.Insert(val)
	case pos == 1:
		n.SetNext(l.head)
		l.head = n
	default:
		node := l.head
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
func (l *LinkedList) Get(pos int) interface{} {
	fmt.Println(l.length)
	if pos > l.length {
		return nil
	}

	node := l.head
	// Position - 1 since we want the value in the given position
	for i := 0; i < pos-1; i++ {
		node = node.Next()
	}

	return node.Value()
}

// Delete value at the given position
func (l *LinkedList) Delete(pos int) bool {
	if pos > l.length {
		return false
	}

	node := l.head
	if pos == 1 {
		l.head = node.Next()
	} else {
		for i := 1; i < pos-1; i++ {
			node = node.Next()
		}
		node.SetNext(node.Next().Next())
	}
	l.length = l.length - 1
	return true
}

//DeleteMount Desmonta la particion, seria para el comano unmount
func (l *LinkedList) DeleteMount(ID string) {
	if l.head != nil {
		aux := l.head
		node := l.head
		for ; node != nil; node = node.Next() {
			if ID == node.Value().ID {
				break
			}
			aux = node
		}
		if node == nil {
			fmt.Println("Particion no Montada")
			return
		}

		if aux == node {
			l.head = node.Next()
		} else {
			aux.next = node.Next()
		}
		l.length = l.length - 1

	}
	fmt.Println("PARTICION ELIMINADA CON EXITO")
	l.Print()
}

// Each method to apply to each element in the linked list a
// function who receives an interface and don't return any value
func (l *LinkedList) Each(f func(val structs_lwh.MountDisk)) {
	for n := l.head; n != nil; n = n.Next() {
		f(n.Value())
	}
}

//GetMountedPart obtiene un objeto Mount, que ya fue insertado con anterioridad en la lista
func (l *LinkedList) GetMountedPart(ID string) (structs_lwh.MountDisk, bool) {
	var disk structs_lwh.MountDisk
	if l.head != nil {
		for node := l.head; node != nil; node = node.Next() {
			if ID == node.Value().ID {
				disk = node.Value()
				return disk, true
			}
		}
	}
	return disk, false
}

//MountedPart Verifica si ya se monto alguna particion del disco
func (l *LinkedList) MountedPart(path string, name string) bool {
	if l.head != nil {
		for node := l.head; node != nil; node = node.Next() {
			if path == node.Value().Path && name == node.Value().Name {
				fmt.Println("SI ENTRO ACA xd")
				return true
			}
		}
	}
	return false
}

//SetLetter genera la letra para el ID del mount, como lo indica el enunciado
func (l *LinkedList) SetLetter(path string) byte {
	var letter byte = 97
	if l.head != nil {
		for node := l.head; node != nil; node = node.Next() {
			if path == node.Value().Path {
				letter = node.Value().ID[2]
				break
			} else if letter <= node.Value().ID[2] {
				letter++
			}
		}
	}
	return letter
}

//SetNumber genera el numero para el id del mount, como lo indica el enunciado
func (l *LinkedList) SetNumber(path string) int {
	var i = 1
	if l.head != nil {
		for node := l.head; node != nil; node = node.Next() {
			if path == node.Value().Path {
				var number [10]byte
				copy(number[:], "")
				for x := 3; x < len(node.Value().ID); x++ {

					aux := node.Value().ID[x]
					fmt.Println(aux)
					number[x-3] = aux
				}
				var auxiliar = ""
				for _, k := range number {
					if k != 0 {
						auxiliar += string(k)
					}
				}
				i, _ = strconv.Atoi(auxiliar)
				i++
			}
		}
	}
	return i
}

// Print LinkedList
func (l *LinkedList) Print() {
	if l.head != nil {
		fmt.Println("Lista de Particiones Montadas: ")
		for node := l.head; node != nil; node = node.Next() {
			fmt.Println("ID:", node.Value().ID, "Nombre:", node.Value().Name, "Directorio:", node.Value().Path)
		}
	}
}
