package lwh

import (
	"MIA-PROYECTO1/datastructure"
	"MIA-PROYECTO1/structs_lwh"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func MakeFile(Root Node) {
	var path string = ""
	var cont string = ""
	var id string = ""
	var size int64 = -1
	var p bool = false
	for _, i := range Root.Children {
		if i.TypeToken == "PATH" {
			path = i.Value
		} else if i.TypeToken == "CONT" {
			cont = i.Value
		} else if i.TypeToken == "SIZE" {
			k, err := strconv.ParseInt(i.Value, 10, 64)
			if err != nil {
				log.Fatalln("ERROR DE CONVERSION")
			}
			size = k
		} else if i.TypeToken == "P" {
			p = true
		} else if i.TypeToken == "ID" {
			id = i.Value
		}
		for _, j := range i.Children {
			if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "CONT" {
				cont = j.Value
			} else if j.TypeToken == "SIZE" {
				k, err := strconv.ParseInt(j.Value, 10, 64)
				if err != nil {
					log.Fatalln("ERROR DE CONVERSION")
				}
				size = k
			} else if j.TypeToken == "P" {
				p = true
			} else if j.TypeToken == "ID" {
				id = j.Value
			}
		}
	}

	disk, err := lista.GetMountedPart(id)
	if err {
		getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

		if getData.GetSize != 0 && getData.GetStart != 0 {
			RecorrerAVD(disk.GetPath(), getData.GetStart, size, path, cont, p)
		}

	} else if err == false {
		fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
	}
}

func RecorrerAVD(pathDisk string, start int64, size int64, path string, cont string, p bool) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	auxPath := strings.Split(path, "/")

	auxPath = remove(auxPath, "")

	f.Seek(start, io.SeekStart)

	sb := readFileSB(f, err)

	var pos int = 0

	var apIndirecto int64 = 0

	avd := ReadAVD(apIndirecto, f, err, sb.SbApTreeDirectory)

	if p {
		for pos < len(auxPath) {
			if pos < len(auxPath)-1 {
				if existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1 {

				} else if existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) == -1 {

				}
			}
		}
	}

}

func SetDirectory(path string) (string, bool) {

	index := strings.LastIndex(path, "/")
	if strings.Contains(path, "\"") {
		path = strings.ReplaceAll(path, "\"", "")
		if index > -1 {
			path = path[:index]
			path += "/"
			return path, true
		}
	} else if !strings.Contains(path, "\"") {
		if index > -1 {
			path = path[:index]
			path += "/"
			return path, true
		}
	}
	return "", false
}

func createPointer(name [20]byte, ppointer int64) structs_lwh.Pointer {
	var pointer structs_lwh.Pointer

	pointer.Name = name
	pointer.PPointer = ppointer

	return pointer
}

func createListPointer(f *os.File, err error, SbApTreeDirectory int64, avd structs_lwh.AVD) datastructure.LinkedListP {

	var listaP datastructure.LinkedListP

	for i := 0; i < 6; i++ {
		if avd.AvdApArraySubdirectories[i] != -1 {
			aux := ReadAVD(avd.AvdApArraySubdirectories[i], f, err, SbApTreeDirectory)

			listaP.Insert(createPointer(aux.AvdNameDirectory, avd.AvdApArraySubdirectories[i]))
		}

	}

	if avd.AvdApTreeVirtualDirectory != -1 {

		aux := ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, SbApTreeDirectory)

		auxList := createListPointer(f, err, SbApTreeDirectory, aux)

		if auxList.Head != nil {
			for node := auxList.Head; node != nil; node = node.Next() {

				point := node.Value()

				listaP.Insert(point)
			}
		}

	}

	return listaP
}

func existDirectoryAVD(f *os.File, err error, startAvd int64, avd structs_lwh.AVD, name string) int64 {
	auxName := converNameSBToByte(name)
	for i := 0; i < 6; i++ {
		if avd.AvdApArraySubdirectories[i] != -1 {
			checkAvd := ReadAVD(avd.AvdApArraySubdirectories[i], f, err, startAvd)

			if checkAvd.AvdNameDirectory == auxName {
				return avd.AvdApArraySubdirectories[i]
			}
		}
	}

	if avd.AvdApTreeVirtualDirectory != -1 {
		auxAVD := ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, startAvd)
		return existDirectoryAVD(f, err, startAvd, auxAVD, name)
	}

	return -1
}

func converNameSBToByte(name string) [20]byte {
	var auxName [20]byte
	for i, j := range []byte(name) {
		auxName[i] = byte(j)
	}
	return auxName
}

func convertBByteToString(name [25]byte) string {
	var DbData string = ""
	for _, name := range name {
		if name != 0 {
			DbData += string(name)
		}
	}
	return DbData
}
