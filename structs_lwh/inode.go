package structs_lwh

type INodo struct {
	ICountInodo         int64
	ISizeArchive        int64
	ICountBlocksAsigned int64
	IArrayBlocks        [4]int64
	IApIndirecto        int64
	IProper             int64
	IGid                int64
	IPerm               int64
}

type Block struct {
	DbData [25]byte
}
