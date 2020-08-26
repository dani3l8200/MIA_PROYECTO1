package analyzers

import (
	"MIA-PROYECTO1/lwh"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Execute() {
	eqn := "Fdisk -sizE->135 -path->/home/dani3l8200/Escritorio/MisDiscos/Juan6.dsk -name->Particion3 -fit->BF -type->E"
	/*fmt.Println("************************************************************************************************")
	fmt.Println("                                       SISTEMA DE ARCHIVOS LWH                                 ")
	fmt.Println("************************************************************************************************")*/
	//fi := bufio.NewReader(os.Stdin)

	/*for {
	var eqn string
	var ok bool

	fmt.Printf("~$: ")
	if eqn, ok = readline(fi); ok {*/
	ExecuteComands(eqn)
	/*	} else {
			break
		}
	}*/

}

func ExecuteComands(entra string) {
	l := newLexer(bytes.NewBufferString(entra), os.Stdout, "file.name")
	if yyParse(l) == 0 {
		SelectCommands(Root)
	}
	yyDebug = 0
	yyErrorVerbose = true
}

// SelectCommands utilizaro para ver de que comando se trata
func SelectCommands(command lwh.Node) {
	if command.TypeToken == "PAUSE" {
		Pause()
	} else if command.TypeToken == "MKDISK" {
		if lwh.CheckMKdisk(command) {
			lwh.MakeMK(command)
		}
	} else if command.TypeToken == "RMDISK" {
		if lwh.CheckRMdisk(command) {
			lwh.DeleteDisk(command.Children[0].Value)
		}
	} else if command.TypeToken == "FDISK" {
		if lwh.ChekcFDisk(command) {
			lwh.MakeFdisk(command)
		}
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
