package structs_lwh

//User estructura para almacenar los usuarios
type User struct {
	Usr [20]byte
	Pwd [20]byte
	Grp [20]byte
	Gid int64
	UID int64
	ID  [20]byte
}

//Group ...
type Group struct {
	Gid int64
	Grp [20]byte
}

//CheckPermiso ...
type CheckPermiso struct {
	Read  bool
	Write bool
	Exec  bool
}

//Permisos ...
type Permisos struct {
	Permiso [3]CheckPermiso
}
