package LWH

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Execute() {
	fmt.Println("*************************************************************************************************")
	fmt.Println("                                       SISTEMA DE ARCHIVOS LWH                                   ")
	fmt.Println("*************************************************************************************************")
	fi := bufio.NewReader(os.Stdin)

	for {
		var eqn string
		var ok bool

		fmt.Printf("~$: ")
		if eqn, ok = readline(fi); ok {
			ExecuteComands(eqn)
		} else {
			break
		}
	}

}

func ExecuteComands(entra string) {
	l := newLexer(bytes.NewBufferString(entra), os.Stdout, "file.name")
	if yyParse(l) == 0 {
		SelectCommands(root)
	}
	yyDebug = 0
	yyErrorVerbose = true
}

func SelectCommands(command node) {
	if command.typeToken == "PAUSE" {
		Pause()
	} else if command.typeToken == "MKDISK" {
		fmt.Println("AQUI ENTRO XD MKDISK")
	} else if command.typeToken == "RMDISK" {
		fmt.Println("RMDISK")
	} else if command.typeToken == "FDISK" {
		fmt.Println("FDISK")
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	s = strings.Replace(s, `\n`, "", -1)
	if err != nil {
		return "", false
	}
	return s, true
}

func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
