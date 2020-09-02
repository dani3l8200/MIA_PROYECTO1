package analyzers

import (
	"MIA-PROYECTO1/lwh"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Execute realiza el iniciamiento del proyecto :D
func Execute() {
	//eqn := "mkdisk -size->10 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Disco5.dsk -unit->m"
	//eqn := "Fdisk -sizE->40 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion1 \n"
	//eqn := "Fdisk -sizE->50 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion2 -fit->BF"
	//eqn := "Fdisk -sizE->63 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion3  -fit->BF"
	//eqn := "Fdisk -sizE->883 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion4 -type->E"
	//eqn := "Fdisk -sizE->12 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion5 -type->L"
	//eqn := "Fdisk -sizE->25 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion6 -type->L"
	//eqn := "Fdisk -sizE->45 -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion7 -type->L"
	//eqn := "mount -path->/home/dani3l8200/Escritorio/MisDiscos/Disco5.dsk -name->Particion1"

	//eqn := "exec -path->/home/dani3l8200/Escritorio/MisDiscos/ht1.txt"

	fmt.Println("************************************************************************************************")
	fmt.Println("                                       SISTEMA DE ARCHIVOS LWH                                 ")
	fmt.Println("************************************************************************************************")

	var flag = false
	for !flag {
		fmt.Printf("~$: ")
		str := readLine()
		if strings.EqualFold(str, "exit") {
			flag = true
		} else {
			for strings.Contains(str, "\\*") {
				str = strings.ReplaceAll(str, "\\*", "")
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("> ")
				scanner.Scan()
				str += scanner.Text()
			}
			ExecuteComands(str)
		}

	}
	//ExecuteComands(eqn)

}

// ExecuteComands ejecuta todos los comandos definidos en el enunciado
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
	} else if command.TypeToken == "EXEC" {
		ReadMyFile(command.Children[0].Value)
	} else if command.TypeToken == "COMENTARIO" {
		fmt.Println(command.Value)
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
	} else if command.TypeToken == "MOUNT" {
		if lwh.CheckMount(command) {
			lwh.MountPartitions(command)
		}
	} else if command.TypeToken == "UNMOUNT" {
		if lwh.CheckUnmount(command) {
			lwh.UnmountPartitions(command)
		}
	}
	for _, i := range command.Children {
		if i.TypeToken == "PAUSE" {
			Pause()
		} else if i.TypeToken == "EXEC" {
			ReadMyFile(i.Value)
		} else if i.TypeToken == "COMENTARIO" {
			fmt.Println(i.Value)
		} else if i.TypeToken == "MKDISK" {
			if lwh.CheckMKdisk(i) {
				lwh.MakeMK(i)
			}
		} else if i.TypeToken == "RMDISK" {
			if lwh.CheckRMdisk(i) {
				lwh.DeleteDisk(i.Children[0].Value)
			}
		} else if i.TypeToken == "FDISK" {
			if lwh.ChekcFDisk(i) {
				lwh.MakeFdisk(i)
			}
		} else if i.TypeToken == "MOUNT" {
			if lwh.CheckMount(i) {
				lwh.MountPartitions(i)
			}
		} else if i.TypeToken == "UNMOUNT" {
			if lwh.CheckUnmount(i) {
				lwh.UnmountPartitions(i)
			}
		}
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	return strings.TrimSpace(str)
}

// ReadMyFile realiza la accion para el comano exec xd
func ReadMyFile(ruta string) {
	ruta = strings.ReplaceAll(ruta, "\"", "")
	archivo, err := os.Open(ruta)
	defer func() {
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err != nil {
		panic("~$: 'El fichero o directorio no existe'\n")
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		if linea := scanner.Text(); linea != "" {
			for strings.Contains(linea, "\\*") {
				fmt.Println(linea)
				linea = strings.ReplaceAll(linea, "\\*", "")
				scanner.Scan()
				linea += scanner.Text()
				fmt.Println(linea)
			}

			ExecuteComands(linea)
		}
	}
}

// Pause realiza el comando pausa, descrito en el enunciado
func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
