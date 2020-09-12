package structs_lwh

type SB struct {
	SbNameHd                      [20]byte
	SbTreeVirtualCount            int64
	SbDetailDirectoryCount        int64
	SbInodosCount                 int64
	SbBlocksCount                 int64
	SbTreeVirtualFree             int64
	SbDetailDirectoryFree         int64
	SbInodosFree                  int64
	SbBlocksFree                  int64
	SbDateCreation                [25]byte
	SbDateLastMount               [25]byte
	SbMontajesCount               int64
	SbApBitmapTreeDirectory       int64
	SbApTreeDirectory             int64
	SbApBitmapDetailDirectory     int64
	SbApDetailDirectory           int64
	SbApBitmapTableInodo          int64
	SbApTableInodo                int64
	SbApBitmapBlocks              int64
	SbApBlocks                    int64
	SbApLog                       int64
	SbSizeStrucTreeDirectory      int64
	SbSizeStrucDetailDirectory    int64
	SbSizeStrucInodo              int64
	SbSizeStrucBloque             int64
	SbFirstFreeBitTreeDirectory   int64
	SbFirstFreeBitDetailDirectory int64
	SbFirstFreeBitTableInodo      int64
	SbFirstFreeBitBlocks          int64
	SbMagicNum                    int64
	SbSize                        int64
}

type Pointer struct {
	Name     [20]byte
	PPointer int64
}
