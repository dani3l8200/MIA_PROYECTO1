package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// MakeFdisk crear las particiones primarias, extendidas y logicas
func MakeFdisk(root Node) {
	var path string = ""
	var name string = ""
	var size int32 = 0
	var unit byte = 'K'
	var tipo byte = 'P'
	var fit byte = 'W'
	var del int8 = -1
	var add int16 = 0
	for _, i := range root.Children {
		if i.TypeToken == "SIZE" {
			k, err := strconv.Atoi(i.Value)
			if err != nil {
				log.Fatal(err)
			}
			size = int32(k)
			fmt.Println(size)
		} else if i.TypeToken == "UNIT" {
			if strings.EqualFold(i.Value, "k") {
				unit = 'K'
				fmt.Printf("%c\n", unit)
			} else if strings.EqualFold(i.Value, "m") {
				unit = 'M'
				fmt.Printf("%c\n", unit)
			} else if strings.EqualFold(i.Value, "b") {
				unit = 'B'
				fmt.Printf("%c\n", unit)
			}
		} else if i.TypeToken == "PATH" {
			path = i.Value
			fmt.Printf("%+v\n", path)
		} else if i.TypeToken == "TYPE" {
			if strings.EqualFold(i.Value, "p") {
				tipo = 'P'
				fmt.Printf("%c\n", tipo)
			} else if strings.EqualFold(i.Value, "e") {
				tipo = 'E'
				fmt.Printf("%c\n", tipo)
			} else if strings.EqualFold(i.Value, "l") {
				tipo = 'L'
				fmt.Printf("%c\n", tipo)
			}
		} else if i.TypeToken == "FIT" {
			if strings.EqualFold(i.Value, "BF") {
				fit = 'B'
				fmt.Printf("%c\n", fit)
			} else if strings.EqualFold(i.Value, "FF") {
				fit = 'F'
				fmt.Printf("%c\n", fit)
			} else if strings.EqualFold(i.Value, "WF") {
				unit = 'W'
				fmt.Printf("%c\n", fit)
			}
		} else if i.TypeToken == "DELETE" {
			if strings.EqualFold(i.Value, "fast") {
				del = 0
				fmt.Printf("%+v\n", del)
			} else if strings.EqualFold(i.Value, "full") {
				del = 1
				fmt.Printf("%+v\n", del)
			}
		} else if i.TypeToken == "NAME" {
			name = i.Value
			fmt.Printf("%+v\n", name)
		} else if i.TypeToken == "ADD" {
			k, err := strconv.Atoi(i.Value)
			if err != nil {
				log.Fatal(err)
			}
			add = int16(k)
			fmt.Println(add)
		}
		for _, j := range i.Children {
			if j.TypeToken == "SIZE" {
				k, err := strconv.Atoi(j.Value)
				if err != nil {
					log.Fatal(err)
				}
				size = int32(k)
				fmt.Println(size)
			} else if j.TypeToken == "UNIT" {
				if strings.EqualFold(j.Value, "k") {
					unit = 'K'
					fmt.Printf("%c\n", unit)
				} else if strings.EqualFold(j.Value, "m") {
					unit = 'M'
					fmt.Printf("%c\n", unit)
				} else if strings.EqualFold(j.Value, "b") {
					unit = 'B'
					fmt.Printf("%c\n", unit)
				}
			} else if j.TypeToken == "PATH" {
				path = j.Value
				fmt.Printf("%+v\n", path)
			} else if j.TypeToken == "TYPE" {
				if strings.EqualFold(j.Value, "p") {
					tipo = 'P'
					fmt.Printf("%c\n", tipo)
				} else if strings.EqualFold(j.Value, "e") {
					tipo = 'E'
					fmt.Printf("%c\n", tipo)
				} else if strings.EqualFold(j.Value, "l") {
					tipo = 'L'
					fmt.Printf("%c\n", tipo)
				}
			} else if j.TypeToken == "FIT" {
				if strings.EqualFold(j.Value, "BF") {
					fit = 'B'
					fmt.Printf("%c\n", fit)
				} else if strings.EqualFold(j.Value, "FF") {
					fit = 'F'
					fmt.Printf("%c\n", fit)
				} else if strings.EqualFold(j.Value, "WF") {
					unit = 'W'
					fmt.Printf("%c\n", fit)
				}
			} else if j.TypeToken == "DELETE" {
				if strings.EqualFold(j.Value, "fast") {
					del = 0
					fmt.Printf("%+v\n", del)
				} else if strings.EqualFold(j.Value, "full") {
					del = 1
					fmt.Printf("%+v\n", del)
				}
			} else if j.TypeToken == "NAME" {
				name = j.Value
				fmt.Printf("%+v\n", name)
			} else if j.TypeToken == "ADD" {
				k, err := strconv.Atoi(j.Value)
				if err != nil {
					log.Fatal(err)
				}
				add = int16(k)
				fmt.Println(add)
			}
		}
	}

	if _, err := os.Stat(path); err == nil {
		if size != 0 {
			if tipo == 'P' {
				ParticionPrimaria(path, name, tipo, fit, unit, size)
			} else if tipo == 'E' {
				ParticionExtendida(path, name, tipo, fit, unit, size)
			} else if tipo == 'L' {

			}
		} else if add != 0 {

		} else if del != -1 {

		}
	} else if os.IsNotExist(err) {
		panic(err)
	}

}

// ParticionPrimaria este metodo es usado para crear la particion primaria en el disco
func ParticionPrimaria(path string, name string, tipo byte, fit byte, unit byte, size int32) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	fmt.Printf("%+v\n", auxSize)
	m := readFileDisk(f, err)
	var a int32 = 0
	for i := 0; i < 4; i++ {
		if m.Partition[i].PartStatus != '1' {
			a += m.Partition[i].PartSize
		}
	}

	if m.MbrSize-a >= auxSize {
		var checkPartition bool = false
		fmt.Println(checkPartition)
		for i := 0; i < 4; i++ {
			if m.Partition[i].PartName == auxName {
				checkPartition = true
				break
			} else if m.Partition[i].PartType == 'E' {
				f.Seek(int64(m.MbrSize), 0)
				e := readFileEBR(f, err)
				fmt.Println(e.PartStartE)
				pos, _ := f.Seek(0, os.SEEK_CUR)
				fmt.Println(pos)
				for pos < int64(m.Partition[i].PartSize)+int64(m.Partition[i].PartStart) {
					if e.PartNameE == auxName {
						checkPartition = true
					}
					if e.PartNextE == -1 {
						break
					}
				}
			}

		}

		if !checkPartition {
			var index int = 0
			if fit == 'B' {
				index = BestFit(m, auxSize)
				fmt.Println(index)
			} else if fit == 'F' {
				index = FirstFit(m, auxSize)
			} else if fit == 'W' {
				index = WorstFit(m, auxSize)
			}

			if index != -1 {
				copy(m.Partition[index].PartName[:], name)
				m.Partition[index].PartFit = fit
				m.Partition[index].PartType = tipo
				m.Partition[index].PartSize = auxSize
				m.Partition[index].PartStatus = '0'

				if index == 0 {
					m.Partition[index].PartStart = int32(binary.Size(m))
				} else {
					m.Partition[index].PartStart = m.Partition[index-1].PartStart + m.Partition[index-1].PartSize
				}
				writeInPrimaryPartition(path, m, index)
				fmt.Println("SE CREO LA PARTICION PRIMARIA :D")
			} else {
				fmt.Println("YA SE HA CREADO 4 PARTICIONES YA NO SE PUEDEN CREAR MAS :(")
			}
		} else {
			fmt.Println("LA PARTICION YA HA SIDO CREADA CON ESTE NOMBRE ", name)
		}

	} else {
		fmt.Println("YA NO CUENTA CON ESPACIO EN EL DISCO, EL DISTO TIENE UN ESPACIO DE: ", string(m.MbrSize), " Y ESTA TRATANDO DE HACER UNA PARTICION DE ", string(size))
	}
}

// ParticionExtendida Crea unicamente un particion extendida, escribiendo en el struct del mbr en su respetiva particion y en el EBR
func ParticionExtendida(path string, name string, tipo byte, fit byte, unit byte, size int32) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	fmt.Printf("%+v\n", auxSize)
	m := readFileDisk(f, err)
	e := readFileEBR(f, err)
	var ebr structs_lwh.EBR
	var a int32 = 0

	for i := 0; i < 4; i++ {
		if m.Partition[i].PartType == 'E' {
			fmt.Println("SOLO UNA PARTICION EXTENDIDA ES PERMITIDA")
			return
		}
	}

	for i := 0; i < 4; i++ {
		if m.Partition[i].PartStatus != '1' {
			a += m.Partition[i].PartSize
		}
	}

	if m.MbrSize-a >= auxSize {
		var checkPartition bool = false
		for i := 0; i < 4; i++ {
			if m.Partition[i].PartName == auxName {
				checkPartition = true
				break
			} else if m.Partition[i].PartType == 'E' {
				f.Seek(int64(m.MbrSize), int(m.Partition[i].PartStart))
				pos, _ := f.Seek(0, os.SEEK_CUR)
				eSize := binary.Size(e)
				fmt.Println(pos)
				for eSize != 0 && pos < int64(m.Partition[i].PartSize)+int64(m.Partition[i].PartStart) {
					if e.PartNameE == auxName {
						checkPartition = true
					}
					if e.PartNextE == -1 {
						break
					}
				}
			}
		}
		if !checkPartition {
			var index int = -1
			if fit == 'B' {
				index = BestFit(m, auxSize)
			} else if fit == 'F' {
				index = FirstFit(m, auxSize)
			} else if fit == 'W' {
				index = WorstFit(m, auxSize)
			}

			if index != -1 {
				copy(m.Partition[index].PartName[:], name)
				m.Partition[index].PartFit = fit
				m.Partition[index].PartSize = auxSize
				m.Partition[index].PartType = tipo
				m.Partition[index].PartStatus = '0'

				if index == 0 {
					m.Partition[index].PartStart = int32(binary.Size(m))
				} else {
					m.Partition[index].PartStart = m.Partition[index-1].PartStart + m.Partition[index-1].PartSize
				}

				ebr.PartStatusE = '0'
				copy(ebr.PartNameE[:], name)
				ebr.PartFitE = fit
				ebr.PartNextE = -1
				ebr.PartSizeE = 0
				ebr.PartStartE = m.Partition[index].PartStart
				writeInExtenderPartition(path, m, ebr, index)
				fmt.Println("SE CREO LA PARTICION EXTENDIDA, SI SALE AUN ARCHIVOS :D")
			} else {
				fmt.Println("YA SE HA CREADO 4 PARTICIONES YA NO SE PUEDEN CREAR MAS :(")
			}
		} else {
			fmt.Println("LA PARTICION YA HA SIDO CREADA CON ESTE NOMBRE ", name)
		}
	} else {
		fmt.Println("YA NO CUENTA CON ESPACIO EN EL DISCO, EL DISTO TIENE UN ESPACIO DE: ", string(m.MbrSize), " Y ESTA TRATANDO DE HACER UNA PARTICION DE ", string(size))
	}

}

// FirstFit Crear el primer ajuste
func FirstFit(m structs_lwh.MBR, size int32) int {
	var check bool = false
	fmt.Println(check)
	var indexSelect int = 0
	for ; indexSelect < 4; indexSelect++ {
		if m.Partition[indexSelect].PartStart == -1 || (m.Partition[indexSelect].PartStatus == '1' && m.Partition[indexSelect].PartSize >= size) {
			check = true
			break
		}
	}
	if check {
		return indexSelect
	}
	return -1
}

// BestFit crear el mejor ajuste
func BestFit(m structs_lwh.MBR, size int32) int {
	var check bool = false
	var indexSelect int = 0
	for i := 0; i < 4; i++ {
		if m.Partition[i].PartStart == -1 || (m.Partition[i].PartStatus == '1' && m.Partition[i].PartSize >= size) {
			check = true
			if i != indexSelect && m.Partition[indexSelect].PartSize > m.Partition[i].PartSize {
				indexSelect = i
			}
		}
	}
	if check {
		return indexSelect
	}
	return -1
}

// WorstFit el peor ajuste :(
func WorstFit(m structs_lwh.MBR, size int32) int {
	var check bool = false
	var indexSelect int = 0
	for i := 0; i < 4; i++ {
		if m.Partition[i].PartStart == -1 || (m.Partition[i].PartStatus == '1' && m.Partition[i].PartSize >= size) {
			check = true
			if i != indexSelect && m.Partition[indexSelect].PartSize < m.Partition[i].PartSize {
				indexSelect = i
			}
		}
	}
	if check {
		return indexSelect
	}
	return -1
}

func readFileDisk(f *os.File, err error) structs_lwh.MBR {
	var m structs_lwh.MBR
	sizeRead := binary.Size(m)
	fmt.Printf("%+v\n", sizeRead)
	data := readNextBytes(f, sizeRead)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failled", err)
	}
	return m
}

func readFileEBR(f *os.File, err error) structs_lwh.EBR {
	var e structs_lwh.EBR
	sizeRead := binary.Size(e)
	fmt.Printf("%+v\n", sizeRead)
	data := readNextBytes(f, sizeRead)
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &e)
	if err != nil {
		log.Fatal("binary.Read failled", err)
	}
	return e
}

func converNameToByte(name string) [16]byte {
	var auxName [16]byte
	for i, j := range []byte(name) {
		auxName[i] = byte(j)
	}
	return auxName
}

func writeInPrimaryPartition(path string, m structs_lwh.MBR, index int) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//var firstPos byte = '1'

	//start := &firstPos
	//struC := &m
	f.Seek(int64(m.MbrSize), int(m.Partition[index].PartStart))
	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &m)

	writeNextBytes(f, binario.Bytes())
}

func writeInExtenderPartition(path string, m structs_lwh.MBR, e structs_lwh.EBR, index int) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	f.Seek(int64(m.MbrSize), int(m.Partition[index].PartStart))
	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &m)
	binary.Write(&binario, binary.BigEndian, &e)

	writeNextBytes(f, binario.Bytes())

}
