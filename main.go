package main

import "MIA-PROYECTO1/analyzers"

func main() {
	analyzers.Execute()
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

/*if x.value == "mkdisk" {
	fmt.Println(x.value)
	for i, s := range x.children {
		fmt.Println(i, s.value)
	}
} else if x.value == "exec" {
	fmt.Println(x.value)
	for k, m := range x.children {
		fmt.Println(k, m.value)
	}


}*/
