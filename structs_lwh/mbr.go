package structs_lwh

// Partitions structura para el mbr
type Partitions struct {
	PartStatus byte
	PartType   byte
	PartFit    byte
	PartStart  int64
	PartSize   int64
	PartName   [16]byte
}

// MBR estructura para el master boot record
type MBR struct {
	MbrSize          int64
	MbrTime          [25]byte
	MbrDiskSignature int64
	Partition        [4]Partitions
}

// MakeMK es una funcion para el mkdisk
