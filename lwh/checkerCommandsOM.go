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
		} else if i.TypeToken == "UNIT" {
			unitChekc = true
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
				unitChekc = true
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

// CheckRMdisk metodo utilizado para verificar que venga solo path y de manera obligatoria
func CheckRMdisk(aux Node) bool {
	var pathCheck bool = false
	for _, i := range aux.Children {
		if i.TypeToken == "PATH" {
			if pathCheck {
				fmt.Println("COMANDO REPETIDO PATH :(")
				return false
			}
			pathCheck = true
		}
	}
	var globalFlag bool = true
	if !pathCheck {
		fmt.Println("HACE FALTA EL COMANDO PATH")
		globalFlag = false
	}
	return globalFlag
}

// ChekcFDisk metodo utilizado para verificar que venga los comandos especificados en el enunciado para el comando fdisk
func ChekcFDisk(aux Node) bool {
	var generalFlag bool = true
	var sizeCheck bool = false
	var unitCheck bool = false
	var pathCheck bool = false
	var typeCheck bool = false
	var fitCheck bool = false
	var deleteCheck bool = false
	var nameCheck bool = false
	var addCheck bool = false
	for _, i := range aux.Children {
		if i.TypeToken == "SIZE" {
			sizeCheck = true
			for _, char := range i.Value {
				if char < '0' || char > '9' {
					fmt.Println("error numeros negativos ")
					return false
				}
			}
		} else if i.TypeToken == "UNIT" {
			unitCheck = true
		} else if i.TypeToken == "PATH" {
			pathCheck = true
		} else if i.TypeToken == "TYPE" {
			typeCheck = true
		} else if i.TypeToken == "FIT" {
			fitCheck = true
		} else if i.TypeToken == "DELETE" {
			deleteCheck = true
		} else if i.TypeToken == "NAME" {
			nameCheck = true
		} else if i.TypeToken == "ADD" {
			addCheck = true
		}
		for _, j := range i.Children {
			if j.TypeToken == "SIZE" {
				if sizeCheck {
					fmt.Println("COMANDO SIZE REPETIDO :(")
					return false
				}
				sizeCheck = true
				for _, char := range j.Value {
					if char < '0' || char > '9' {
						fmt.Println("error numeros negativos ")
						return false
					}
				}
			} else if j.TypeToken == "UNIT" {
				if unitCheck {
					fmt.Println("COMANDO UNIT REPETIDO :(")
					return false
				}
				unitCheck = true
			} else if j.TypeToken == "PATH" {
				if pathCheck {
					fmt.Println("COMANDO PATH REPETIDO :( ")
					return false
				}
				pathCheck = true
			} else if j.TypeToken == "TYPE" {
				if typeCheck {
					fmt.Println("COMANDO TYPE REPETIDO :(")
					return false
				}
				typeCheck = true
			} else if j.TypeToken == "FIT" {
				if fitCheck {
					fmt.Println("COMANDO FIT REPETIDO :(")
					return false
				}
				fitCheck = true
			} else if j.TypeToken == "DELETE" {
				if deleteCheck {
					fmt.Println("COMANDO DELETE REPETIDO :(")
					return false
				}
				deleteCheck = true

			} else if j.TypeToken == "NAME" {
				if nameCheck {
					fmt.Println("COMANDO NAME REPETIDO :(")
					return false
				}
				fmt.Println(j.Value)
				nameCheck = true
			} else if j.TypeToken == "ADD" {
				if addCheck {
					fmt.Println("COMANDO ADD REPETIDO :(")
					return false
				}
				addCheck = true
			}
		}
	}

	if !nameCheck {
		fmt.Println("ERROR EN FDISK FALTA EL COMANDO NAME :(")
		generalFlag = false
	}

	if !pathCheck {
		fmt.Println("ERROR EN FDISK FALTA EL COMANDO PATH :(")
		generalFlag = false
	}

	if addCheck && deleteCheck {
		fmt.Println("ERRO EN FDISK NO TIENE SENTIDO USAR ADD Y DELETE :)")
		generalFlag = false
	}

	if !sizeCheck && !addCheck && !deleteCheck {
		fmt.Println("ERROR EN FDISK FALTAN LOS COMANDOS SIZE, ADD, DELETE :(")
		generalFlag = false
	}

	if sizeCheck && addCheck {
		fmt.Println("ERROR EN FDISK NO TIENE SENTIDO SIZE Y ADD :)")
		generalFlag = false
	}

	if sizeCheck && deleteCheck {
		fmt.Println("ERROR EN FDISK NO TIENE SENTIDO SIZE Y DELETE :)")
		generalFlag = false
	}

	fmt.Println(sizeCheck, unitCheck, pathCheck, typeCheck, fitCheck, deleteCheck, nameCheck, addCheck)
	return generalFlag
}
