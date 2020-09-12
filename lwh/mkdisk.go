package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
				log.Fatal(err)
			}
			size = k
			fmt.Println(size)
			fmt.Printf("%+v\n", unsafe.Sizeof(k))
			fmt.Printf("%+v\n", unsafe.Sizeof(size))
		} else if v.TypeToken == "PATH" {
			path = v.Value
			fmt.Printf("%+v\n", path)
		} else if v.TypeToken == "NAME" {
			name = v.Value
			fmt.Printf("%+v\n", name)
		} else if v.TypeToken == "UNIT" {
			if strings.EqualFold(v.Value, "k") {
				unit = 'k'
				fmt.Printf("%+v\n", unit)
			} else if strings.EqualFold(v.Value, "m") {
				unit = 'm'
				fmt.Printf("%+v\n", unit)
			}
		}
		for _, j := range v.Children {
			if j.TypeToken == "SIZE" {
				k, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				size = k
				fmt.Println(j.Value)
				fmt.Printf("%+v\n", unsafe.Sizeof(size))
			} else if j.TypeToken == "PATH" {
				path = j.Value
				fmt.Printf("%+v\n", path)
			} else if j.TypeToken == "NAME" {
				name = j.Value
				fmt.Printf("%+v\n", name)
			} else if j.TypeToken == "UNIT" {
				if strings.EqualFold(j.Value, "k") {
					unit = 'K'
					//size = size * 1024
					fmt.Printf("%+v\n", size)
				} else if strings.EqualFold(j.Value, "m") {
					unit = 'M'
					//size = size * 1024 * 1024
					fmt.Printf("%+v\n", unit)
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

	var directory string = ""
	var newPath string = ""
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

	indexPath := strings.LastIndex(path, ".")
	if indexPath > -1 {
		aux1 := path[:indexPath]
		aux1 += "_copy.dsk"
		newPath += aux1
		fmt.Println(newPath)
	} else {
		fmt.Println("Route not copy :(")
	}
	indexSlash := strings.LastIndex(path, "/")
	if indexSlash > -1 {
		aux2 := path[:indexSlash]
		aux2 += "/"
		directory += aux2
		fmt.Println(directory)
	} else {
		fmt.Println("No se pudo obtener la ruta del directorio")
	}

	pathDisk += directory
	nameDisk += name
	// creando el directorio
	makeDirectory(directory)
	// Creando el archivo binario
	pathaux := directory + name
	writeFile(pathaux, mbr)
	// Creando una copia del archivo binario
	writeFile(newPath, mbr)
	//	read(path)
	/*f1, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()*/

}

/*func read(path string) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var m structs_lwh.MBR
	var size int = int(unsafe.Sizeof(m))
	size2 := binary.Size(mbr)

	fmt.Println(size)
	fmt.Println(size2)

	data := readNextBytes(f, size)

	buffer := bytes.NewBuffer(data)

	fmt.Println("AQUI ESTA", data)
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failled", err)
	}
	fmt.Printf("size: %d\nDiskSignature: %d\nTime: %+v\nCaracter:%+v\n ", m.MbrSize, m.MbrDiskSignature, m.MbrTime, m.Partition)

}*/

func writeFile(path string, m structs_lwh.MBR) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	return bytes
}

func writeNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
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
