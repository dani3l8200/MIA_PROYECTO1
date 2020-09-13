package structs_lwh

type Log struct {
	LogTypeOperacion [20]byte
	LogType          [16]byte
	LogName          [16]byte
	LogContent       [16]byte
	LogDate          [25]byte
}
