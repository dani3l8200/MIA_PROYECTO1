package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var mbr structs_lwh.MBR
var pathDisk string = ""
var nameDisk string = ""
var startDisk int64 = 0
var fitDisk byte

// MakeMK funcion para el mkdisk
func MakeMK(root Node) {
	var size int64 = 0
	var path string = ""
	var name string = ""
	var unit byte = 'K'

	for _, v := range root.Children {
		if v.TypeToken == "SIZE" {
			k, err := strconv.ParseInt(v.Value, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			size = k
		} else if v.TypeToken == "PATH" {
			path = v.Value
		} else if v.TypeToken == "NAME" {
			name = v.Value
		} else if v.TypeToken == "UNIT" {
			if strings.EqualFold(v.Value, "k") {
				unit = 'k'
			} else if strings.EqualFold(v.Value, "m") {
				unit = 'm'
			}
		}
		for _, j := range v.Children {
			if j.TypeToken == "SIZE" {
				k, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					fmt.Println(err)
				}
				size = k
			} else if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "NAME" {
				name = j.Value
			} else if j.TypeToken == "UNIT" {
				if strings.EqualFold(j.Value, "k") {
					unit = 'K'
				} else if strings.EqualFold(j.Value, "m") {
					unit = 'M'
				}
			}
		}
	}

	diskSignature := GenerateSignature()
	CreateDisk(path, size, diskSignature, name, unit)

}

// CreateDisk Funcion para crear el disco y tambien inicializar el mbr
func CreateDisk(path string, size int64, diskSignature int64, name string, unit byte) {
	auxSize := verifySize(unit, size)
	var auxName string = ""
	times := time.Now()
	fixFormat := times.Format("01-02-2006 15:04:00")

	copy(mbr.MbrTime[:], fixFormat)

	mbr.MbrDiskSignature = diskSignature
	mbr.MbrSize = auxSize
	for i := 0; i < 4; i++ {
		mbr.Partition[i].PartStatus = '0'
		mbr.Partition[i].PartStart = -1
		mbr.Partition[i].PartSize = 0
		copy(mbr.Partition[i].PartName[:], "")
	}
	if strings.Contains(path, "\"") {
		path, _ = SetDirectory(path)
	}

	if strings.Contains(name, "\"") {
		auxName, _ = SetDirectory(name)
	}

	if auxName != "" {
		name = auxName
	}

	if strings.Contains(path, ".d") {
		path = RemakeDirectory(path) + "/"
	}

	nameDisk += name
	pathaux := path + name
	// creando el directorio
	makeDirectory(path)
	// Creando el archivo binario

	writeFile(pathaux, mbr)

	fmt.Println("SE CREO EL DISCO EXICTOSAMENTE :)")

	Pause()
	//	read(path)
	/*f1, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()*/

}

func writeFile(path string, m structs_lwh.MBR) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var pos [2]byte
	copy(pos[:], "\\0")

	var binario3 bytes.Buffer

	binary.Write(&binario3, binary.BigEndian, &pos)

	writeNextBytes(f, binario3.Bytes())

	f.Seek(m.MbrSize, 0)

	var binario2 bytes.Buffer

	binary.Write(&binario2, binary.BigEndian, &pos)

	writeNextBytes(f, binario2.Bytes())

	f.Seek(0, 0)

	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &m)

	writeNextBytes(f, binario.Bytes())
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}

// GenerateRandomSignature Genera un signature random
func GenerateRandomSignature(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// GenerateSignature devuelve un signature de tipo random
func GenerateSignature() int64 {
	rand.Seed(time.Now().UnixNano())
	randNumber := GenerateRandomSignature(1, 1000)
	return randNumber
}

func makeDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

func verifySize(unit byte, size int64) int64 {
	if unit == 'K' {
		size = size * 1020
	} else if unit == 'M' {
		size = size * 1020 * 1020
	} else if unit == 'B' {
		size = size * 1
	}
	return size
}

//Pause ...
func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

//RemakeDirectory ...
func RemakeDirectory(path string) string {
	index := strings.LastIndex(path, "/")
	if strings.Contains(path, ".d") {
		if index > -1 {
			path = path[:index]
		}
	}
	return path
}
