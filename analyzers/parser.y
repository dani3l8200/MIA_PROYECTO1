%{
package analyzers

import (
  "fmt"
  "bytes"
  "io"
)

type node struct {
  name string
  children []node
}

func (n node) String() string {
  buf := new(bytes.Buffer)
  n.print(buf, " ")
  return buf.String()
}

func (n node) print(out io.Writer, indent string) {
  fmt.Fprintf(out, "\n%v%v", indent, n.name)
  for _, nn := range n.children { nn.print(out, indent + "  ") }
}

func Node(name string) node { return node{name: name} }
func (n node) append(nn...node) node { n.children = append(n.children, nn...); return n }

%}

%union{
    node node
    token string
}

%token  INT BOOL FLOAT INT CHAR STRING IDENT NUM  

%type <token> BOOL FLOAT INT CHAR STRING IDENT NUM 
%type <node> Input Type

%%

Input: /* empty */ { }
     | Input Type {fmt.Println($2)}
     ;
Type: INT ';' {$$ = Node("int")}
    |  FLOAT ';' {$$ = Node("float")}
    |  CHAR ';' {$$ = Node("char")}
    |  STRING ';' {$$ = Node("string")}
    |  BOOL ';' {$$ = Node("bool")}
    |  IDENT '}' {$$ = Node()}




