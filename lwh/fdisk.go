package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// MakeFdisk crear las particiones primarias, extendidas y logicas
func MakeFdisk(root Node) {
	var path string = ""
	var name string = ""
	var size int64 = 0
	var unit byte = 'K'
	var tipo byte = 'P'
	var fit byte = 'F'
	var del int64 = -1
	var add int64 = 0
	for _, i := range root.Children {
		if i.TypeToken == "SIZE" {
			k, err := strconv.ParseInt(i.Value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			size = k
		} else if i.TypeToken == "UNIT" {
			if strings.EqualFold(i.Value, "k") {
				unit = 'K'
			} else if strings.EqualFold(i.Value, "m") {
				unit = 'M'
			} else if strings.EqualFold(i.Value, "b") {
				unit = 'B'
			}
		} else if i.TypeToken == "PATH" {
			path = i.Value
		} else if i.TypeToken == "TYPE" {
			if strings.EqualFold(i.Value, "p") {
				tipo = 'P'
			} else if strings.EqualFold(i.Value, "e") {
				tipo = 'E'
			} else if strings.EqualFold(i.Value, "l") {
				tipo = 'L'
			}
		} else if i.TypeToken == "FIT" {
			if strings.EqualFold(i.Value, "BF") {
				fit = 'B'
			} else if strings.EqualFold(i.Value, "FF") {
				fit = 'F'
			} else if strings.EqualFold(i.Value, "WF") {
				unit = 'W'
			}
		} else if i.TypeToken == "DELETE" {
			if strings.EqualFold(i.Value, "fast") {
				del = 0
			} else if strings.EqualFold(i.Value, "full") {
				del = 1
			}
		} else if i.TypeToken == "NAME" {
			name = i.Value
		} else if i.TypeToken == "ADD" {
			k, err := strconv.ParseInt(i.Value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			add = k
			fmt.Println(add)
		}
		for _, j := range i.Children {
			if j.TypeToken == "SIZE" {
				k, err := strconv.ParseInt(j.Value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				size = k
			} else if j.TypeToken == "UNIT" {
				if strings.EqualFold(j.Value, "k") {
					unit = 'K'
				} else if strings.EqualFold(j.Value, "m") {
					unit = 'M'
				} else if strings.EqualFold(j.Value, "b") {
					unit = 'B'
				}
			} else if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "TYPE" {
				if strings.EqualFold(j.Value, "p") {
					tipo = 'P'
				} else if strings.EqualFold(j.Value, "e") {
					tipo = 'E'
				} else if strings.EqualFold(j.Value, "l") {
					tipo = 'L'
				}
			} else if j.TypeToken == "FIT" {
				if strings.EqualFold(j.Value, "BF") {
					fit = 'B'
				} else if strings.EqualFold(j.Value, "FF") {
					fit = 'F'
				} else if strings.EqualFold(j.Value, "WF") {
					unit = 'W'
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
			} else if j.TypeToken == "ADD" {
				k, err := strconv.ParseInt(j.Value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				add = k
				fmt.Println(add)
			}
		}
	}

	if _, err := os.Stat(path); err == nil {
		if size != 0 {
			if tipo == 'P' {
				ParticionPrimaria(path, name, tipo, fit, unit, size)
				printMBR(path)
			} else if tipo == 'E' {
				ParticionExtendida(path, name, tipo, fit, unit, size)
				printMBR(path)
			} else if tipo == 'L' {
				ParticionLogica(path, name, tipo, fit, unit, size)
				printMBR(path)
			}
		} else if del != -1 {

		} else {
			AddSize(path, name, unit, add)
		}
	} else if os.IsNotExist(err) {
		panic(err)
	}

}

// AddSize anade mas espacio o quita dependiendo del valor que tenga el comando add
func AddSize(path string, name string, unit byte, size int64) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	addOrM := strconv.Itoa(int(size))
	flagCheck := CheckNumbers(addOrM)
	m := readFileDisk(f, err)
	var index = 0
	var existPartition = false
	var part structs_lwh.Partitions

	for k, i := range m.Partition {
		if auxName == i.PartName {
			existPartition = true
			index = k
			part = i
			break
		}
	}

	if existPartition {
		if flagCheck {
			if math.Abs(float64(auxSize)) >= float64(part.PartSize) {
				fmt.Println("EL ESPACIO A QUITAR ES MAS GRANDE QUE EL ESPACIO DE LA PARTICION")
			} else {
				part.PartSize = part.PartSize - int64(math.Abs(float64(auxSize)))
				m.Partition[index] = part
				f.Seek(0, 0)
				var binario bytes.Buffer

				binary.Write(&binario, binary.BigEndian, &m)

				writeNextBytes(f, binario.Bytes())
			}
		} else {
			if index == 3 {
				x := m.MbrSize - (part.PartStart + part.PartSize)
				if x >= auxSize {
					part.PartSize = part.PartSize + auxSize
					m.Partition[index] = part
					f.Seek(0, 0)
					var binario bytes.Buffer

					binary.Write(&binario, binary.BigEndian, &m)

					writeNextBytes(f, binario.Bytes())
				} else {
					fmt.Println("EL ESPACION ANADIR ES MAYOR AL QUE TIENE LA PARTICION")
				}
			} else {
				var nextPartition structs_lwh.Partitions
				var rest int64 = 0
				var nextIndex = index + 1
				for nextIndex < 5 {
					if nextIndex == 4 {
						rest = m.MbrSize - (part.PartStart + part.PartSize)
						break
					} else {
						nextPartition = m.Partition[nextIndex]
						nextIndex++
						if nextPartition.PartStatus == '0' && nextPartition.PartStart != -1 {
							rest = nextPartition.PartStart - (part.PartStart + part.PartSize)
							break
						}
					}
				}
				if rest >= auxSize {
					part.PartSize = part.PartSize + auxSize
					m.Partition[index] = part
					f.Seek(0, 0)
					var binario bytes.Buffer

					binary.Write(&binario, binary.BigEndian, &m)

					writeNextBytes(f, binario.Bytes())
				} else {
					fmt.Println("EL ESPACIO A AGREGAR ES MAYOR A LO DISPONIBLE PARA LA SIGUIENTE PARTICION")
				}
			}
		}
	}

}

// ParticionPrimaria este metodo es usado para crear la particion primaria en el disco
func ParticionPrimaria(path string, name string, tipo byte, fit byte, unit byte, size int64) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	m := readFileDisk(f, err)
	var ebr structs_lwh.EBR
	//ReportMBR("/home/dani3l8200/Escritorio/MisDiscos/archivo.dot", m)

	var a int64 = 0
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
				f.Seek(m.Partition[i].PartStart, 0)
				ebr = readFileEBR(f, err)
				tam := binary.Size(ebr)
				pos, _ := f.Seek(0, os.SEEK_CUR)
				fmt.Println(pos)
				for tam != 0 && pos < int64(m.Partition[i].PartSize)+int64(m.Partition[i].PartStart) {
					if ebr.PartNameE == auxName {
						checkPartition = true
					} else if ebr.PartNextE == -1 {
						break
					} else {
						f.Seek(ebr.PartNextE, 0)
						ebr = readFileEBR(f, err)
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
					m.Partition[index].PartStart = int64(binary.Size(m)) + 1
				} else {
					m.Partition[index].PartStart = m.Partition[index-1].PartStart + m.Partition[index-1].PartSize
				}

				f.Seek(0, 0)

				var binario bytes.Buffer

				binary.Write(&binario, binary.BigEndian, &m)

				writeNextBytes(f, binario.Bytes())

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
func ParticionExtendida(path string, name string, tipo byte, fit byte, unit byte, size int64) {
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	m := readFileDisk(f, err)
	var ebr structs_lwh.EBR
	var a int64 = 0

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
				pos, _ := f.Seek(0, os.SEEK_CUR)
				e := readFileEBR(f, err)
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
					m.Partition[index].PartStart = int64(binary.Size(m)) + 1
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

func printMBR(path string) {
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Seek(0, 0)
	m := readFileDisk(f, err)
	var indexCheck int = 0
	var fecha string = ""
	var nameParticion string = ""
	var exiteEBR bool = false
	var e structs_lwh.EBR
	fmt.Println("--------------------------------------MBR-------------------------------------")
	fmt.Println("SIZE:", strconv.Itoa(int(m.MbrSize)))
	fmt.Println("Signature:", strconv.Itoa(int(m.MbrDiskSignature)))
	for _, k := range m.MbrTime {
		if k != 0 {
			fecha += string(k)
		}
	}
	fmt.Println("FECHA:", fecha)
	for i, mbr := range m.Partition {
		fmt.Println("PARTCION NO.", strconv.Itoa(i+1))
		fmt.Println("FIT:", string(mbr.PartFit))
		for _, name := range mbr.PartName {
			if name != 0 {
				nameParticion += string(name)
			}
		}
		fmt.Println("Nombre Particion:", nameParticion)
		fmt.Println("Size de Particion:", strconv.Itoa(int(mbr.PartSize)))
		fmt.Println("Start de paticion:", strconv.Itoa(int(mbr.PartStart)))
		fmt.Println("Status de Particion:", string(mbr.PartStatus))
		fmt.Println("Tipo de Particion:", string(mbr.PartType))
		if mbr.PartType == 'E' {
			indexCheck = i
			exiteEBR = true
		}
		nameParticion = ""
	}
	x := 0
	if exiteEBR {
		f.Seek(m.Partition[indexCheck].PartStart, 0)
		e = readFileEBR(f, err)
		pos, _ := f.Seek(0, os.SEEK_CUR)
		fmt.Println("---------------------EBR--------------------------------")
		for pos < (m.Partition[indexCheck].PartStart + m.Partition[indexCheck].PartSize) {

			if e.PartStatusE != '1' {
				fmt.Println("PARTICION LOGICA NO.", strconv.Itoa(x+1))
				fmt.Println("FIT_E:", string(e.PartFitE))
				for _, k := range e.PartNameE {
					if k != 0 {
						nameParticion += string(k)
					}
				}
				fmt.Println("NAME_E:", nameParticion)
				fmt.Println("NEXT_E:", strconv.Itoa(int(e.PartNextE)))
				fmt.Println("SIZE_E:", strconv.Itoa(int(e.PartSizeE)))
				fmt.Println("START_E:", strconv.Itoa(int(e.PartStartE)))
				fmt.Println("STATUS_E:", string(e.PartStatusE))
				nameParticion = ""
			}
			if e.PartNextE == -1 {
				break
			} else {
				f.Seek(e.PartNextE, 0)
				e = readFileEBR(f, err)
			}
			x++

		}

	}

}

//ParticionLogica Crea particiones logicas
func ParticionLogica(path string, name string, tipo byte, fit byte, unit byte, size int64) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	auxSize := verifySize(unit, size)
	auxName := converNameToByte(name)
	fmt.Printf("%+v\n", auxSize)
	f.Seek(0, 0)
	m := readFileDisk(f, err)
	var ebr structs_lwh.EBR
	var checkPartition bool = false

	for i := 0; i < 4; i++ {
		if m.Partition[i].PartName == auxName {
			checkPartition = true
			break
		} else if m.Partition[i].PartType == 'E' {
			f.Seek(int64(m.Partition[i].PartStart), 0)
			ebr = readFileEBR(f, err)
			pos, _ := f.Seek(0, os.SEEK_CUR)
			fmt.Println(pos)
			eSize := binary.Size(ebr)
			fmt.Println(eSize)
			for eSize != 0 && pos < int64(m.Partition[i].PartSize)+int64(m.Partition[i].PartStart) {
				if ebr.PartNameE == auxName {
					checkPartition = true
				}
				if ebr.PartNextE == -1 {
					break
				} else {
					f.Seek(ebr.PartNextE, 0)
					ebr = readFileEBR(f, err)
				}
			}
		}
	}
	if !checkPartition {
		var index int = -1

		for i := 0; i < 4; i++ {
			if m.Partition[i].PartType == 'E' {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Println("EL ENUNCIADO INDICA QUE PARA LOGICA SE NECESITA UNA EXTENDIDA :)")
		}
		f.Seek(m.Partition[index].PartStart, 0)
		ebr = readFileEBR(f, err)
		if ebr.PartSizeE == 0 {
			if m.Partition[index].PartSize < auxSize {
				fmt.Println("El size de la particion logica debe ser menor a la extendida ya creada :p")
			} else {
				sizeM := binary.Size(ebr)
				pos, _ := f.Seek(0, os.SEEK_CUR)
				fmt.Println(pos)
				ebr.PartStatusE = '0'
				ebr.PartFitE = fit
				ebr.PartStartE = pos - int64(sizeM)
				ebr.PartSizeE = auxSize
				ebr.PartNextE = -1
				copy(ebr.PartNameE[:], name)
				writeInLogicalPartition(path, m, ebr, index)
				f.Seek(m.Partition[index].PartStart, 0)
				auxMI := readFileEBR(f, err)
				fmt.Println(auxMI)
				fmt.Println("SE CREO LA PARTICION LOGICA")
			}
		} else {
			pos, _ := f.Seek(0, os.SEEK_CUR)
			for ebr.PartNextE != -1 && pos < (m.Partition[index].PartSize+m.Partition[index].PartStart) {
				f.Seek(ebr.PartNextE, 0)
				ebr = readFileEBR(f, err)

			}
			size1 := ebr.PartStartE + ebr.PartSizeE + auxSize
			if size1 <= (m.Partition[index].PartSize + m.Partition[index].PartStart) {
				ebr.PartNextE = ebr.PartStartE + ebr.PartSizeE

				sizeM := binary.Size(ebr)
				pos, _ := f.Seek(0, os.SEEK_CUR)
				b := pos - int64(sizeM)
				f.Seek(b, 0)

				var binario bytes.Buffer

				binary.Write(&binario, binary.BigEndian, &ebr)

				writeNextBytes(f, binario.Bytes())

				f.Seek(int64(ebr.PartStartE+ebr.PartSizeE), 0)
				pos1, _ := f.Seek(0, os.SEEK_CUR)
				ebr.PartFitE = fit
				copy(ebr.PartNameE[:], name)
				ebr.PartNextE = -1
				ebr.PartSizeE = auxSize
				ebr.PartStartE = pos1
				ebr.PartStatusE = '0'

				var binario2 bytes.Buffer

				binary.Write(&binario2, binary.BigEndian, &ebr)

				writeNextBytes(f, binario2.Bytes())

				fmt.Println("SE CREO LA PARTICION LOGICA")
			} else {
				fmt.Println("LA PARTICION LOGICA SOBREPASA EL PESO DE LA EXTENDIDA")
			}
		}

	} else {
		fmt.Println("LA PARTICION LOGICA CON ESTE NOMBRE YA HA SIDO CREADA", name)
	}
}

// FirstFit Crear el primer ajuste
func FirstFit(m structs_lwh.MBR, size int64) int {
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
func BestFit(m structs_lwh.MBR, size int64) int {
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
func WorstFit(m structs_lwh.MBR, size int64) int {
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

func writeInLogicalPartition(path string, m structs_lwh.MBR, e structs_lwh.EBR, index int) {
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(int64(m.Partition[index].PartStart), 0)
	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &e)

	writeNextBytes(f, binario.Bytes())
}

func writeInExtenderPartition(path string, m structs_lwh.MBR, e structs_lwh.EBR, index int) {
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Seek(0, 0)

	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &m)

	writeNextBytes(f, binario.Bytes())

	f.Seek(m.Partition[index].PartStart, 0)

	fmt.Println(strconv.Itoa(int(m.Partition[index].PartStart)))

	var binario2 bytes.Buffer

	binary.Write(&binario2, binary.BigEndian, &e)

	writeNextBytes(f, binario2.Bytes())
}

func writeInExtenderPartitionIfExist(path string, m structs_lwh.MBR, e structs_lwh.EBR, index int, fit byte, auxSize int32) {

}

//CheckNumbers revisa si el valor para el comando add es positivio o negativo
func CheckNumbers(value string) bool {
	var check = false
	for _, char := range value {
		if char < '0' || char > '9' {
			check = true
			return check
		}
	}
	return check
}
