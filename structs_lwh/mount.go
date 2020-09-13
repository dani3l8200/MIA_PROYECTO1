package structs_lwh

// MountDisk estructura donde montará una partición del disco en el sistema.
type MountDisk struct {
	ID   string
	Path string
	Name string
}

type GetMounDisk struct {
	GetSize          int64
	GetStart         int64
	GetName          string
	GetNamePartition string
}

//FMountDisk Inicializa la structura mount
func (mount *MountDisk) FMountDisk(ID string, Path string, Name string) MountDisk {
	return MountDisk{ID: ID, Path: Path, Name: Name}
}

//GetID retorna le id generado automaticamente
func (mount *MountDisk) GetID() string {
	return mount.ID
}

//SetID edita algun id, pero es opcional xd
func (mount *MountDisk) SetID(ID string) {
	mount.ID = ID
}

//GetPath retornal el directorio introducido
func (mount *MountDisk) GetPath() string {
	return mount.Path
}

//SetPath edita algun directorio, si lo fuera necesario
func (mount *MountDisk) SetPath(Path string) {
	mount.Path = Path
}

//GetName retorna el nombre del mount
func (mount *MountDisk) GetName() string {
	return mount.Name
}

//SetName edita algun nombre del mount, si fuera necesario
func (mount *MountDisk) SetName(Name string) {
	mount.Name = Name
}
