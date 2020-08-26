// Package structs_lwh paquete que contiene los strucs
package structs_lwh

// EBR struct para el Extender Boot Record
type EBR struct {
	PartStatusE byte
	PartFitE    byte
	PartStartE  int32
	PartSizeE   int32
	PartNextE   int32
	PartNameE   [16]byte
}
