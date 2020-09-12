package lwh

import (
	"MIA-PROYECTO1/datastructure"
	"MIA-PROYECTO1/structs_lwh"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

var lista datastructure.LinkedList

func MountPartitions(root Node) {
	var name string = ""
	var path string = ""

	for _, i := range root.Children {
		if i.TypeToken == "PATH" {
			path = i.Value
			fmt.Println(path)
		} else if i.TypeToken == "NAME" {
			name = i.Value
			fmt.Println(name)
		}
		for _, j := range i.Children {
			if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "NAME" {
				name = j.Value
			}
		}
	}
	if name != "" && path != "" {

		if _, err := os.Stat(path); err == nil {
			auxName := converNameToByte(name)
			f, err := os.OpenFile(path, os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			f.Seek(0, 0)
			m := readFileDisk(f, err)

			var index = -1
			var CheckExtend = -1
			//Revisamos las particiones del mbr
			for v, i := range m.Partition {
				PartName := i.PartName
				if auxName == PartName {
					index = v
					break
				}
				if i.PartType == 'E' {
					CheckExtend = v
				}
			}
			//Revisamos las particiones del ebr
			if index == -1 && CheckExtend != -1 {
				var ebr structs_lwh.EBR
				f.Seek(m.Partition[CheckExtend].PartStart, 0)
				ebr = readFileEBR(f, err)
				pos, _ := f.Seek(0, os.SEEK_CUR)
				size := binary.Size(ebr)
				for size != 0 && pos < (m.Partition[CheckExtend].PartSize+m.Partition[CheckExtend].PartStart) {
					PartName := ebr.PartNameE
					if PartName == auxName {
						index = CheckExtend
						break
					}
					if ebr.PartNextE == -1 {
						break
					} else {
						f.Seek(ebr.PartNextE, 0)
						ebr = readFileEBR(f, err)
					}
				}
			}
			if index != -1 {
				if !lista.MountedPart(path, name) {
					letter := lista.SetLetter(path)
					number := lista.SetNumber(path)
					var mount structs_lwh.MountDisk
					var id string = ""

					var auxID [16]byte
					copy(auxID[:], "vd ")
					auxID[2] = letter
					for _, v := range auxID {
						if v != 0 {
							id += string(v)
						}
					}
					id += strconv.Itoa(number)
					lista.Insert(mount.FMountDisk(id, path, name))
					lista.Print()
				} else {
					fmt.Println("La particion con l nombre", name, "ya ha sido montara anteriormente")
				}
			} else {
				fmt.Println("LA PARTICION NO HA SIDO ENONTRADO EN EL DISCO")
			}

		} else if os.IsNotExist(err) {
			panic(err)
		}
	} else {
		lista.Print()
	}
}

func UnmountPartitions(root Node) {
	for _, i := range root.Children {
		if i.TypeToken == "ID" {
			lista.DeleteMount(i.Value)
		}
		for _, j := range i.Children {
			lista.DeleteMount(j.Value)
		}
	}
}
