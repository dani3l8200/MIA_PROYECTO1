package LWH

import (
	"bytes"
	"fmt"
	"io"
)

type node struct {
	value     string
	typeToken string
	children  []node
}

func (n node) String() string {
	buf := new(bytes.Buffer)
	n.print(buf, " ")
	return buf.String()
}

func (n node) print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v", indent, n.value)
	for _, nn := range n.children {
		nn.print(out, indent+"  ")
	}
}

func Node(typeToken string, value string) node { return node{value: value, typeToken: typeToken} }
func (n node) append(nn ...node) node          { n.children = append(n.children, nn...); return n }
