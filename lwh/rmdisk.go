package lwh

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DeleteDisk Funcion para eliminar el disco binario creado con rmdisk
func DeleteDisk(path string) {
	reader := bufio.NewReader(os.Stdin)
	if _, err := os.Stat(path); err == nil {
		fmt.Println(path)
		for {
			fmt.Println("SEGURO QUE DESEA ELIMINAR EL DISCO ", path, "? (y/n)")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("Y", text) == 0 || strings.Compare("y", text) == 0 {
				deleteFile(path)
				break
			} else if strings.Compare("N", text) == 0 || strings.Compare("n", text) == 0 {
				break
			}

		}
	} else if os.IsNotExist(err) {
		panic(err)
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
