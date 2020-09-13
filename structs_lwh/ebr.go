// Package structs_lwh paquete que contiene los strucs
package structs_lwh

// EBR struct para el Extender Boot Record
type EBR struct {
	PartStatusE byte
	PartFitE    byte
	PartStartE  int64
	PartSizeE   int64
	PartNextE   int64
	PartNameE   [16]byte
}
