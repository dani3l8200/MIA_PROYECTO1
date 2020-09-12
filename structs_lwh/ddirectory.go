package structs_lwh

type DDArrayFile struct {
	DdFileName             [25]byte
	DdFileApInodo          int64
	DdFileDateCreation     [25]byte
	DdFileDateModification [25]byte
}

type DDirectory struct {
	DDArrayBlock        [5]DDArrayFile
	DdApDetailDirectory int64
}
