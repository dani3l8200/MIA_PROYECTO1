package structs_lwh

// Partitions structura para el mbr
type Partitions struct {
	PartStatus byte
	PartType   byte
	PartFit    byte
	PartStart  int32
	PartSize   int32
	PartName   [16]byte
}

// MBR estructura para el master boot record
type MBR struct {
	MbrSize          int32
	MbrTime          [25]byte
	MbrDiskSignature int32
	Partition        [4]Partitions
}

// MakeMK es una funcion para el mkdisk
