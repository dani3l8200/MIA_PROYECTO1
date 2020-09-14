package lwh

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DeleteDisk Funcion para eliminar el disco binario creado con rmdisk
func DeleteDisk(path string) {
	var directory string = ""
	if strings.Contains(path, "\"") {
		directory, _ = SetDirectory(path)
	}
	if directory != "" {
		path = directory
	}
	reader := bufio.NewReader(os.Stdin)
	if _, err := os.Stat(path); err == nil {
		for {
			fmt.Println("SEGURO QUE DESEA ELIMINAR EL DISCO ", path, "? (y/n)")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("Y", text) == 0 || strings.Compare("y", text) == 0 {
				deleteFile(path)
				fmt.Println("SE ELIMINO EL DISCO CORRECTAMENTE")
				Pause()
				break
			} else if strings.Compare("N", text) == 0 || strings.Compare("n", text) == 0 {
				return
			}

		}
	} else if os.IsNotExist(err) {
		Pause()
		fmt.Println(err)
	}

}
func deleteFile(path string) {
	var err = os.Remove(path)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
