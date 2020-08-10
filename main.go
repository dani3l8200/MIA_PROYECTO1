package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			l := newLexer(bytes.NewBufferString(eqn), os.Stdout, "file.name")
			yyParse(l)
		} else {
			break
		}
	}
	yyDebug = 0
	yyErrorVerbose = true

}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
