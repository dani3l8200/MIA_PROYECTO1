package structs_lwh

type AVD struct {
	AvdDateCreation           [25]byte
	AvdNameDirectory          [20]byte
	AvdApArraySubdirectories  [6]int64
	AvdApDetailDirectory      int64
	AvdApTreeVirtualDirectory int64
	AvdProper                 int64
	AvdGid                    int64
	AvdPerm                   int64
}
