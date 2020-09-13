package lwh

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MakeMkdir(Root Node) {
	var id string = ""
	var path string = ""
	var p bool = false

	for _, i := range Root.Children {
		if i.TypeToken == "PATH" {
			path = i.Value
		} else if i.TypeToken == "ID" {
			id = i.Value
		} else if i.TypeToken == "P" {
			p = true
		}

		for _, j := range i.Children {
			if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "ID" {
				id = j.Value
			} else if j.TypeToken == "P" {
				p = true
			}
		}
	}

	disk, err := lista.GetMountedPart(id)
	if err {
		getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

		if getData.GetSize != 0 && getData.GetStart != 0 {
			auxPath, errorPath := SetDirectory(path)
			if errorPath {
				path = auxPath

				MakeDirectorys(disk.GetPath(), getData.GetStart, path, p)

			} else if errorPath == false {
				log.Fatal("ERROR CON EL PATH")
			}

		}

	} else if err == false {
		fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
	}
}

func MakeDirectorys(diskPath string, start int64, path string, p bool) {

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
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
				if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					WriteFilesAVD(start, f, err, auxPath[pos], apIndirecto, myusers.UID, myusers.Gid, avd, sb)
					fmt.Println("SE CREO LA CARPETA CON EXITO")
					return
				} else if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					fmt.Println("NO SE CREO LA CARPETA ")
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
				if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					WriteFilesAVD(start, f, err, auxPath[pos], apIndirecto, myusers.UID, myusers.Gid, avd, sb)
					fmt.Println("SE CREO LA CARPETA CON EXITO")
					return
				} else if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					fmt.Println("NO SE CREO LA CARPETA ")
					return
				}
			}

		}
	}
}
