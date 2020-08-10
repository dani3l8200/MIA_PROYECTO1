package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))
	yyDebug = 0
	yyErrorVerbose = true
	for {
		var eqn string
		var ok bool

		fmt.Printf("Ingrese el comando: ")
		if eqn, ok = readline(fi); ok {
			l := analyzers.newLexer(bytes.NewBufferString(eqn), os.Stdout, "file.name")
			yyParse(l)
		} else {
			break
		}
	}

}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
