package lwh

import (
	"MIA-PROYECTO1/datastructure"
	"MIA-PROYECTO1/structs_lwh"
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

//MakeReports ...
func MakeReports(Root Node) {
	var path string = ""
	var id string = ""
	var ruta string = ""
	var name string = ""

	for _, i := range Root.Children {

		if i.TypeToken == "PATH" {
			path = i.Value
		} else if i.TypeToken == "ID" {
			id = i.Value
		} else if i.TypeToken == "RUTA" {
			ruta = i.Value
			fmt.Println(ruta)
		} else if i.TypeToken == "NAME" {
			name = i.Value
		}

		for _, j := range i.Children {
			if j.TypeToken == "PATH" {
				path = j.Value
			} else if j.TypeToken == "ID" {
				id = j.Value
			} else if j.TypeToken == "RUTA" {
				ruta = j.Value
			} else if j.TypeToken == "NAME" {
				name = j.Value
			}
		}
	}

	if strings.EqualFold(name, "mbr") {
		ReportMBR(path, id)
		return
	} else if strings.EqualFold(name, "disk") {
		ReportDisk(path, id)
		return
	} else if strings.EqualFold(name, "sb") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportSB(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "bm_arbdir") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportBmArbdir(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "bm_detdir") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportBmDetdir(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "bm_inode") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportBmInode(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "bm_block") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportBmBlock(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "directorio") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportDirectory(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "tree_complete") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportTreeComplete(path, disk.GetPath(), getData.GetStart, getData.GetName)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "tree_file") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportTreeFile(path, disk.GetPath(), getData.GetStart, getData.GetName, ruta)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	} else if strings.EqualFold(name, "tree_directorio") {
		disk, err := lista.GetMountedPart(id)
		if err == true {
			getData := GetDiskMount(disk.GetPath(), disk.GetName(), false)

			if getData.GetSize != 0 && getData.GetStart != 0 {
				ReportDirectoryFile(path, disk.GetPath(), getData.GetStart, getData.GetName, ruta)
				return
			}

		} else if err == false {
			fmt.Println("ERROR NO SE ENCONTRO LA PARTICION")
			return
		}
	}
}

//ReportMBR ...
func ReportMBR(path string, id string) {

	var report string = ""
	var namePartition string = ""

	disk, check := lista.GetMountedPart(id)

	if !check {
		fmt.Println("LA PARTICION CON ESTE ID NO ESTA MONTADA, F")
		return
	}

	pathDirectory := disk.GetPath()

	f, err := os.OpenFile(pathDirectory, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.Seek(0, io.SeekStart)

	m := readFileDisk(f, err)

	diskName, _ := GetNameDisk(disk.GetPath())

	var nameLogica string = ""

	testTime := string(m.MbrTime[:19])

	report += "digraph MBR{\n"
	report += "\tgraph[bgcolor = lemonchiffon1,label=\"REPORT MBR\"];\n"
	report += "\trandir=TB;\n\n"

	report += "\tnode0[shape=plaintext, label=<\n"
	report += "\t\t<table border='0' cellborder='1' color=\"green1\" cellspacing='0' cellpadding='4'>\n"

	report += "\t\t\t<tr><td bgcolor=\"brown1\" colspan='2'>MBR " + diskName + "</td></tr>\n"
	report += "\t\t\t<tr>  <td bgcolor=\"azure2\">Nombre</td>  <td bgcolor=\"azure2\">Valor</td>  </tr>\n"
	report += "\t\t\t<tr>  <td bgcolor=\"gold\">mbr_tamaño</td>  <td bgcolor=\"gold\">" + strconv.Itoa(int(m.MbrSize)) + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_fecha_creacion</td>  <td>" + testTime + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td bgcolor=\"gold\">mbr_disk_signature</td>  <td bgcolor=\"gold\">" + strconv.Itoa(int(m.MbrDiskSignature)) + "</td>  </tr>\n"

	indexLogica := -1
	exist := false
	for j, i := range m.Partition {
		auxID := j + 1
		x := strconv.Itoa(auxID)
		if i.PartStart != -1 && i.PartStatus != '1' {
			for _, k := range i.PartName {
				if k != 0 {
					namePartition += string(k)
				}

			}
			report += "\t\t\t<tr>  <td>part_status_" + x + "</td>  <td>" + string(i.PartStatus) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td bgcolor=\"gold\">part_type_" + x + "</td>  <td bgcolor=\"gold\">" + string(i.PartType) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_fit_" + x + "</td>  <td>" + string(i.PartFit) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td bgcolor=\"gold\">part_start_" + x + "</td>  <td bgcolor=\"gold\">" + strconv.Itoa(int(i.PartStart)) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_size_" + x + "</td>  <td>" + strconv.Itoa(int(i.PartSize)) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td bgcolor=\"gold\">part_name_" + x + "</td>  <td bgcolor=\"gold\">" + namePartition + "</td>  </tr>\n"
			namePartition = ""

			if i.PartType == 'E' {
				indexLogica = j
				exist = true
			}
		}
	}

	report += "\t\t</table>\n"
	report += "\t>];\n\n"

	x := 1

	if exist {
		f.Seek(m.Partition[indexLogica].PartStart, io.SeekStart)
		e := readFileEBR(f, err)
		pos, _ := f.Seek(0, os.SEEK_CUR)
		for pos < (m.Partition[indexLogica].PartStart + m.Partition[indexLogica].PartSize) {

			if e.PartStatusE != '1' {
				report += "\tnode" + strconv.Itoa(x) + "[shape=plaintext, label=<\n"
				report += "\t\t<table border='0' cellborder='1' color=\"purple\" cellspacing='0' cellpadding='4'>\n"
				report += "\t\t\t<tr><td colspan='2' bgcolor=\"aquamarine\">EBR_" + strconv.Itoa(x) + "</td></tr>\n"
				report += "\t\t\t<tr>  <td bgcolor=\"SteelBlue1\">Nombre</td>  <td bgcolor=\"SteelBlue1\">Valor</td>  </tr>"

				for _, k := range e.PartNameE {
					if k != 0 {
						nameLogica += string(k)
					}
				}
				report += "\t\t\t<tr>  <td>part_status_1</td>  <td>" + string(e.PartStatusE) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td bgcolor=\"DarkSlateGray1\">part_fit_1</td>  <td bgcolor=\"DarkSlateGray1\">" + string(e.PartFitE) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_start_1</td>  <td>" + strconv.Itoa(int(e.PartStartE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td bgcolor=\"DarkSlateGray1\">part_size_1</td>  <td bgcolor=\"DarkSlateGray1\">" + strconv.Itoa(int(e.PartSizeE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_next_1</td>  <td>" + strconv.Itoa(int(e.PartNextE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td bgcolor=\"DarkSlateGray1\">part_name_1</td>  <td bgcolor=\"DarkSlateGray1\">" + nameLogica + "</td>  </tr>"

				report += "\t\t</table>\n"
				report += "\t>];\n\n"

				nameLogica = ""
			}
			if e.PartNextE == -1 {
				break
			} else {
				f.Seek(e.PartNextE, 0)
				e = readFileEBR(f, err)
			}
			x++

		}
	}

	report += "\n}\n"
	GenerateDot(path, "mbrReport.dot", report)
}

//ReportDisk ...
func ReportDisk(path string, id string) {
	var report string = ""
	//var namePartition string = ""

	disk, check := lista.GetMountedPart(id)

	if !check {
		fmt.Println("LA PARTICION CON ESTE ID NO ESTA MONTADA, F")
		return
	}

	pathDirectory := disk.GetPath()

	f, err := os.OpenFile(pathDirectory, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.Seek(0, io.SeekStart)

	m := readFileDisk(f, err)

	diskName, _ := GetNameDisk(disk.GetPath())

	//var nameLogica string = ""

	report += "digraph DISK{\n"
	report += "\tgraph[bgcolor = lemonchiffon1,label=\"Reporte DISK\"];\n"
	report += "\trankdir=TB;\n\n"

	var total float64 = float64(m.MbrSize)

	var space float64 = 0

	report += "\tnodeO[shape=plaintext, label=<\n"
	report += "\t\t<table border='1' cellborder='1'  cellspacing='0' cellpadding='4'>\n"
	report += "\t\t\t<tr>\n\t\t\t\t<td bgcolor=\"tan3\">MBR " + diskName + "</td>\n"

	for j, i := range m.Partition {
		var sizePartition float64 = float64(i.PartSize)

		if i.PartStart != -1 {
			cant := sizePartition * 100 / total

			auxTotal := sizePartition

			if i.PartStatus != '1' {
				if i.PartType == 'E' {
					report += "\t\t\t\t<td>\n\t\t\t\t\t<table border='1' cellspacing='0' cellborder='1'>\n"
					report += "\t\t\t\t\t\t<tr>\n"
					report += "\t\t\t\t\t\t\t<td bgcolor=\"violet\">EXTENDIDA " + FloatToString(cant) + "% del Disco </td>\n"
					report += "\t\t\t\t\t\t</tr>\n\t\t\t\t\t\t<tr>\n"

					f.Seek(i.PartStart, io.SeekStart)
					e := readFileEBR(f, err)
					pos, _ := f.Seek(0, os.SEEK_CUR)

					if e.PartSizeE != 0 {
						f.Seek(i.PartStart, io.SeekStart)

						for pos < (i.PartStart + i.PartSize) {

							sizePartition = float64(e.PartSizeE)
							cant = sizePartition * 100 / auxTotal
							space += cant
							if cant != 0 {
								if e.PartStatusE != '1' {
									report += "\t\t\t\t\t<td bgcolor=\"LightSeaGreen\">EBR</td>\n"
									report += "\t\t\t\t\t<td bgcolor=\"HotPink2\">LOGICA <br/>" + FloatToString(cant) + "% del SubDisco</td>\n"
								} else {
									report += "\t\t\t\t\t<td bgcolor=\"gold\">LIBRE<br/>" + FloatToString(cant) + "% del SubDisco</td>\n"
								}
							}

							if e.PartNextE == -1 {
								sizePartition = 100 - space
								cant = sizePartition
								if cant != 0 {
									report += "\t\t\t\t\t<td bgcolor=\"gold\">LIBRE<br/>" + FloatToString(cant) + "% del SubDisco</td>\n"
								}
								break
							} else {
								f.Seek(e.PartNextE, 0)
								e = readFileEBR(f, err)
							}

						}

					} else {
						report += "\t\t\t\t\t<td bgcolor=\"gold\">" + strconv.Itoa(int(cant)) + "% del Disco</td>\n"
					}

					report += "\t\t\t\t</tr>\n\t\t\t</table>\n\t\t\t</td>\n"

					var nextS3 int64 = 0

					if j != 3 {
						nextS3 += m.Partition[j+1].PartStart
					}

					report += checkNextSpace(i.PartStart, i.PartSize, nextS3, j, total)

				} else {
					report += "\t\t\t<td bgcolor=\"LawnGreen\">PRIMARIA<br/> " + FloatToString(cant) + "% del disco</td>\n"
					var nextS3 int64 = 0
					if j != 3 {
						nextS3 += m.Partition[j+1].PartStart
					}

					report += checkNextSpace(i.PartStart, i.PartSize, nextS3, j, total)
				}
			} else {
				report += "<td bgcolor=\"gold\">LIBRE <br/>" + FloatToString(cant) + "% del Disco</td>\n"
			}
		}
	}

	report += "</tr></table>\n\t>];\n\n"

	report += "\n}\n"

	GenerateDot(path, "ResportMount.dot", report)
}

//ReportSB ...
func ReportSB(path string, diskPath string, start int64, name string) {

	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sb := readFileSB(f, err, start)

	report += "digraph D{\n"
	report += "graph [pad=\"0.5\", nodesep=\"0.5\", ranksep=\"2\"];\n"
	report += "node [shape=plain]\n"
	report += "rankdir=LR;\n"
	report += "arset [label=<\n"
	report += "<table border=\"0\" cellborder=\"1\" color=\"green1\" cellspacing=\"0\">\n"
	report += "<tr> <td colspan=\"2\" bgcolor=\"green1\"> Reporte SB Ubicado en el disco " + name + " </td> </tr>\n"
	report += "<tr> <td bgcolor=\"green1\"> DATO EN LA ESTRUCTURA </td><td  bgcolor=\"green1\"> DESCRIPCIÓN </td> </tr>"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_nombre_hd: </td><td bgcolor=\"LightSalmon1\">" + converByteLToString(sb.SbNameHd) + "</td></tr>\n"
	report += "<tr> <td> sb_arbol_virtual_count: </td><td>" + strconv.Itoa(int(sb.SbTreeVirtualCount)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_detalle_directory_count: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbDetailDirectoryCount)) + "</td></tr>\n"
	report += "<tr> <td> sb_inodos_count: </td><td>" + strconv.Itoa(int(sb.SbInodosCount)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_bloques_count: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbBlocksCount)) + "</td></tr>\n"
	report += "<tr> <td> sb_arbol_virtual_free: </td><td>" + strconv.Itoa(int(sb.SbTreeVirtualFree)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_detalle_directory_free: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbDetailDirectoryFree)) + "</td></tr>\n"
	report += "<tr> <td> sb_inodos_free: </td><td>" + strconv.Itoa(int(sb.SbInodosFree)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_bloques_free: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbBlocksFree)) + "</td></tr>\n"
	report += "<tr> <td> sb_date_creacion: </td><td>" + string(sb.SbDateCreation[:19]) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_date_ultimo_montaje: </td><td bgcolor=\"LightSalmon1\">" + string(sb.SbDateLastMount[:19]) + "</td></tr>\n"
	report += "<tr> <td> sb_montaje_count: </td><td>" + strconv.Itoa(int(sb.SbMontajesCount)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_ap_bitmap_arbol_directorio: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbApBitmapTreeDirectory)) + "</td></tr>\n"
	report += "<tr> <td> sb_ap_arbol_directory: </td><td>" + strconv.Itoa(int(sb.SbApTreeDirectory)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_ap_bitmap_detalle_directory: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbApBitmapDetailDirectory)) + "</td></tr>\n"
	report += "<tr> <td> sb_ap_detalle_directory: </td><td>" + strconv.Itoa(int(sb.SbApDetailDirectory)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_ap_bitmap_table_inodo: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbApBitmapTableInodo)) + "</td></tr>\n"
	report += "<tr> <td> sb_ap_table_inodo: </td><td>" + strconv.Itoa(int(sb.SbApTableInodo)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_ap_bitmap_bloques: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbApBitmapBlocks)) + "</td></tr>\n"
	report += "<tr> <td> sb_ap_bloques: </td><td>" + strconv.Itoa(int(sb.SbApBlocks)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_ap_log: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbApLog)) + "</td></tr>\n"
	report += "<tr> <td> size_struct_arbol_directorio: </td><td>" + strconv.Itoa(int(sb.SbSizeStrucTreeDirectory)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> size_struct_detalle_directorio: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbSizeStrucDetailDirectory)) + "</td></tr>\n"
	report += "<tr> <td> size_struct_inodo: </td><td>" + strconv.Itoa(int(sb.SbSizeStrucInodo)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> size_struct_bloque: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbSizeStrucBloque)) + "</td></tr>\n"
	report += "<tr> <td> sb_first_free_bit_arbol_directorio: </td><td>" + strconv.Itoa(int(sb.SbFirstFreeBitTreeDirectory)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_first_free_bit_detalle_directorio: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbFirstFreeBitDetailDirectory)) + "</td></tr>\n"
	report += "<tr> <td> sb_first_free_bit_tabla_inodo: </td><td>" + strconv.Itoa(int(sb.SbFirstFreeBitTableInodo)) + "</td></tr>\n"
	report += "<tr> <td bgcolor=\"LightSalmon1\"> sb_first_free_bit_bloques: </td><td bgcolor=\"LightSalmon1\">" + strconv.Itoa(int(sb.SbFirstFreeBitBlocks)) + "</td></tr>\n"
	report += "<tr> <td> sb_magic_num: </td><td>" + strconv.Itoa(int(sb.SbMagicNum)) + "</td></tr>\n"
	report += "</table>\n"
	report += ">]\n}"
	GenerateDot(path, "Report3.dot", report)
}

//ReportBmArbdir ...
func ReportBmArbdir(path string, diskPath string, start int64, name string) {

	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var i int64 = 0
	var j int = 0

	sb := readFileSB(f, err, start)

	bitMapAvd := getSBBitMap(f, err, sb.SbTreeVirtualCount, sb.SbApBitmapTreeDirectory)
	report += "----------------------------REPORTE BITMAP ARBOL DE DIRECTORIO-------------------------\n"
	report += "╔══════════════════════════════════════════════════════════╗\n"

	for i <= 20 && j < len(bitMapAvd) {
		if i < 20 {
			report += "║" + string(bitMapAvd[j]) + "║"
		} else if (i) == 20 {
			i = -1
			report += "\n"
			j = j - 1
		}
		i++
		j++
	}

	report += "\n╚══════════════════════════════════════════════════════════╝"

	GenerateText(path, report)

}

//ReportBmDetdir ...
func ReportBmDetdir(path string, diskPath string, start int64, name string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var i int64 = 0
	var j int = 0

	sb := readFileSB(f, err, start)

	bitMapDD := getSBBitMap(f, err, sb.SbDetailDirectoryCount, sb.SbApBitmapDetailDirectory)
	report += "----------------------------REPORTE BITMAP DETALLE DE DIRECTORIO-------------------------\n"
	report += "╔══════════════════════════════════════════════════════════╗\n"

	for i <= 20 && j < len(bitMapDD) {
		if i < 20 {
			report += "║" + string(bitMapDD[j]) + "║"
		} else if (i) == 20 {
			i = -1
			report += "\n"
			j = j - 1
		}
		i++
		j++
	}

	report += "\n╚══════════════════════════════════════════════════════════╝"

	GenerateText(path, report)
}

//ReportBmInode ...
func ReportBmInode(path string, diskPath string, start int64, name string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var i int64 = 0
	var j int = 0

	sb := readFileSB(f, err, start)

	bitMapInode := getSBBitMap(f, err, sb.SbInodosCount, sb.SbApBitmapTableInodo)
	report += "----------------------------REPORTE BITMAP TABLA DE INODOS-------------------------\n"
	report += "╔══════════════════════════════════════════════════════════╗\n"

	for i <= 20 && j < len(bitMapInode) {
		if i < 20 {
			report += "║" + string(bitMapInode[j]) + "║"
		} else if (i) == 20 {
			i = -1
			report += "\n"
			j = j - 1
		}
		i++
		j++
	}

	report += "\n╚══════════════════════════════════════════════════════════╝"

	GenerateText(path, report)
}

//ReportBmBlock ...
func ReportBmBlock(path string, diskPath string, start int64, name string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var i int64 = 0
	var j int = 0

	sb := readFileSB(f, err, start)

	bitMapBlock := getSBBitMap(f, err, sb.SbBlocksCount, sb.SbApBitmapBlocks)
	report += "----------------------------REPORTE BITMAP BLOQUES-------------------------\n"
	report += "╔══════════════════════════════════════════════════════════╗\n"
	for i <= 20 && j < len(bitMapBlock) {
		if i < 20 {
			report += "║" + string(bitMapBlock[j]) + "║"
		} else if (i) == 20 {
			i = -1
			report += "\n"
			j = j - 1
		}
		i++
		j++
	}

	report += "\n╚══════════════════════════════════════════════════════════╝"

	GenerateText(path, report)
}

//ReportDirectory ...
func ReportDirectory(path string, diskPath string, start int64, name string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sb := readFileSB(f, err, start)

	avd := ReadAVD(0, f, err, sb.SbApTreeDirectory)

	report += "digraph structs {\n"
	report += "\tgraph[label=\"Reporte Directorio\"];\n"
	report += "splines = ortho\n"
	report += GenerateDotAVD(f, err, avd, sb.SbApTreeDirectory, 0, "LightGoldenrod1")
	report += "}\n"

	GenerateDot(path, "reporDirectory.dot", report)
}

//GenerateDotAVD ...
func GenerateDotAVD(f *os.File, err error, avd structs_lwh.AVD, startAvd int64, pos int64, graphColor string) string {
	var report string = ""

	report += "\tAVD" + strconv.Itoa(int(pos)) + " [\n\t"
	report += "\t\tshape = none;\n"
	report += "\t\tlabel = <\n"
	report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"
	report += "\t\t\t\t<tr><td colspan=\"8\" bgcolor=\"" + graphColor + "\" >" + converByteLToString(avd.AvdNameDirectory) + "</td></tr>\n"
	report += "\t\t\t\t<tr>\n"

	if avd.AvdApArraySubdirectories[0] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[0])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[0] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[1] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[1])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[1] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[2] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[2])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[2] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[3] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[3])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[3] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[4] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[4])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[4] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[5] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[5])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[5] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApTreeVirtualDirectory != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"cyan3\">" + strconv.Itoa(int(avd.AvdApTreeVirtualDirectory)) + "</td>\n"
	} else if avd.AvdApTreeVirtualDirectory == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"cyan3\"> &nbsp; </td>\n"
	}

	if avd.AvdApDetailDirectory != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"greenyellow\">" + strconv.Itoa(int(avd.AvdApDetailDirectory)) + "</td>\n"
	} else if avd.AvdApDetailDirectory == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"greenyellow\"> &nbsp; </td>\n"
	}

	report += "\t\t\t\t\t</tr>\n"
	report += "\t\t\t\t</table>\n"
	report += "\t\t\t>\n"
	report += "    ];\n\n"

	for i := 0; i < 6; i++ {
		if avd.AvdApArraySubdirectories[i] != -1 {
			report += "AVD" + strconv.Itoa(int(pos)) + "->AVD" + strconv.Itoa(int(avd.AvdApArraySubdirectories[i])) + ";\n"

			noIndirecto := ReadAVD(avd.AvdApArraySubdirectories[i], f, err, startAvd)

			report += GenerateDotAVD(f, err, noIndirecto, startAvd, avd.AvdApArraySubdirectories[i], "bisque1")
		}
	}

	if avd.AvdApTreeVirtualDirectory != -1 {
		report += "AVD" + strconv.Itoa(int(pos)) + "->AVD" + strconv.Itoa(int(avd.AvdApTreeVirtualDirectory)) + ";\n"

		indirecto := ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, startAvd)

		report += GenerateDotAVD(f, err, indirecto, startAvd, avd.AvdApTreeVirtualDirectory, "yellow")
	}

	return report
}

//ReportTreeComplete ...
func ReportTreeComplete(path string, diskPath string, start int64, name string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sb := readFileSB(f, err, start)

	avd := ReadAVD(0, f, err, sb.SbApTreeDirectory)

	report += "digraph TreeComplete {\n"

	report += "\tgraph[label=\"Reporte Arbol Completo\"];\n"

	report += "splines = ortho\n"

	report += GenerateDotFullTree(f, err, avd, sb.SbApTreeDirectory, 0, "LightGoldenrod1", sb.SbApDetailDirectory, sb.SbApTableInodo, sb.SbApBlocks)

	report += "}\n"

	GenerateDot(path, "reportTreeComplete.dot", report)
}

//GenerateDotFullTree ...
func GenerateDotFullTree(f *os.File, err error, avd structs_lwh.AVD, startAvd int64, pos int64, graphColor string, starDD int64, startInodo int64, startBlock int64) string {
	var report string = ""

	report += "\tAVD" + strconv.Itoa(int(pos)) + " [\n\t"
	report += "\t\tshape = none;\n"
	report += "\t\tlabel = <\n"
	report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"
	report += "\t\t\t\t<tr><td colspan=\"8\" bgcolor=\"" + graphColor + "\" >" + converByteLToString(avd.AvdNameDirectory) + "</td></tr>\n"
	report += "\t\t\t\t<tr>\n"

	if avd.AvdApArraySubdirectories[0] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[0])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[0] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[1] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[1])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[1] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[2] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[2])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[2] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[3] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[3])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[3] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[4] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[4])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[4] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApArraySubdirectories[5] != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"OrangeRed1\">" + strconv.Itoa(int(avd.AvdApArraySubdirectories[5])) + "</td>\n"
	} else if avd.AvdApArraySubdirectories[5] == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
	}

	if avd.AvdApTreeVirtualDirectory != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"cyan3\">" + strconv.Itoa(int(avd.AvdApTreeVirtualDirectory)) + "</td>\n"
	} else if avd.AvdApTreeVirtualDirectory == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"cyan3\"> &nbsp; </td>\n"
	}

	if avd.AvdApDetailDirectory != -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"greenyellow\">" + strconv.Itoa(int(avd.AvdApDetailDirectory)) + "</td>\n"
	} else if avd.AvdApDetailDirectory == -1 {
		report += "\t\t\t\t\t\t<td bgcolor = \"greenyellow\"> &nbsp; </td>\n"
	}

	report += "\t\t\t\t\t</tr>\n"
	report += "\t\t\t\t</table>\n"
	report += "\t\t\t>\n"
	report += "    ];\n\n"

	for i := 0; i < 6; i++ {
		if avd.AvdApArraySubdirectories[i] != -1 {
			report += "AVD" + strconv.Itoa(int(pos)) + "->AVD" + strconv.Itoa(int(avd.AvdApArraySubdirectories[i])) + ";\n"

			noIndirecto := ReadAVD(avd.AvdApArraySubdirectories[i], f, err, startAvd)

			report += GenerateDotFullTree(f, err, noIndirecto, startAvd, avd.AvdApArraySubdirectories[i], "bisque1", starDD, startInodo, startBlock)
		}
	}

	if avd.AvdApTreeVirtualDirectory != -1 {
		report += "AVD" + strconv.Itoa(int(pos)) + "->AVD" + strconv.Itoa(int(avd.AvdApTreeVirtualDirectory)) + ";\n"

		indirecto := ReadAVD(avd.AvdApTreeVirtualDirectory, f, err, startAvd)

		report += GenerateDotFullTree(f, err, indirecto, startAvd, avd.AvdApTreeVirtualDirectory, "yellow", starDD, startInodo, startBlock)
	}

	if avd.AvdApDetailDirectory != -1 {
		report += "AVD" + strconv.Itoa(int(pos)) + "->DD" + strconv.Itoa(int(avd.AvdApDetailDirectory)) + ";\n"

		report += GenerateDotDD(f, err, ReadDD(avd.AvdApDetailDirectory, f, err, starDD), starDD, avd.AvdApDetailDirectory, startInodo, startBlock)
	}

	return report
}

//GenerateDotDD ...
func GenerateDotDD(f *os.File, err error, dd structs_lwh.DDirectory, startDd int64, pos int64, startInode int64, startBlock int64) string {
	var report string = ""
	var graphColor string = ""

	report += "\tDD" + strconv.Itoa(int(pos)) + " [\n\t"
	report += "\t\tshape = none;\n"
	report += "\t\tlabel = <\n"
	report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"
	report += "\t\t\t\t<tr><td colspan=\"2\" bgcolor=\"dodgerblue\" >Directory" + strconv.Itoa(int(pos)) + "</td></tr>\n"

	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			graphColor = "deepskyblue"
		} else if i%2 == 0 {
			graphColor = "lightskyblue1"
		}

		if dd.DDArrayBlock[i].DdFileApInodo != -1 {
			report += "\t\t\t\t<tr>\n"
			report += "\t\t\t\t\t<td bgcolor=\"" + graphColor + "\">" + convertBByteToString(dd.DDArrayBlock[i].DdFileName) + "</td>\n"
			report += "\t\t\t\t\t<td bgcolor=\"" + graphColor + "\">" + strconv.Itoa(int(dd.DDArrayBlock[i].DdFileApInodo)) + "</td>\n"
			report += "\t\t\t\t</tr>\n"
		} else if dd.DDArrayBlock[i].DdFileApInodo == -1 {
			report += "<tr><td bgcolor=\"" + graphColor + "\"> </td><td bgcolor=\"" + graphColor + "\"> </td></tr>\n"
		}
	}

	if dd.DdApDetailDirectory != -1 {
		report += "\t\t\t\t<tr>\n"
		report += "\t\t\t\t\t<td colspan=\"2\" bgcolor=\"greenyellow\">" + strconv.Itoa(int(dd.DdApDetailDirectory)) + "</td>\n"
		report += "\t\t\t\t<tr>\n"
	} else if dd.DdApDetailDirectory == -1 {
		report += "<tr><td colspan=\"2\" bgcolor=\"greenyellow\"> </td></tr>\n"
	}

	report += "\t\t\t\t\t</table>\n"
	report += "\t\t\t\t>\n"
	report += "\t\t\t];\n\n"

	for i := 0; i < 5; i++ {
		if dd.DDArrayBlock[i].DdFileApInodo != -1 {
			inodo := ReadTInodo(dd.DDArrayBlock[i].DdFileApInodo, f, err, startInode)

			report += "DD" + strconv.Itoa(int(pos)) + "->INODO" + strconv.Itoa(int(dd.DDArrayBlock[i].DdFileApInodo)) + ";\n"

			report += GenerateDotInode(f, err, inodo, startInode, dd.DDArrayBlock[i].DdFileApInodo, startBlock)
		}
	}

	if dd.DdApDetailDirectory != -1 {
		report += "DD" + strconv.Itoa(int(pos)) + "->DD" + strconv.Itoa(int(dd.DdApDetailDirectory)) + ";\n"
		report += GenerateDotDD(f, err, ReadDD(dd.DdApDetailDirectory, f, err, startDd), startDd, dd.DdApDetailDirectory, startInode, startBlock)
	}

	return report
}

//GenerateDotInode ...
func GenerateDotInode(f *os.File, err error, inodo structs_lwh.INodo, startInodo int64, pos int64, startBlock int64) string {
	var report string = ""

	report += "\n"

	report += "\tINODO" + strconv.Itoa(int(pos)) + " [\n"

	report += "\t\tshape = none;\n"

	report += "\t\tlabel = <\n"

	report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"

	report += "\t\t\t\t<tr><td bgcolor=\"dodgerblue\" >Inodo" + strconv.Itoa(int(pos)) + "</td></tr>\n"

	for i := 0; i < 4; i++ {
		if inodo.IArrayBlocks[i] != -1 {
			report += "<tr><td>" + strconv.Itoa(int(inodo.IArrayBlocks[i])) + "</td></tr>\n"
		} else if inodo.IArrayBlocks[i] == -1 {
			report += "<tr><td> </td></tr>\n"
		}
	}

	if inodo.IApIndirecto != -1 {
		report += "<tr><td>" + strconv.Itoa(int(inodo.IApIndirecto)) + "</td></tr>\n"
	} else if inodo.IApIndirecto == -1 {
		report += "<tr><td> </td></tr>\n"
	}

	report += "\t\t\t\t</table>\n"

	report += "\t\t\t>\n"

	report += "\t\t];\n\n"

	for i := 0; i < 4; i++ {
		if inodo.IArrayBlocks[i] != -1 {

			report += "\t\tINODO" + strconv.Itoa(int(pos)) + "->BLOCK" + strconv.Itoa(int(inodo.IArrayBlocks[i])) + ";\n"

			block := ReadBlock(inodo.IArrayBlocks[i], f, err, startBlock)

			report += "BLOCK" + strconv.Itoa(int(inodo.IArrayBlocks[i])) + "[shape=\"box\" label=\"" + convertBByteToString(block.DbData) + "\"]\n"
		}
	}

	if inodo.IApIndirecto != -1 {

		report += "INODO" + strconv.Itoa(int(pos)) + "->INODO" + strconv.Itoa(int(inodo.IApIndirecto)) + ";\n"

		report += GenerateDotInode(f, err, ReadTInodo(inodo.IApIndirecto, f, err, startInodo), startInodo, inodo.IApIndirecto, startBlock)
	}

	return report
}

//GenerateDotTreeDd ...
func GenerateDotTreeDd(f *os.File, err error, dd structs_lwh.DDirectory, pos int64, startDd int64) string {
	var graphColor string = ""

	var report string = ""

	report += "\tDD" + strconv.Itoa(int(pos)) + " [\n"
	report += "\t\tshape = none;\n"
	report += "\t\tlabel = <\n"
	report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"
	report += "\t\t\t\t<tr><td colspan=\"2\" bgcolor=\"dodgerblue\" >Detalle " + strconv.Itoa(int(pos)) + "</td></tr>\n"

	for i := 0; i < 5; i++ {
		if (i % 2) == 1 {
			graphColor = "deepskyblue"
		} else if (i % 2) == 0 {
			graphColor = "lightskyblue1"
		}

		if dd.DDArrayBlock[i].DdFileApInodo != -1 {
			report += "\t\t\t\t\t<tr>\n"
			report += "\t\t\t\t\t\t<td bgcolor=\"" + graphColor + "\">" + convertBByteToString(dd.DDArrayBlock[i].DdFileName) + "</td>\n"
			report += "\t\t\t\t\t\t<td bgcolor=\"" + graphColor + "\">" + strconv.Itoa(int(dd.DDArrayBlock[i].DdFileApInodo)) + "</td>\n"
			report += "\t\t\t\t\t</tr>\n"
		} else if dd.DDArrayBlock[i].DdFileApInodo == -1 {
			report += "<tr><td bgcolor=\"" + graphColor + "\"> </td><td bgcolor=\"" + graphColor + "\"> </td></tr>\n"
		}
	}

	if dd.DdApDetailDirectory != -1 {
		report += "\t\t\t\t\t<tr>\n"
		report += "\t\t\t\t\t\t<td colspan=\"2\" bgcolor=\"greenyellow\">" + strconv.Itoa(int(dd.DdApDetailDirectory)) + "</td>\n"
		report += "\t\t\t\t\t</tr>\n"
	} else if dd.DdApDetailDirectory == -1 {
		report += "<tr><td colspan=\"2\" bgcolor=\"greenyellow\"> </td></tr>\n"
	}

	report += "\t\t\t\t</table>\n"
	report += "\t\t\t>\n"
	report += "\t\t];\n\n"

	if dd.DdApDetailDirectory != -1 {
		report += "DD" + strconv.Itoa(int(pos)) + "->DD" + strconv.Itoa(int(dd.DdApDetailDirectory)) + ";\n"
		report += GenerateDotTreeDd(f, err, ReadDD(dd.DdApDetailDirectory, f, err, startDd), dd.DdApDetailDirectory, startDd)
	}

	return report
}

//ReportTreeFile ...
func ReportTreeFile(path string, diskPath string, start int64, name string, ruta string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sb := readFileSB(f, err, start)

	k, _ := SetDirectory(ruta)

	ruta = k

	auxPath := strings.Split(ruta, "/")

	auxPath = deleteEmpty(auxPath)

	var pos int = 0

	var apavd int64 = 0

	avd := ReadAVD(apavd, f, err, sb.SbApTreeDirectory)

	if ruta == "/" {
		archivos := CreateListPointerDd(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory))

		if archivos.Head != nil {
			i := 0
			for node := archivos.Head; node != nil; node = node.Next() {

				fmt.Println(strconv.Itoa(i+1), convertBByteToString(node.Value().Name))
				i++
			}
		}

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("INGRESE UNA OPCION")

			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)

			newRuta, _ := strconv.Atoi(text)

			if newRuta < archivos.Length()+1 {

				report += "digraph structs {\n"

				report += "splines = ortho\n"

				inodo := ReadTInodo(archivos.Get(newRuta-1), f, err, sb.SbApTableInodo)

				report += GenerateDotInode(f, err, inodo, sb.SbApTableInodo, archivos.Get(newRuta-1), sb.SbApBlocks)

				report += "}\n"

				GenerateDot(path, "ReporFile.dot", report)

				return
			}
			break
		}

	} else if ruta != "/" {
		for pos < len(auxPath) {
			if pos < len(auxPath)-1 {
				if !(existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1) {
					fmt.Println("LA CARPETA NO EXISTE")
					Pause()
				}
				apavd = existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos])

				if len(GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid).AvdNameDirectory) != 0 {
					avd = GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid)
				} else if len(GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid).AvdNameDirectory) == 0 {
					fmt.Println("SIN PERMISOS DE LECTURA")
					Pause()
				}
				pos++
			} else if pos == len(auxPath)-1 {
				if existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos]) != -1 {
					apavd = existDirectoryAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos])

					avd = GetFilesAVD(f, err, sb.SbApTreeDirectory, avd, auxPath[pos], myusers.UID, myusers.Gid)

					archivos := CreateListPointerDd(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory))

					if archivos.Head != nil {
						i := 0
						for node := archivos.Head; node != nil; node = node.Next() {

							fmt.Println(strconv.Itoa(i+1), convertBByteToString(node.Value().Name))
							i++
						}
					}

					reader := bufio.NewReader(os.Stdin)

					for {
						fmt.Println("INGRESE UNA OPCION")

						text, _ := reader.ReadString('\n')

						text = strings.Replace(text, "\n", "", -1)

						newRuta, _ := strconv.Atoi(text)

						if newRuta < archivos.Length()+1 {

							report += "digraph structs {\n"

							report += "splines = ortho\n"

							inodo := ReadTInodo(archivos.Get(newRuta-1), f, err, sb.SbApTableInodo)

							report += GenerateDotInode(f, err, inodo, sb.SbApTableInodo, archivos.Get(newRuta-1), sb.SbApBlocks)

							report += "}\n"

							GenerateDot(path, "ReporFile.dot", report)

							return
						}
						break
					}
				} else if existDetailDirectory(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), auxPath[pos]) {
					archivos := CreateListPointerDd(f, err, sb.SbApDetailDirectory, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory))

					if archivos.Head != nil {
						i := 0
						for node := archivos.Head; node != nil; node = node.Next() {

							fmt.Println(strconv.Itoa(i+1), convertBByteToString(node.Value().Name))
							i++
						}
					}

					reader := bufio.NewReader(os.Stdin)

					for {
						fmt.Println("INGRESE UNA OPCION")

						text, _ := reader.ReadString('\n')

						text = strings.Replace(text, "\n", "", -1)

						newRuta, _ := strconv.Atoi(text)

						if newRuta < archivos.Length()+1 {

							report += "digraph structs {\n"

							report += "splines = ortho\n"

							inodo := ReadTInodo(archivos.Get(newRuta-1), f, err, sb.SbApTableInodo)

							report += GenerateDotInode(f, err, inodo, sb.SbApTableInodo, archivos.Get(newRuta-1), sb.SbApBlocks)

							report += "}\n"

							GenerateDot(path, "ReporFile.dot", report)

							return
						}
						break
					}
				}
				pos++
			}
		}
	}
}

//ReportDirectoryFile ...
func ReportDirectoryFile(path string, diskPath string, start int64, name string, ruta string) {
	var report string = ""

	f, err := os.OpenFile(diskPath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sb := readFileSB(f, err, start)

	k, _ := SetDirectory(ruta)

	ruta = k

	auxPath := strings.Split(ruta, "/")

	auxPath = deleteEmpty(auxPath)

	avd := ReadAVD(0, f, err, sb.SbApTreeDirectory)

	var directorio datastructure.LinkedListP

	directorio.Insert(createPointer("\"/\"", 0))

	var exit bool = true

	for exit != false {
		if directorio.Length() > 0 {
			if directorio.Head != nil {
				i := 0
				for node := directorio.Head; node != nil; node = node.Next() {

					fmt.Println(strconv.Itoa(i+1), convertBByteToString(node.Value().Name))
					i++
				}
			}
			fmt.Println("INGRESE UNA OPCION")
			reader := bufio.NewReader(os.Stdin)

			text, _ := reader.ReadString('\n')

			text = strings.Replace(text, "\n", "", -1)

			newRuta, _ := strconv.Atoi(text)

			if newRuta < directorio.Length()+1 {
				fmt.Println("a. Buscar Siguiente Carpeta")
				fmt.Println("b. Generar Reporte")

				text, _ := reader.ReadString('\n')

				text = strings.Replace(text, "\n", "", -1)

				if strings.EqualFold(text, "a") {
					avd = ReadAVD(directorio.Get(newRuta-1), f, err, sb.SbApTreeDirectory)
					directorio = CreateListPointerAvd(f, err, sb.SbApTreeDirectory, avd)
				} else if strings.EqualFold(text, "b") {
					avd = ReadAVD(directorio.Get(newRuta-1), f, err, sb.SbApTreeDirectory)

					report += "digraph structs {\n"
					report += "splines = ortho\n"

					report += "\tAVD" + strconv.Itoa(int(directorio.Get(newRuta-1))) + " [\n"
					report += "\t\tshape = none;\n"
					report += "\t\tlabel = <\n"
					report += "\t\t\t<table border=\"0\" cellborder=\"2\" cellspacing=\"2\" color=\"cyan4\">\n"
					report += "\t\t\t\t<tr><td colspan=\"8\" bgcolor=\"white\" >" + converByteLToString(avd.AvdNameDirectory) + "</td></tr>\n"
					report += "\t\t\t\t<tr>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"lightcyan\"> &nbsp; </td>\n"
					report += "\t\t\t\t\t<td bgcolor = \"cyan3\"> &nbsp; </td>\n"

					if avd.AvdApDetailDirectory != -1 {
						report += "\t\t\t\t\t<td bgcolor = \"greenyellow\">" + strconv.Itoa(int(avd.AvdApDetailDirectory)) + "</td>\n"
					} else if avd.AvdApDetailDirectory == -1 {
						report += "\t\t\t\t\t<td bgcolor = \"greenyellow\"> &nbsp; </td>\n"
					}

					report += "\t\t\t\t</tr>\n"
					report += "\t\t\t</table>\n"
					report += "\t\t>\n"
					report += "\t];\n\n"

					report += "AVD" + strconv.Itoa(int(directorio.Get(newRuta-1))) + "->DD" + strconv.Itoa(int(avd.AvdApDetailDirectory))

					report += GenerateDotTreeDd(f, err, ReadDD(avd.AvdApDetailDirectory, f, err, sb.SbApDetailDirectory), avd.AvdApDetailDirectory, sb.SbApDetailDirectory)

					report += "}\n"

					GenerateDot(path, "TreeDirectorio.dot", report)

					exit = false
				}
			}
		}
	}

}

func checkNextSpace(start int64, size int64, nextStart int64, index int, tot float64) string {
	var report string = ""
	var m structs_lwh.MBR
	if index != 3 {
		var p1 float64 = float64(start + size)
		var size float64 = tot + float64(int64(binary.Size(m)))

		if nextStart == -1 && size != p1 {

			var freeS float64 = size - p1 + float64(binary.Size(m))

			var percentage float64 = freeS * 100 / tot

			if percentage == 0 {
				report = ""
			} else {
				report = "\t\t\t<td bgcolor=\"gold\"> LIBRE<br/> "
				report += FloatToString(percentage)
				report += "% del disco</td>"
			}

		}
	} else {
		var p1 float64 = float64(start + size)

		var size float64 = tot + float64(int64(binary.Size(m)))

		if size != p1 {
			var freeS float64 = size - p1 + float64(binary.Size(m))

			var percentage float64 = freeS * 100 / tot

			if percentage == 0 {
				report = ""
			} else {
				report = "\t\t\t<td bgcolor=\"gold\"> LIBRE<br/> "
				report += FloatToString(percentage)
				report += "% del disco</td>"
			}
		}
	}
	return report
}

func writeFileReport(path string, content string) {
	// open file using READ & WRITE permission
	createFile(path)
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0777)
	checkError(err)
	defer file.Close()
	// write some text to file
	_, err = file.WriteString(content)
	checkError(err)

	// save changes
	err = file.Sync()
	checkError(err)
}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

//GenerateText ...
func GenerateText(path string, content string) {
	directory := GetNameReport(path)

	makeDirectory(directory)
	aux, _ := SetDirectory(path)

	writeFileReport(aux, content)

	archive, _ := SetDirectory(path)

	openMyReport := GetOpenFile(path)

	x, _ := SetDirectory(openMyReport)

	openMyReport = x

	fmt.Println("-----------REPORTE GENERADO CON EXITO!--------------")

	ViewReport(openMyReport, archive)

	Pause()
}

//GenerateDot ...
func GenerateDot(path string, nameReport string, content string) {

	directory := GetNameReport(path)

	makeDirectory(directory)

	nameReport = directory + nameReport

	writeFileReport(nameReport, content)

	tipo := GetTypeFile(path)

	archive, _ := SetDirectory(path)

	execProcess(tipo, nameReport, archive)

	openMyReport := GetOpenFile(path)

	x, _ := SetDirectory(openMyReport)

	openMyReport = x
	fmt.Println("-----------REPORTE GENERADO CON EXITO!--------------")

	ViewReport(openMyReport, archive)

	Pause()

}

//FloatToString ...
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

func execProcess(tipo string, inputFile string, outputFile string) int {
	app := "dot"

	arg0 := tipo
	arg1 := inputFile
	arg2 := "-o"
	arg3 := outputFile

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err == nil {
		return 0
	}

	// Figure out the exit code
	if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		if ws.Exited() {
			return ws.ExitStatus()
		}

		if ws.Signaled() {
			return -int(ws.Signal())
		}
	}
	return -1
}

//ViewReport ...
func ViewReport(tipo string, outputFile string) int {
	if tipo == ".jpg" || tipo == ".png" {
		app := "eog"
		arg0 := outputFile

		cmd := exec.Command(app, arg0)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err == nil {
			return 0
		}

		// Figure out the exit code
		if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			if ws.Exited() {
				return ws.ExitStatus()
			}

			if ws.Signaled() {
				return -int(ws.Signal())
			}
		}
	} else if tipo == ".pdf" {
		app := "xdg-open"
		arg0 := outputFile

		cmd := exec.Command(app, arg0)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err == nil {
			return 0
		}

		// Figure out the exit code
		if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			if ws.Exited() {
				return ws.ExitStatus()
			}

			if ws.Signaled() {
				return -int(ws.Signal())
			}
		}
	} else if tipo == ".txt" {
		app := "gedit"
		arg0 := outputFile

		cmd := exec.Command(app, arg0)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err == nil {
			return 0
		}

		// Figure out the exit code
		if ws, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			if ws.Exited() {
				return ws.ExitStatus()
			}

			if ws.Signaled() {
				return -int(ws.Signal())
			}
		}
	}
	return -1
}

//GetTypeFile ...
func GetTypeFile(path string) string {

	index := strings.LastIndex(path, ".")
	if index > -1 {
		path = path[index:]
		x, _ := SetDirectory(path)
		path = x
		if path == ".pdf" {
			return "-Tpdf"
		} else if path == ".png" {
			return "-Tpng"
		} else if path == ".jpg" {
			return "-Tjpg"
		}
	}
	return "-Tpng"
}

//GetOpenFile ...
func GetOpenFile(path string) string {
	index := strings.LastIndex(path, ".")
	if index > -1 {
		path = path[index:]
		return path
	}
	return "-Tpng"
}

//GetNameReport ...
func GetNameReport(path string) string {
	if strings.Contains(path, "\"") {
		path = strings.ReplaceAll(path, "\"", "")

		index := strings.LastIndex(path, "/")

		if index > -1 {
			path = path[:index] + "/"
			return path
		}

	} else if !strings.Contains(path, "\"") {
		index := strings.LastIndex(path, "/")

		if index > -1 {
			path = path[:index] + "/"

			return path
		}
	}
	return ""
}
