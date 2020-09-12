package lwh

import (
	"MIA-PROYECTO1/datastructure"
	"MIA-PROYECTO1/structs_lwh"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var myusers structs_lwh.User
var listU datastructure.LinkedListU
var listG datastructure.LinkedListG

//MakeLogin ....
func MakeLogin(Root Node) {
	var usr string = ""
	var pwd string = ""
	var id string = ""

	for _, i := range Root.Children {
		if i.TypeToken == "USER" {
			usr = i.Value
		} else if i.TypeToken == "PWD" {
			pwd = i.Value
		} else if i.TypeToken == "ID" {
			id = i.Value
		}
		for _, j := range i.Children {
			if j.TypeToken == "USER" {
				usr = j.Value
			} else if j.TypeToken == "PWD" {
				pwd = j.Value
			} else if j.TypeToken == "ID" {
				id = j.Value
			}
		}
	}

	disk, err := lista.GetMountedPart(id)
	if err {
		getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

		if getData.GetSize != 0 && getData.GetStart != 0 {
			Login(disk.GetPath(), usr, pwd, id, getData.GetStart)
		}

	} else if err == false {
		fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
	}
}

// Login Hace el inicio de sesion xd
func Login(path string, usr string, pwd string, id string, start int64) bool {
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Seek(start, io.SeekStart)
	sb := readFileSB(f, err)
	inodo := ReadTInodo(0, f, err, sb.SbApTableInodo)
	us := ReadUsers(f, err, sb, inodo)

	GetRecoveryUsers(us, id)

	if listU.Head != nil {

		for node := listU.Head; node != nil; node = node.Next() {
			xa := converNameSBToByte(usr)
			xs := converNameSBToByte(pwd)
			if node.Value().Usr == xa && node.Value().Pwd == xs {

				myusers = node.Value()

				return true

			} else if node.Value().Usr != xa && node.Value().Pwd != xs {
				fmt.Println("REVISE LOS DATOS INTRODUCIDOS")
				return false
			}
		}
	} else {
		fmt.Println("NO HAY USUARIOS EN EL SISTEMA O INTRODUJO MAL LOS DATOS")
	}

	return false
}

//MakeLogout ....
func MakeLogout(Root Node) {
	var flagCheckLogin = false

	if Root.TypeToken == "LOGOUT" {
		flagCheckLogin = true
		Logout(flagCheckLogin)
	}
}

//Logout ...
func Logout(check bool) {
	auxName := converByteLToString(myusers.Usr)
	pasWord := converByteLToString(myusers.Pwd)
	if check && auxName != "" && pasWord != "" {
		myusers = WriteUser("", "", "", 0, 0, "")
		return
	} else if check == false && auxName == "" && pasWord == "" {
		fmt.Println("NO HAY USUARIOS LOGUEADOS :(")
		return
	}
}

//ReadUsers obtiene todos los usuarios ingresados en formato byte xd
func ReadUsers(f *os.File, err error, sb structs_lwh.SB, inodo structs_lwh.INodo) []byte {

	test := make([]byte, sb.SbBlocksCount*25)
	var content string = ""
	for i := 0; i < 4; i++ {
		if inodo.IArrayBlocks[i] != -1 {
			block := ReadBlock(inodo.IArrayBlocks[i], f, err, sb.SbApBlocks)
			auxName := convertBByteToString(block.DbData)
			content += auxName
		}
	}

	if inodo.IApIndirecto != -1 {
		as := ReadUsers(f, err, sb, ReadTInodo(inodo.IApIndirecto, f, err, sb.SbApTableInodo))
		for _, i := range as {
			if i != 0 {
				content += string(i)
			}
		}
	}
	copy(test[:], content)
	return test
}

//Permisos Crea los permisos que usara el usuario
func Permisos(perm int64) structs_lwh.Permisos {
	userPerm := perm / 100

	groupPerm := (perm - userPerm*100) / 10

	otherPerm := perm - userPerm*100

	var permisos structs_lwh.Permisos
	//Permisos Usuarios
	//------------------------------------------------------------------------------------------
	if userPerm%2 == 1 {
		permisos.Permiso[0].Exec = true
	} else if userPerm%2 == 0 {
		permisos.Permiso[0].Exec = false
	}

	if (userPerm/2)%2 == 1 {
		permisos.Permiso[0].Write = true
	} else if (userPerm/2)%2 == 0 {
		permisos.Permiso[0].Write = false
	}

	if ((userPerm/2)/2)%2 == 1 {
		permisos.Permiso[0].Read = true
	} else if ((userPerm/2)/2)%2 == 0 {
		permisos.Permiso[0].Read = false
	}
	//-----------------------------------------------------------------------------------------------

	if groupPerm%2 == 1 {
		permisos.Permiso[1].Exec = true
	} else if groupPerm%2 == 0 {
		permisos.Permiso[1].Exec = false
	}

	if (groupPerm/2)%2 == 1 {
		permisos.Permiso[1].Write = true
	} else if (groupPerm/2)%2 == 0 {
		permisos.Permiso[1].Write = false
	}

	if ((groupPerm/2)/2)%2 == 1 {
		permisos.Permiso[1].Read = true
	} else if ((groupPerm/2)/2)%2 == 0 {
		permisos.Permiso[1].Read = false
	}
	//------------------------------------------------------------------------------------------------

	if otherPerm%2 == 1 {
		permisos.Permiso[2].Exec = true
	} else if otherPerm%2 == 0 {
		permisos.Permiso[2].Exec = true
	}

	if (otherPerm/2)%2 == 1 {
		permisos.Permiso[2].Write = true
	} else if (otherPerm/2)%2 == 0 {
		permisos.Permiso[2].Write = false
	}

	if ((otherPerm/2)/2)%2 == 1 {
		permisos.Permiso[2].Read = true
	} else if ((otherPerm/2)/2)%2 == 0 {
		permisos.Permiso[2].Read = false
	}
	return permisos
}

//ReadUserByID  Obtiene el usuario ingresando su id
func ReadUserByID(id int64) structs_lwh.User {
	empty := WriteUser("", "", "", -1, -1, "")

	if listU.Head != nil {

		for node := listU.Head; node != nil; node = node.Next() {

			if node.Value().Gid == id {

				return node.Value()

			}

		}

	}

	return empty
}

//ReadUserByName Obtiene el usuario ingresando su usuario
func ReadUserByName(name string) structs_lwh.User {
	auxName := converNameSBToByte(name)

	empty := WriteUser("", "", "", -1, -1, "")

	if listU.Head != nil {

		for node := listU.Head; node != nil; node = node.Next() {

			if node.Value().Usr == auxName {

				return node.Value()

			}

		}
	}

	return empty
}

//WriteUser crea usuarios
func WriteUser(Usr string, Pwd string, Grp string, gid int64, uid int64, id string) structs_lwh.User {
	var user structs_lwh.User
	copy(user.Grp[:], Grp)
	copy(user.Usr[:], Usr)
	copy(user.Pwd[:], Pwd)
	copy(user.ID[:], id)
	user.UID = uid
	user.Gid = gid
	return user
}

//WriteGP ...
func WriteGP(name string, id int64) structs_lwh.Group {
	var group structs_lwh.Group
	group.Gid = id
	copy(group.Grp[:], name)
	return group
}

//GetPRead obtiene los permisos de lectura
func GetPRead(proper int64, idUser int64, permiso structs_lwh.Permisos) bool {

	if idUser == 1 {
		return true
	}

	if proper == idUser {
		return permiso.Permiso[0].Read
	} else if ReadUserByID(idUser).Gid == ReadUserByID(proper).Gid {
		return permiso.Permiso[1].Read
	} else {
		return permiso.Permiso[2].Read
	}
}

//GetPWrite obtiene los permisos de escritura
func GetPWrite(proper int64, idUser int64, permiso structs_lwh.Permisos) bool {
	if idUser == 1 {
		return true
	}

	if proper == idUser {
		return permiso.Permiso[0].Write
	} else if ReadUserByID(idUser).Gid == ReadUserByID(idUser).Gid {
		return permiso.Permiso[1].Write
	} else {
		return permiso.Permiso[2].Write
	}
}

//GetCurrentUser ...
func GetCurrentUser() structs_lwh.User {
	return myusers
}

//GetRecoveryUsers ...
func GetRecoveryUsers(us []byte, id string) {
	//----------------------------------------------------------------------------------------------------------------------------------------------------
	listU.Delete()
	listG.Delete()
	var users string = ""
	for _, name := range us {
		if name != 0 {
			users += string(name)
		}
	}
	//--------------------------------------------------------------------------------------------------------------------------------------------------------
	x := strings.Split(users, "\n")

	x = remove(x, "")

	aux := make([]string, len(x))

	users = ""
	for i := 0; i < len(x); i++ {

		xd := strings.ReplaceAll(x[i], " ", "")

		aux[i] = xd
	}

	for i := 0; i < len(aux); i++ {

		x = strings.Split(aux[i], ",")

		if x[1] == "G" {

			k, _ := strconv.Atoi(x[0])

			listG.Insert(WriteGP(x[2], int64(k)))

		} else if x[1] == "U" {

			if listG.Head != nil {

				xs := converNameSBToByte(x[2])

				for node := listG.Head; node != nil; node = node.Next() {

					if node.Value().Grp == xs {

						k, _ := strconv.Atoi(x[0])

						listU.Insert(WriteUser(x[2], x[4], x[3], node.Value().Gid, int64(k), id))

					}
				}
			}
		}
	}

	//-----------------------------------------------------------------------------------------------------------------------------------------------------------

}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

//converByteLToString .....
func converByteLToString(name [20]byte) string {
	var auxName string
	for _, j := range name {
		auxName += string(j)
	}
	return auxName
}
