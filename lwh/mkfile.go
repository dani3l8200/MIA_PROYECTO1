package lwh

import (
	"MIA-PROYECTO1/datastructure"
	"MIA-PROYECTO1/structs_lwh"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//MakeFile ...
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

	disk, err := Lista.GetMountedPart(id)
	if err {
		getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

		if getData.GetSize != 0 && getData.GetStart != 0 {
			auxPath, errorPath := SetDirectory(path)
			if errorPath {
				path = auxPath
				RecorrerAVD(disk.GetPath(), getData.GetStart, size, path, cont, p)

			} else if errorPath == false {
				log.Fatal("ERROR CON EL PATH")
			}

		}

	} else if err == false {
		fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
	}
}

//RecorrerAVD ....
func RecorrerAVD(pathDisk string, start int64, size int64, path string, cont string, p bool) {
	f, err := os.OpenFile(pathDisk, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	auxPath := strings.Split(path, "/")

	auxPath = deleteEmpty(auxPath)

	sb := readFileSB(f, err, start)

	var pos int = 0

	var apIndirecto int64 = 0

	avd := ReadAVD(apIndirecto, f, err, sb.SbApTreeDirectory)

	if p {
		for pos < len(auxPath) {
			if pos < len(auxPath)-1 {
				if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					if WriteFilesAVD(start, f, err, auxPath[pos], apIndirecto, myusers.UID, myusers.Gid, avd, sb) {
						sb = readFileSB(f, err, start)
						avd = ReadAVD(apIndirecto, f, err, sb.SbApTreeDirectory)
					} else if !WriteFilesAVD(start, f, err, auxPath[pos], apIndirecto, myusers.UID, myusers.Gid, avd, sb) {
						fmt.Println("EL USUARIO NO TIENE PERMISOS!")
						return
					}
				}

				apIndirecto = existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos])

				if len(GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid).AvdNameDirectory) != 0 {
					sb = readFileSB(f, err, start)
					avd = GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid)

				} else if len(GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid).AvdNameDirectory) == 0 {
					fmt.Println("NO TIENE PERMISOS DE LECTURA!")
					return
				}
				pos++

			} else if pos == len(auxPath)-1 {
				sb = readFileSB(f, err, start)
				if !existDetailDirectory(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), auxPath[pos]) {
					sb = readFileSB(f, err, start)
					WriteFilesDD(start, f, err, sb, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), avd.AvdApDetailDirectory, auxPath[pos], size, cont, 1, 0, 664)

					pos++

					fmt.Println("SE CREO EL ARCHIVO CON EXITO")
				} else {
					fmt.Println("YA EXISTE EL ARCHIVO")
					return
				}
			}
		}
	} else {
		for pos < len(auxPath) {

			if pos < len(auxPath)-1 {
				if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					log.Fatalln("NO EXISTE LA CARPETA USE EL COMANDO P")
					return
				}

				apIndirecto = existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos])

				if len(GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid).AvdNameDirectory) != 0 {
					sb = readFileSB(f, err, start)
					avd = GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid)

				} else {
					log.Fatalln("NO TIENE PERMISOS DE LECTURA")
					return
				}

				pos++
			} else if pos == len(auxPath)-1 {
				sb = readFileSB(f, err, start)
				if !existDetailDirectory(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), auxPath[pos]) {
					sb = readFileSB(f, err, start)
					WriteFilesDD(start, f, err, sb, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), avd.AvdApDetailDirectory, auxPath[pos], size, cont, 1, 0, 664)

					pos++

					fmt.Println("SE CREO EL ARCHIVO CON EXITO")
				} else {
					fmt.Println("YA EXISTE EL ARCHIVO")
					return
				}
			}

		}
	}

}

//WriteFilesAVD ...
func WriteFilesAVD(start int64, f *os.File, err error, name string, aptIndirecto int64, idUser int64, idGrp int64, avd structs_lwh.AVD, sb structs_lwh.SB) bool {
	if sb.SbTreeVirtualFree != 0 && sb.SbDetailDirectoryFree != 0 {
		perm := Permisos(avd.AvdPerm)

		if GetPWrite(avd.AvdProper, idUser, perm) {

			for i := 0; i < 6; i++ {
				if avd.AvdApArraySubdirectories[i] == -1 {

					avdBitmap := getSBBitMap(f, err, sb.SbTreeVirtualCount, sb.SbApBitmapTreeDirectory)

					ddBitmap := getSBBitMap(f, err, sb.SbDetailDirectoryCount, sb.SbApBitmapDetailDirectory)

					bitFreeAvd := GetNextBitFree(avdBitmap)

					bitFreeDd := GetNextBitFree(ddBitmap)

					avd.AvdApArraySubdirectories[i] = bitFreeAvd

					sb.SbTreeVirtualFree = sb.SbTreeVirtualFree - 1

					sb.SbDetailDirectoryFree = sb.SbDetailDirectoryFree - 1

					avdBitmap[bitFreeAvd] = '1'

					ddBitmap[bitFreeDd] = '1'

					dd := createDDirectory()

					avdFile := createAVD(name, bitFreeDd, idUser, idGrp, 664)

					writeDD(bitFreeDd, f, err, sb, dd)

					WriteAVD(bitFreeAvd, f, err, sb, avdFile)

					WriteAVD(aptIndirecto, f, err, sb, avd)

					var fakeBlock []byte
					var fakeInodo []byte

					sb = WriteOneInBitmap(f, sb, avdBitmap, ddBitmap, fakeInodo, fakeBlock)

					writeSB(f, err, sb, start)

					return true
				}
			}

			if avd.AvdApTreeVirtualDirectory != -1 {
				return WriteFilesAVD(start, f, err, name, avd.AvdApTreeVirtualDirectory, idUser, idGrp, ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, sb.SbApTreeDirectory), sb)
			} else if avd.AvdApTreeVirtualDirectory == -1 {
				avdBitmap := getSBBitMap(f, err, sb.SbTreeVirtualCount, sb.SbApBitmapTreeDirectory)

				ddBitmap := getSBBitMap(f, err, sb.SbDetailDirectoryCount, sb.SbApBitmapDetailDirectory)

				bitFreeAvd := GetNextBitFree(avdBitmap)

				sb.SbTreeVirtualFree = sb.SbTreeVirtualFree - 1

				avdBitmap[bitFreeAvd] = '1'

				avd.AvdApTreeVirtualDirectory = bitFreeAvd

				apIndirecto := createAVD(converByteLToString(avd.AvdNameDirectory), -1, idUser, idGrp, 664)

				WriteAVD(bitFreeAvd, f, err, sb, apIndirecto)

				WriteAVD(aptIndirecto, f, err, sb, avd)

				var fakeBlock []byte

				var fakeInodo []byte

				sb = WriteOneInBitmap(f, sb, avdBitmap, ddBitmap, fakeInodo, fakeBlock)

				writeSB(f, err, sb, start)

				WriteFilesAVD(start, f, err, name, bitFreeAvd, idUser, idGrp, apIndirecto, sb)

			}
		}
	} else if sb.SbTreeVirtualFree == 0 && sb.SbDetailDirectoryFree == 0 {
		fmt.Println("ESTRUCTURAS YA LLENAS ")
	}
	return false
}

//WriteFilesDD ...
func WriteFilesDD(start int64, f *os.File, err error, sb structs_lwh.SB, dd structs_lwh.DDirectory, aptDD int64, name string, size int64, cont string, idUser int64, idGrp int64, perm int64) bool {
	if sb.SbDetailDirectoryFree != 0 {
		inodoBitmap := getSBBitMap(f, err, sb.SbInodosCount, sb.SbApBitmapTableInodo)

		inodoFree := GetNextBitFree(inodoBitmap)

		auxInodo := createTableInodo(inodoFree, size, idUser, idGrp, perm)

		for i := 0; i < 5; i++ {
			if dd.DDArrayBlock[i].DdFileApInodo == -1 {

				inodoBitmap[inodoFree] = '1'

				dd.DDArrayBlock[i] = createArrayFile(name, inodoFree)
				sb.SbInodosFree--

				writeDD(aptDD, f, err, sb, dd)

				writeTInodo(inodoFree, f, err, sb, auxInodo)

				var fakeAVD []byte
				var fakeDD []byte
				var fakeBlock []byte

				sb = WriteOneInBitmap(f, sb, fakeAVD, fakeDD, inodoBitmap, fakeBlock)

				writeSB(f, err, sb, start)

				if cont != "" {

					WriteFilesInodos(f, err, sb, auxInodo, inodoFree, cont, start)

					return true

				} else if cont == "" {

					cont = GenerateContForEmpty(size)

					WriteFilesInodos(f, err, sb, auxInodo, inodoFree, cont, start)

					return true
				}

			}
		}

		if dd.DdApDetailDirectory != -1 {
			return WriteFilesDD(start, f, err, sb, ReadDD(dd.DdApDetailDirectory, f, err, sb.SbApDetailDirectory), dd.DdApDetailDirectory, name, size, cont, idUser, idGrp, perm)
		} else if dd.DdApDetailDirectory == -1 {

			ddBitmap := getSBBitMap(f, err, sb.SbDetailDirectoryCount, sb.SbApBitmapDetailDirectory)

			ddFree := GetNextBitFree(ddBitmap)

			ddBitmap[ddFree] = '1'

			sb.SbDetailDirectoryFree--

			dd.DdApDetailDirectory = ddFree

			auxDD := createDDirectory()

			writeDD(aptDD, f, err, sb, dd)

			writeDD(ddFree, f, err, sb, auxDD)

			sb = WriteOneInBitmap(f, sb, nil, ddBitmap, nil, nil)

			writeSB(f, err, sb, start)

			return WriteFilesDD(start, f, err, sb, auxDD, ddFree, name, size, cont, idUser, idGrp, perm)
		}
	} else if sb.SbDetailDirectoryFree == 0 {
		fmt.Println("YA OCUPO TODAS LAS ESTRUCTURAS DISPONIBLES")
	}

	return false
}

//WriteFilesInodos ...
func WriteFilesInodos(f *os.File, err error, sb structs_lwh.SB, inodo structs_lwh.INodo, aptInodo int64, cont string, start int64) bool {
	if sb.SbInodosFree != 0 {

		inodoBitmap := getSBBitMap(f, err, sb.SbInodosCount, sb.SbApBitmapTableInodo)

		blockBitmap := getSBBitMap(f, err, sb.SbBlocksCount, sb.SbApBitmapBlocks)

		var limit [25]byte

		var block structs_lwh.Block

		for i := 0; i < 4; i++ {

			if len(cont) > 25 {

				if inodo.IArrayBlocks[i] == -1 {

					var dataN string = ""

					copy(limit[:], cont)

					data := convertBByteToString(limit)

					blockFree := GetNextBitFree(blockBitmap)

					blockBitmap[blockFree] = '1'

					inodo.IArrayBlocks[i] = blockFree

					block = createBlock(data)

					sb.SbBlocksFree--

					writeBlock(blockFree, f, err, sb, block)

					limitN := make([]byte, len(cont))

					for j := 25; j < len(cont); j++ {

						limitN[j] = cont[j]

					}

					for _, j := range limitN {
						if j != 0 {
							dataN += string(j)
						}
					}

					cont = dataN

				} else if inodo.IArrayBlocks[i] != -1 {

					var dataN string = ""

					copy(limit[:], cont)

					data := convertBByteToString(limit)

					block = createBlock(data)

					writeBlock(inodo.IArrayBlocks[i], f, err, sb, block)

					limitN := make([]byte, len(cont))

					for j := 25; j < len(cont); j++ {

						limitN[j] = cont[j]

					}

					for _, j := range limitN {
						if j != 0 {
							dataN += string(j)
						}
					}

					cont = dataN
				}
			} else {
				if inodo.IArrayBlocks[i] == -1 {

					blockFree := GetNextBitFree(blockBitmap)

					blockBitmap[blockFree] = '1'

					inodo.IArrayBlocks[i] = blockFree

					block = createBlock(cont)

					sb.SbBlocksFree--

					writeBlock(blockFree, f, err, sb, block)

					cont = ""

					writeTInodo(aptInodo, f, err, sb, inodo)

					var fakeAvd []byte
					var fakeDD []byte
					var fakeInodo []byte

					sb = WriteOneInBitmap(f, sb, fakeAvd, fakeDD, fakeInodo, blockBitmap)

					writeSB(f, err, sb, start)

					return true
				} else if inodo.IArrayBlocks[i] != -1 {

					block = createBlock(cont)

					writeBlock(inodo.IArrayBlocks[i], f, err, sb, block)

					cont = ""

					writeTInodo(aptInodo, f, err, sb, inodo)

					var fakeAvd []byte
					var fakeDD []byte
					var fakeInodo []byte

					sb = WriteOneInBitmap(f, sb, fakeAvd, fakeDD, fakeInodo, blockBitmap)

					writeSB(f, err, sb, start)
					return true
				}
			}
		}

		if inodo.IApIndirecto == -1 {

			inodoFree := GetNextBitFree(inodoBitmap)

			inodo.IApIndirecto = inodoFree

			sb.SbInodosFree--

			inodoBitmap[inodoFree] = '1'

			auxInodo := createTableInodo(inodoFree, inodo.ISizeArchive, inodo.IProper, inodo.IGid, inodo.IPerm)

			writeTInodo(inodoFree, f, err, sb, auxInodo)

			writeTInodo(aptInodo, f, err, sb, inodo)

			sb = WriteOneInBitmap(f, sb, nil, nil, inodoBitmap, blockBitmap)

			writeSB(f, err, sb, start)

			return WriteFilesInodos(f, err, sb, auxInodo, inodoFree, cont, start)

		} else if inodo.IApIndirecto != -1 {

			writeTInodo(aptInodo, f, err, sb, inodo)

			sb = WriteOneInBitmap(f, sb, nil, nil, inodoBitmap, blockBitmap)

			writeSB(f, err, sb, start)

			return WriteFilesInodos(f, err, sb, ReadTInodo(inodo.IApIndirecto, f, err, sb.SbApTableInodo), inodo.IApIndirecto, cont, start)
		}
	} else if sb.SbInodosFree == 0 && sb.SbBlocksFree == 0 {
		fmt.Println("YA OCUPO TODAS LAS ESTRUCTURAS DISPONIBLES")
	}
	return false
}

//GetFilesAVD ...
func GetFilesAVD(f *os.File, err error, start int64, avd structs_lwh.AVD, name string, idUser int64, idGrp int64) structs_lwh.AVD {
	auxAvd := createAVD("", -1, 0, 0, 664)
	perm := Permisos(avd.AvdPerm)
	auxName := converNameSBToByte(name)
	if GetPRead(avd.AvdProper, idUser, perm) {
		for i := 0; i < 6; i++ {
			if avd.AvdApArraySubdirectories[i] != -1 {
				if ReadAVD(avd.AvdApArraySubdirectories[i], f, err, start).AvdNameDirectory == auxName {
					return ReadAVD(avd.AvdApArraySubdirectories[i], f, err, start)
				}
			}
		}

		if avd.AvdApTreeVirtualDirectory != -1 {
			return GetFilesAVD(f, err, start, ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, start), name, idUser, idGrp)
		}
	}

	return auxAvd
}

//SetDirectory ...
func SetDirectory(path string) (string, bool) {

	//index := strings.LastIndex(path, "/")
	if strings.Contains(path, "\"") {
		path = strings.ReplaceAll(path, "\"", "")
		/*if index > -1 {
			path = path[:index]

		}*/
		return path, true
	} else if !strings.Contains(path, "\"") {
		/*if index > -1 {
			path = path[:index]
			return path, true
		}*/
	}
	return path, true
}

func createPointer(name string, ppointer int64) structs_lwh.Pointer {
	var pointer structs_lwh.Pointer

	pointer.Name = converNameDDToByte(name)
	pointer.PPointer = ppointer

	return pointer
}

//CreateListPointerAvd obtiene las carpetas del avd
func CreateListPointerAvd(f *os.File, err error, SbApTreeDirectory int64, avd structs_lwh.AVD) datastructure.LinkedListP {

	var listaP datastructure.LinkedListP
	listaP.Delete()
	for i := 0; i < 6; i++ {
		if avd.AvdApArraySubdirectories[i] != -1 {
			aux := ReadAVD(avd.AvdApArraySubdirectories[i], f, err, SbApTreeDirectory)

			listaP.Insert(createPointer(converByteLToString(aux.AvdNameDirectory), avd.AvdApArraySubdirectories[i]))
		}

	}

	if avd.AvdApTreeVirtualDirectory != -1 {

		aux := ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, SbApTreeDirectory)

		auxList := CreateListPointerAvd(f, err, SbApTreeDirectory, aux)

		if auxList.Head != nil {
			for node := auxList.Head; node != nil; node = node.Next() {

				point := node.Value()

				listaP.Insert(point)
			}
		}

	}

	return listaP
}

//CreateListPointerDd ...
func CreateListPointerDd(f *os.File, err error, SbApDetailDirectory int64, dd structs_lwh.DDirectory) datastructure.LinkedListP {
	var listaP datastructure.LinkedListP
	listaP.Delete()

	for i := 0; i < 5; i++ {
		if dd.DDArrayBlock[i].DdFileApInodo != -1 {
			listaP.Insert(createPointer(convertBByteToString(dd.DDArrayBlock[i].DdFileName), dd.DDArrayBlock[i].DdFileApInodo))
		}
	}

	if dd.DdApDetailDirectory != -1 {
		auxList := CreateListPointerDd(f, err, SbApDetailDirectory, ReadDD(dd.DdApDetailDirectory, f, err, SbApDetailDirectory))

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

func existDetailDirectory(f *os.File, err error, startDD int64, dd structs_lwh.DDirectory, name string) bool {
	auxName := converNameDDToByte(name)
	for i := 0; i < 5; i++ {
		if dd.DDArrayBlock[i].DdFileApInodo != -1 {
			if dd.DDArrayBlock[i].DdFileName == auxName {
				return true
			}
		}
	}

	if dd.DdApDetailDirectory != -1 {
		return existDetailDirectory(f, err, startDD, ReadDD(dd.DdApDetailDirectory, f, err, startDD), name)
	}
	return false
}

func converNameSBToByte(name string) [20]byte {
	var auxName [20]byte
	for i, j := range []byte(name) {
		auxName[i] = byte(j)
	}
	return auxName
}

func converNameDDToByte(name string) [25]byte {
	var auxName [25]byte
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

func convertPartitionNameToString(name [16]byte) string {
	var DbData string = ""
	for _, name := range name {
		if name != 0 {
			DbData += string(name)
		}
	}
	return DbData
}

func getSBBitMap(f *os.File, err error, sizeBitmap int64, x int64) []byte {
	var i int64 = 0

	var data0 byte = ' '

	test := make([]byte, sizeBitmap)

	for ; i < sizeBitmap; i++ {

		f.Seek(x+i, io.SeekStart)

		sizeRead := binary.Size(data0)

		data := readNextBytes(f, sizeRead)

		buffer := bytes.NewBuffer(data)

		err = binary.Read(buffer, binary.BigEndian, &data0)

		if err != nil {
			log.Fatal("binary.Read failed", err)
		}

		test[i] = data0
	}
	return test
}

//GenerateContForEmpty ...
func GenerateContForEmpty(size int64) string {
	var name string = ""
	var mybyte byte = 97
	var i int64 = 0
	var j int64 = 0
	for i <= 25 && j <= size {
		if i < 25 {
			if mybyte <= 122 {
				name += string(mybyte)
				mybyte++
			}
		} else if i == 25 {
			mybyte = 97
			i = -1
		}
		i++
		j++
	}
	return name
}
