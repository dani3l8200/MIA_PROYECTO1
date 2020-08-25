// Package lwh Contiene los analizadores usados para la lectura inicial
package lwh

import (
	"bytes"
	"fmt"
	"io"
)

// Node es una estructura que va a contener los tokens, donde le token contiene tipo y su valore correspondiente
type Node struct {
	Value     string
	TypeToken string
	Children  []Node
	Size      int
}

// METODO TOSTRING PARA COMPROBAR AL IMPRIMIR TODOS LOS TOKENS
func (n Node) String() string {
	buf := new(bytes.Buffer)
	n.print(buf, " ")
	return buf.String()
}

func (n Node) print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v%v", indent, n.Value, n.TypeToken)
	for _, nn := range n.Children {
		nn.print(out, indent+"  ")
	}
}

//********************************************************************************************

// NodeF es una clase donde se almacenaran los atributos tipo y valor del token
func NodeF(TypeToken string, Value string) Node {

	return Node{Value: Value, TypeToken: TypeToken}
}

// Length de vuelve el size de nuestra lista que contiene los tokens
func (n Node) Length() int {
	return n.Size
}

//********************************************************************************************
//******************************** METODO PARA ALMACENAR LOS TOKENS UNA LISTA SE PODRIA DECIR xd ***************************
func (n Node) Append(nn ...Node) Node {
	n.Children = append(n.Children, nn...)
	n.Size = n.Size + 1
	return n
}

//***********************************************************************************************
