package main

import "MIA-PROYECTO1/analyzers"

func main() {
	analyzers.Execute()

	/*	if _, err := os.Stat("/home/Juan Daniel/"); os.IsNotExist(err) {
		os.Mkdir("/home/Juan Daniel/", 0777)
	}*/
	/*	var lista datastructure.LinkedList
		var mount structs_lwh.MountDisk
		name := "JUAN"
		path := "/home/dani33l820"
		letter := lista.SetLetter(path)

		number := lista.SetNumber(path)
		var ids string = ""

		var auxID [16]byte
		copy(auxID[:], "vd ")
		auxID[2] = letter
		for _, v := range auxID {
			if v != 0 {
				ids += string(v)
			}
		}
		ids += strconv.Itoa(number)

		lista.Insert(mount.FMountDisk(ids, path, name))

		letter1 := lista.SetLetter("/home/Escritorio/NAme")

		number1 := lista.SetNumber("/home/Escritorio/NAme")
		ids = ""

		var auxID1 [16]byte

		copy(auxID1[:], "vd ")
		auxID1[2] = letter1
		for _, v := range auxID1 {
			if v != 0 {
				ids += string(v)
			}
		}
		ids += strconv.Itoa(number1)

		fmt.Println(letter1)
		lista.Insert(mount.FMountDisk("A299", "/home/Escritorio/NAme", "Marcela"))
		test, err := lista.GetMountedPart("A2997")
		if !err {
			log.Fatalln("NO SE ENCONTRO EL ID")
		}
		fmt.Println(test.ID, test.Name, test.Path)*/

	//lwh.ReportMBR("/home/dani3l8200/Escritorio/MisDiscos/archivo.dot")

	/*if _, err := os.Stat("/home/dani3l8200/gocode/src/intento/day.go"); err == nil {
		fmt.Println("existe")

	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		log.Fatal(err)

	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Println("F")
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

	}*/

}
