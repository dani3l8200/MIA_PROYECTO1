package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var flagCheckAVD bool = false
var flagCheckDD bool = false
var flagCheckInodo bool = false
var flagCheckBlock bool = false

//GetDiskMount Obtiene las particiones y informacion de las listas
func GetDiskMount(path string, name string, part bool) structs_lwh.GetMounDisk {
	var Data structs_lwh.GetMounDisk
	var m structs_lwh.MBR
	var e structs_lwh.EBR
	auxName := converNameToByte(name)
	Data.GetSize = 0
	Data.GetStart = 0
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Seek(0, 0)
	m = readFileDisk(f, err)
	fmt.Println(m)
	var posExt int = 0

	for pos, i := range m.Partition {
		if i.PartName == auxName {
			aux, flag := GetNameDisk(path)
			if !flag {
				log.Fatal("OCURRIO UN ERROR AL OBTENER EL NOMBRE DEL DISCO")
			}
			Data.GetName = aux
			Data.GetSize = i.PartSize
			Data.GetStart = i.PartStart
			Data.GetNamePartition = convertPartitionNameToString(i.PartName)
			if part {
				pathDisk = path
				startDisk = i.PartStart
				fitDisk = i.PartFit
			}
		}
		if i.PartType == 'E' {
			posExt = pos
		}
	}
	if Data.GetSize == 0 && Data.GetStart == 0 {
		f.Seek(m.Partition[posExt].PartStart, 0)
		pos, _ := f.Seek(0, os.SEEK_CUR)
		size := binary.Size(e)
		for size != 0 && pos < (m.Partition[posExt].PartSize+m.Partition[posExt].PartStart) {
			e = readFileEBR(f, err)
			if e.PartNameE == auxName {
				aux, flag := GetNameDisk(path)
				if !flag {
					log.Fatal("OCURRIO UN ERROR AL OBTENER EL NOMBRE DEL DISCO")
				}
				Data.GetName = aux
				Data.GetSize = e.PartSizeE
				Data.GetStart = e.PartStartE
				if part {
					pathDisk = path
					startDisk = e.PartStartE
					fitDisk = e.PartFitE
				}

			}
			if e.PartNextE == -1 {
				break
			} else {
				f.Seek(e.PartNextE, io.SeekStart)
			}

		}
	}

	return Data
}

//MakeFileSystem obtiene los tokens y valores del parser creado xd
func MakeFileSystem(root Node) {
	var id string = ""
	var tipo byte = '2'
	var unit byte = '0'

	for _, i := range root.Children {
		if i.TypeToken == "ID" {
			id = i.Value
			fmt.Println(id)
		} else if i.TypeToken == "TYPE" {
			if strings.EqualFold(i.Value, "fast") {
				tipo = '0'
				fmt.Println(tipo)
			} else if strings.EqualFold(i.Value, "full") {
				tipo = '1'
			}
		} else if i.TypeToken == "UNIT" {
			if strings.EqualFold(i.Value, "k") {
				unit = 'k'
				fmt.Println(unit)
			} else if strings.EqualFold(i.Value, "m") {
				unit = 'm'
			} else if strings.EqualFold(i.Value, "b") {
				unit = 'b'
			}
		}
		for _, j := range i.Children {
			if j.TypeToken == "ID" {
				id = j.Value
			} else if j.TypeToken == "TYPE" {
				if strings.EqualFold(j.Value, "fast") {
					tipo = '0'
					fmt.Println(tipo)
				} else if strings.EqualFold(j.Value, "full") {
					tipo = '1'
				}
			} else if j.TypeToken == "UNIT" {
				if strings.EqualFold(j.Value, "k") {
					unit = 'k'
				} else if strings.EqualFold(j.Value, "m") {
					unit = 'm'
				} else if strings.EqualFold(j.Value, "b") {
					unit = 'b'
				}
			}
		}
	}

	disk, err := lista.GetMountedPart(id)
	if err == true {
		getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

		if getData.GetSize != 0 && getData.GetStart != 0 {
			MakeFormatFast(getData.GetStart, getData.GetSize, disk.GetPath(), getData.GetName)
		}

	} else if err == false {
		fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
	}

}

//MakeFormatFast Realiza el formateo fast xd
func MakeFormatFast(start int64, sizePartition int64, path string, name string) {

	//                  INICIALIZACION

	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sb structs_lwh.SB
	nStructures := MakeSizeStructures(sizePartition)
	fmt.Println(nStructures)

	sb = PushDataSB(nStructures, start, sizePartition, name, 1)

	var avd structs_lwh.AVD
	var dd structs_lwh.DDirectory
	var inodo structs_lwh.INodo
	var block structs_lwh.Block
	var bitsAvd []byte
	var bitsDD []byte
	var bitsInodo []byte
	var bitsBlocks []byte

	/*********************************************AVD*******************************************************************/
	//***************BITMAP DE AVD DONDE ESCRIBIMOS SEGUN EL FORMATO 00000000000000000000 etc xD**************
	bitsAvd = createBitmapAVD(f, err, sb)

	/****************************************Directorio de Detalles*********************************/

	bitsDD = createBitmapDD(f, err, sb)

	/***********************************************Tabla de INODOS***********************************/

	bitsInodo = createBitmapInodo(f, err, sb)

	/*******************************BLOCK*********************************************************/

	bitsBlocks = createBitmapBlocks(f, err, sb)

	//Empezamos a Llenar nuestras estructuras xD
	/*----------------------------------------------------------------------------------------------------------------*/
	avd = createAVD("/", 0, 1, 1, 664)
	WriteAVD(0, f, err, sb, avd)

	/*-------------------------------------------------------------------------------------------------------------------------*/

	dd = createDDirectory()

	dd.DDArrayBlock[0] = createArrayFile("users.txt", 0)

	writeDD(0, f, err, sb, dd)

	/*--------------------------------------------------------------------------------------------------------------------------------*/

	inodo = createTableInodo(1, 0, 1, 1, 777)

	inodo.IArrayBlocks[0] = 0

	inodo.IArrayBlocks[1] = 1

	writeTInodo(0, f, err, sb, inodo)

	/*----------------------------------------------------------------------------------------------------------------------------------*/

	block = createBlock("1,G,root\n1,U,root,root,")

	writeBlock(0, f, err, sb, block)

	block = createBlock("201801364\n")

	writeBlock(1, f, err, sb, block)

	/*---------------------------------------------------------------------------------------------------------------------------------------*/
	bitsAvd[0] = '1'
	bitsDD[0] = '1'
	bitsInodo[0] = '1'
	bitsBlocks[0] = '1'
	bitsBlocks[1] = '1'
	/*----------------------------------------------------------------------------------------------------------------------------------------*/
	sb.SbTreeVirtualFree = sb.SbTreeVirtualFree - 1
	sb.SbDetailDirectoryFree = sb.SbDetailDirectoryFree - 1
	sb.SbInodosFree = sb.SbInodosFree - 1
	sb.SbBlocksFree = sb.SbBlocksFree - 1
	sb.SbBlocksFree = sb.SbBlocksFree - 1

	sb = WriteOneInBitmap(f, sb, bitsAvd, bitsDD, bitsInodo, bitsBlocks)
	/*---------------------------------------------------------------------------------------------------------------------------------------*/

	writeSB(f, err, sb, start)

	fmt.Println("SE TERMINO EL FORMATEO :D")
}

//GetNameDisk Obtiene el nombre del disco
func GetNameDisk(path string) (string, bool) {
	index := strings.LastIndex(path, "/")
	if index > -1 {
		aux1 := path[index:]
		aux1 = strings.ReplaceAll(aux1, "/", "")
		index = strings.LastIndex(aux1, ".")
		if index > -1 {
			aux1 = aux1[:index]
			return aux1, true
		}
	}
	return "", false
}

//MakeSizeStructures calcula el size de la formula nStructures
func MakeSizeStructures(sizePartition int64) int64 {
	var sb structs_lwh.SB
	var avd structs_lwh.AVD
	var dd structs_lwh.DDirectory
	var inodo structs_lwh.INodo
	var block structs_lwh.Block
	var log structs_lwh.Log
	sizeSB := int64(binary.Size(sb))
	sizeAVD := int64(binary.Size(avd))
	sizeDD := int64(binary.Size(dd))
	sizeInodo := int64(binary.Size(inodo))
	sizeBlock := int64(binary.Size(block))
	sizeLog := int64(binary.Size(log))
	numerator := sizePartition - (2 * sizeSB)
	var denominator int64 = (27 + sizeAVD + sizeDD + (5*sizeInodo + (20 * sizeBlock) + sizeLog))
	nStructs := numerator / denominator
	return nStructs
}

//PushDataSB inicializa el super block xD
func PushDataSB(nStructs int64, startPartition int64, sizePartition int64, name string, n int64) structs_lwh.SB {
	var sb structs_lwh.SB
	var avd structs_lwh.AVD
	var dd structs_lwh.DDirectory
	var inodo structs_lwh.INodo
	var block structs_lwh.Block
	sizeSB := int64(binary.Size(sb))
	sizeAVD := int64(binary.Size(avd))
	sizeDD := int64(binary.Size(dd))
	sizeInodo := int64(binary.Size(inodo))
	sizeBlock := int64(binary.Size(block))

	copy(sb.SbNameHd[:], name)
	sb.SbTreeVirtualCount = nStructs
	sb.SbDetailDirectoryCount = nStructs
	sb.SbInodosCount = 5 * nStructs
	sb.SbBlocksCount = 20 * nStructs
	sb.SbTreeVirtualFree = nStructs
	sb.SbDetailDirectoryFree = nStructs
	sb.SbInodosFree = (5 * nStructs)
	sb.SbBlocksFree = (20 * nStructs)
	sb.SbDateCreation = createDateNull()
	sb.SbDateLastMount = createDateNull()
	sb.SbMontajesCount = n
	sb.SbApBitmapTreeDirectory = startPartition + sizeSB
	sb.SbApTreeDirectory = sb.SbApBitmapTreeDirectory + sb.SbTreeVirtualCount
	sb.SbApBitmapDetailDirectory = sb.SbApTreeDirectory + (sizeAVD * sb.SbTreeVirtualCount)
	sb.SbApDetailDirectory = sb.SbApBitmapDetailDirectory + sb.SbDetailDirectoryCount
	sb.SbApBitmapTableInodo = sb.SbApDetailDirectory + (sizeDD * sb.SbDetailDirectoryCount)
	sb.SbApTableInodo = sb.SbApBitmapTableInodo + sb.SbInodosCount
	sb.SbApBitmapBlocks = sb.SbApTableInodo + (sizeInodo * sb.SbInodosCount)
	sb.SbApBlocks = sb.SbApBitmapBlocks + sb.SbBlocksCount
	sb.SbApLog = sb.SbApBlocks + (sizeBlock * sb.SbBlocksCount)
	sb.SbSizeStrucTreeDirectory = sizeAVD
	sb.SbSizeStrucDetailDirectory = sizeDD
	sb.SbSizeStrucInodo = sizeInodo
	sb.SbSizeStrucBloque = sizeBlock
	sb.SbFirstFreeBitTreeDirectory = 0
	sb.SbFirstFreeBitDetailDirectory = 0
	sb.SbFirstFreeBitTableInodo = 0
	sb.SbFirstFreeBitBlocks = 0
	sb.SbMagicNum = 201801364
	return sb
}

func createDateNull() [25]byte {
	var DateCreation [25]byte
	var mytime time.Time
	mytime = time.Now()
	var auxTime string = ""
	auxTime = mytime.Format("01-02-2006 15:04:00")
	copy(DateCreation[:], auxTime)
	return DateCreation
}

func createDateNull2() [25]byte {
	var DateCreation [25]byte
	var mytime time.Time
	var auxTime string = ""
	auxTime = mytime.Format("01-02-2006 15:04:00")
	copy(DateCreation[:], auxTime)
	return DateCreation
}

func readFileSB(f *os.File, err error, start int64) structs_lwh.SB {
	f.Seek(start, io.SeekStart)
	var sb structs_lwh.SB

	sizeRead := binary.Size(sb)

	data := readNextBytes(f, sizeRead)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &sb)

	if err != nil {
		log.Fatal("binary.Read failled", err)
	}

	return sb
}
func createBitmapAVD(f *os.File, err error, sb structs_lwh.SB) []byte {
	var i int64 = 0

	var data byte = '0'

	test := make([]byte, sb.SbTreeVirtualCount)

	for ; i < sb.SbTreeVirtualCount; i++ {
		f.Seek(sb.SbApBitmapTreeDirectory+i, io.SeekStart)

		var binaryBitmapDirectory bytes.Buffer

		binary.Write(&binaryBitmapDirectory, binary.BigEndian, &data)

		writeNextBytes(f, binaryBitmapDirectory.Bytes())

		test[i] = data
	}

	return test
}

func createBitmapDD(f *os.File, err error, sb structs_lwh.SB) []byte {
	var i int64 = 0

	var data byte = '0'

	test := make([]byte, sb.SbDetailDirectoryCount)

	for ; i < sb.SbDetailDirectoryCount; i++ {

		f.Seek(sb.SbApBitmapDetailDirectory+i, io.SeekStart)

		var binaryBitmapDD bytes.Buffer

		binary.Write(&binaryBitmapDD, binary.BigEndian, &data)

		writeNextBytes(f, binaryBitmapDD.Bytes())

		test[i] = data
	}

	return test
}

func createBitmapInodo(f *os.File, err error, sb structs_lwh.SB) []byte {
	var i int64 = 0

	var data byte = '0'

	test := make([]byte, sb.SbInodosCount)

	for ; i < sb.SbInodosCount; i++ {

		f.Seek(sb.SbApBitmapTableInodo+i, io.SeekStart)

		var binaryBitmapInodo bytes.Buffer

		binary.Write(&binaryBitmapInodo, binary.BigEndian, &data)

		writeNextBytes(f, binaryBitmapInodo.Bytes())

		test[i] = data
	}

	return test
}

func createBitmapBlocks(f *os.File, err error, sb structs_lwh.SB) []byte {
	var i int64 = 0

	var data byte = '0'

	test := make([]byte, sb.SbBlocksCount)

	for ; i < sb.SbBlocksCount; i++ {

		f.Seek(sb.SbApBitmapBlocks+i, io.SeekStart)

		var binaryBitmapBlock bytes.Buffer

		binary.Write(&binaryBitmapBlock, binary.BigEndian, &data)

		writeNextBytes(f, binaryBitmapBlock.Bytes())

		test[i] = data
	}

	return test
}

//WriteOneInBitmap ....
func WriteOneInBitmap(f *os.File, sb structs_lwh.SB, avd []byte, dd []byte, inodo []byte, block []byte) structs_lwh.SB {

	if avd != nil {

		ReWriteBitmap(f, sb.SbTreeVirtualCount, sb.SbApBitmapTreeDirectory, avd)

		sb.SbFirstFreeBitTreeDirectory = GetNextBitFree(avd)
	}

	if dd != nil {

		ReWriteBitmap(f, sb.SbDetailDirectoryCount, sb.SbApBitmapDetailDirectory, dd)

		sb.SbFirstFreeBitDetailDirectory = GetNextBitFree(dd)
	}
	if inodo != nil {
		ReWriteBitmap(f, sb.SbInodosCount, sb.SbApBitmapTableInodo, inodo)

		sb.SbFirstFreeBitTableInodo = GetNextBitFree(inodo)
	}

	if block != nil {

		ReWriteBitmap(f, sb.SbBlocksCount, sb.SbApBitmapBlocks, block)

		sb.SbFirstFreeBitBlocks = GetNextBitFree(block)
	}

	return sb
}

//ReWriteBitmap ...
func ReWriteBitmap(f *os.File, delimiter int64, start int64, data []byte) {
	var i int64 = 0
	var data0 byte
	for ; i < delimiter; i++ {

		data0 = data[i]

		f.Seek(start+i, io.SeekStart)

		var binaryBitmapBlock bytes.Buffer

		binary.Write(&binaryBitmapBlock, binary.BigEndian, &data0)

		writeNextBytes(f, binaryBitmapBlock.Bytes())
	}
}

//GetNextBitFree ...
func GetNextBitFree(data []byte) int64 {
	var i int64 = 0

	for ; i < int64(len(data)); i++ {
		if data[i] == '0' {
			return i
		}
	}

	return -1
}

func readBitmapAVD(path string, sb structs_lwh.SB, pos int64) {

}

func createAVD(NameDirectory string, AvdApDetailDirectory int64, AvdProper int64, AvdGid int64, AvdPerm int64) structs_lwh.AVD {
	var avd structs_lwh.AVD

	avd.AvdDateCreation = createDateNull()

	copy(avd.AvdNameDirectory[:], NameDirectory)

	avd.AvdApDetailDirectory = AvdApDetailDirectory

	for i := 0; i < 6; i++ {
		avd.AvdApArraySubdirectories[i] = -1
	}

	avd.AvdApTreeVirtualDirectory = -1

	avd.AvdPerm = AvdPerm

	avd.AvdProper = AvdProper

	avd.AvdGid = AvdGid

	return avd
}

//WriteAVD ...
func WriteAVD(i int64, f *os.File, err error, sb structs_lwh.SB, avd structs_lwh.AVD) {
	f.Seek(sb.SbApTreeDirectory+i*int64(binary.Size(avd)), io.SeekStart)

	var binaryStrucAVD bytes.Buffer

	binary.Write(&binaryStrucAVD, binary.BigEndian, &avd)

	writeNextBytes(f, binaryStrucAVD.Bytes())
}

//ReadAVD ...
func ReadAVD(i int64, f *os.File, err error, SbApTreeDirectory int64) structs_lwh.AVD {

	var avd structs_lwh.AVD

	f.Seek(SbApTreeDirectory+i*int64(binary.Size(avd)), io.SeekStart)

	sizeRead := binary.Size(avd)

	data := readNextBytes(f, sizeRead)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &avd)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return avd
}

func createDDirectory() structs_lwh.DDirectory {
	var dd structs_lwh.DDirectory

	for i := 0; i < 5; i++ {
		dd.DDArrayBlock[i] = createArrayFile("", -1)
	}

	dd.DdApDetailDirectory = -1

	return dd
}

func createArrayFile(name string, apInodo int64) structs_lwh.DDArrayFile {
	var af structs_lwh.DDArrayFile

	copy(af.DdFileName[:], name)

	af.DdFileDateCreation = createDateNull()

	af.DdFileDateModification = createDateNull()

	af.DdFileApInodo = apInodo

	return af
}

func writeDD(i int64, f *os.File, err error, sb structs_lwh.SB, dd structs_lwh.DDirectory) {
	f.Seek(sb.SbApDetailDirectory+i*int64(binary.Size(dd)), io.SeekStart)

	var binaryStrucDD bytes.Buffer

	binary.Write(&binaryStrucDD, binary.BigEndian, &dd)

	writeNextBytes(f, binaryStrucDD.Bytes())
}

//ReadDD ...
func ReadDD(i int64, f *os.File, err error, SbApDetailDirectory int64) structs_lwh.DDirectory {

	var dd structs_lwh.DDirectory

	f.Seek(SbApDetailDirectory+i*int64(binary.Size(dd)), io.SeekStart)

	sizeRead := binary.Size(dd)

	data := readNextBytes(f, sizeRead)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &dd)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return dd
}

func createTableInodo(ICountInodo int64, ISizeArchive int64, IProper int64, IGid int64, IPerm int64) structs_lwh.INodo {
	var inodoTable structs_lwh.INodo

	inodoTable.ICountInodo = ICountInodo

	inodoTable.ISizeArchive = ISizeArchive

	inodoTable.ICountBlocksAsigned = 0

	for i := 0; i < 4; i++ {
		inodoTable.IArrayBlocks[i] = -1
	}

	inodoTable.IApIndirecto = -1

	inodoTable.IProper = IProper

	inodoTable.IGid = IGid

	inodoTable.IPerm = IPerm

	return inodoTable
}

func writeTInodo(i int64, f *os.File, err error, sb structs_lwh.SB, inodo structs_lwh.INodo) {

	f.Seek(sb.SbApTableInodo+i*int64(binary.Size(inodo)), io.SeekStart)

	var binaryStrucInodo bytes.Buffer

	binary.Write(&binaryStrucInodo, binary.BigEndian, &inodo)

	writeNextBytes(f, binaryStrucInodo.Bytes())
}

//ReadTInodo ...
func ReadTInodo(i int64, f *os.File, err error, SbApTableInodo int64) structs_lwh.INodo {

	var inodo structs_lwh.INodo

	f.Seek(SbApTableInodo+i*int64(binary.Size(inodo)), io.SeekStart)

	sizeRead := binary.Size(inodo)

	data := readNextBytes(f, sizeRead)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &inodo)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return inodo
}

func createBlock(data string) structs_lwh.Block {
	var block structs_lwh.Block

	copy(block.DbData[:], data)

	return block
}

func writeBlock(i int64, f *os.File, err error, sb structs_lwh.SB, block structs_lwh.Block) {
	f.Seek(sb.SbApBlocks+i*int64(binary.Size(block)), io.SeekStart)

	var binaryStrucBlock bytes.Buffer

	binary.Write(&binaryStrucBlock, binary.BigEndian, &block)

	writeNextBytes(f, binaryStrucBlock.Bytes())
}

//ReadBlock ...
func ReadBlock(i int64, f *os.File, err error, SbApBlocks int64) structs_lwh.Block {

	var block structs_lwh.Block

	f.Seek(SbApBlocks+i*int64(binary.Size(block)), io.SeekStart)

	sizeRead := binary.Size(block)

	data := readNextBytes(f, sizeRead)

	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &block)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return block
}

func writeSB(f *os.File, err error, sb structs_lwh.SB, start int64) {
	f.Seek(start, io.SeekStart)

	var binario bytes.Buffer

	binary.Write(&binario, binary.BigEndian, &sb)

	writeNextBytes(f, binario.Bytes())
}
