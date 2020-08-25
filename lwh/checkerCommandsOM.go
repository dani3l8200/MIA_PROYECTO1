package lwh

import (
	"fmt"
	"strings"
)

// CheckMKdisk funcion para verificar las entradas obligatorias
func CheckMKdisk(aux Node) bool {
	var sizeCheck bool = false
	var pathCheck bool = false
	var nameCheck bool = false
	var unitChekc bool = false
	for _, i := range aux.Children {
		if i.TypeToken == "SIZE" {
			sizeCheck = true
			for _, char := range i.Value {
				if char < '0' || char > '9' {
					fmt.Println("error numeros negativos ")
					return false
				}
			}
		} else if i.TypeToken == "PATH" {
			fmt.Println("EXISTE UN PATH")
			pathCheck = true
		} else if i.TypeToken == "NAME" {
			fmt.Println("EXISTE UN NAME")
			nameCheck = true
			if !strings.Contains(i.Value, ".dsk") {
				fmt.Println("Error no tiene la extension .dsk")
				return false
			}
		}
		for _, j := range i.Children {
			if j.TypeToken == "SIZE" {
				if sizeCheck {
					fmt.Println("COMANDO REPETIDO SIZE :( ")
					return false
				}
				sizeCheck = true
				for _, char := range j.Value {
					if char < '0' || char > '9' {
						fmt.Println("error numeros negativos ")
						return false
					}
				}
			} else if j.TypeToken == "PATH" {
				if pathCheck {
					fmt.Println("COMANDO REPETIDO PATH :(")
					return false
				}
				pathCheck = true
			} else if j.TypeToken == "NAME" {
				if nameCheck {
					fmt.Println("COMANDO NAME REPETIDO")
					return false
				}
				nameCheck = true
				if !strings.Contains(j.Value, ".dsk") {
					fmt.Println("Error no tiene la extension .dsk")
					return false
				}
			} else if j.TypeToken == "UNIT" {
				if unitChekc {
					fmt.Println("COMANDO REPETIDO UNIT")
					return false
				}
				unitChekc = false
			}

		}
	}
	var generalFlag bool = true
	if !sizeCheck {
		fmt.Println("HACE FALTA EL COMANDO SIZE")
		generalFlag = false
	}

	if !pathCheck {
		fmt.Println("HACE FALTA EL COMANDO PATH")
		generalFlag = false
	}
	if !nameCheck {
		fmt.Println("HACE FALTA EL COMANDO NAME")
		generalFlag = false
	}

	return generalFlag

}
