package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

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
		panic(err)
	}
	defer f.Close()

	f.Seek(0, io.SeekStart)

	m := readFileDisk(f, err)

	diskName, _ := GetNameDisk(disk.GetPath())

	var nameLogica string = ""

	testTime := string(m.MbrTime[:19])

	report += "digraph MBR{\n"
	report += "\tgraph[label=\"REPORT MBR\"];\n"
	report += "\trandir=TB;\n\n"

	report += "\tnode0[shape=plaintext, label=<\n"
	report += "\t\t<table border='0' cellborder='1' cellspacing='0' cellpadding='4'>\n"

	report += "\t\t\t<tr><td colspan='2'>MBR " + diskName + "</td></tr>\n"
	report += "\t\t\t<tr>  <td>Nombre</td>  <td>Valor</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_tamaño</td>  <td>" + strconv.Itoa(int(m.MbrSize)) + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_fecha_creacion</td>  <td>" + testTime + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_disk_signature</td>  <td>" + strconv.Itoa(int(m.MbrDiskSignature)) + "</td>  </tr>\n"

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
			report += "\t\t\t<tr>  <td>part_type_" + x + "</td>  <td>" + string(i.PartType) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_fit_" + x + "</td>  <td>" + string(i.PartFit) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_start_" + x + "</td>  <td>" + strconv.Itoa(int(i.PartStart)) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_size_" + x + "</td>  <td>" + strconv.Itoa(int(i.PartSize)) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_name_" + x + "</td>  <td>" + namePartition + "</td>  </tr>\n"
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
				report += "\t\t<table border='0' cellborder='1'  cellspacing='0' cellpadding='4'>\n"
				report += "\t\t\t<tr><td colspan='2'>EBR_" + strconv.Itoa(x) + "</td></tr>\n"
				report += "\t\t\t<tr>  <td>Nombre</td>  <td>Valor</td>  </tr>"

				for _, k := range e.PartNameE {
					if k != 0 {
						nameLogica += string(k)
					}
				}
				report += "\t\t\t<tr>  <td>part_status_1</td>  <td>" + string(e.PartStatusE) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_fit_1</td>  <td>" + string(e.PartFitE) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_start_1</td>  <td>" + strconv.Itoa(int(e.PartStartE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_size_1</td>  <td>" + strconv.Itoa(int(e.PartSizeE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_next_1</td>  <td>" + strconv.Itoa(int(e.PartNextE)) + "</td>  </tr>"
				report += "\t\t\t<tr>  <td>part_name_1</td>  <td>" + nameLogica + "</td>  </tr>"

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
		panic(err)
	}
	defer f.Close()

	f.Seek(0, io.SeekStart)

	m := readFileDisk(f, err)

	diskName, _ := GetNameDisk(path)

	//var nameLogica string = ""

	report += "digraph DISK{\n"
	report += "\tgraph[label=\"Reporte DISK\"];\n"
	report += "\trankdir=TB;\n\n"

	var total float64 = float64(m.MbrSize)

	var space float64 = 0

	report += "\tnodeO[shape=plaintext, label=<\n"
	report += "\t\t<table border='1' cellborder='1'  cellspacing='0' cellpadding='4'>\n"
	report += "\t\t\t<tr>\n\t\t\t\t<td>MBR " + diskName + "</td>\n"

	for j, i := range m.Partition {
		var sizePartition float64 = float64(i.PartSize)

		if i.PartStart != -1 {
			cant := sizePartition * 100 / total

			space += cant

			if i.PartStatus != '1' {
				if i.PartType == 'E' {
					report += "\t\t\t\t<td>\n\t\t\t\t\t<table border='1' cellspacing='0' cellborder='1'>\n"
					report += "\t\t\t\t\t\t<tr>\n"
					report += "\t\t\t\t\t\t\t<td>EXTENDIDA</td>\n"
					report += "\t\t\t\t\t\t</tr>\n\t\t\t\t\t\t<tr>\n"

					f.Seek(i.PartStart, io.SeekStart)
					e := readFileEBR(f, err)
					pos, _ := f.Seek(0, os.SEEK_CUR)

					if e.PartSizeE != 0 {
						f.Seek(i.PartStart, io.SeekStart)

						for pos < (i.PartStart + i.PartSize) {
							sizePartition = float64(e.PartSizeE)
							cant = sizePartition * 100 / total

							if cant != 0 {
								if e.PartStatusE != '1' {
									report += "\t\t\t\t\t<td>EBR</td>\n"
									report += "\t\t\t\t\t<td>LOGICA<br/>" + FloatToString(cant) + "% del Disco</td>\n"
								} else {
									report += "\t\t\t\t\t<td>LIBRE<br/>" + FloatToString(cant) + "% del Disco</td>\n"
								}
							}

							if e.PartNextE == -1 {
								sizePartition = float64(i.PartStart) + float64(i.PartSize) - float64((e.PartStartE + e.PartStartE))
								cant = sizePartition * 100 / total
								if cant != 0 {
									report += "\t\t\t\t\t<td>LIBRE<br/>" + FloatToString(cant) + "% del Disco</td>\n"
								}
								break
							} else {
								f.Seek(e.PartNextE, 0)
								e = readFileEBR(f, err)
							}

						}

					} else {
						report += "\t\t\t\t\t<td>" + strconv.Itoa(int(cant)) + "% del Disco</td>\n"
					}

					report += "\t\t\t\t</tr>\n\t\t\t</table>\n\t\t\t</td>\n"

					var nextS3 int64 = 0

					if j != 3 {
						nextS3 += m.Partition[j+1].PartStart
					}

					report += checkNextSpace(i.PartStart, i.PartSize, nextS3, j, total)

				} else {
					report += "\t\t\t<td>PRIMARIA <br/> " + FloatToString(cant) + "% del disco</td>\n"
					var nextS3 int64 = 0
					if j != 3 {
						nextS3 += m.Partition[j+1].PartStart
					}

					report += checkNextSpace(i.PartStart, i.PartSize, nextS3, j, total)
				}
			} else {
				report += "<td>LIBRE <br/>" + FloatToString(cant) + "% del Disco</td>\n"
			}
		}
	}

	report += "</tr></table>\n\t>];\n\n"

	report += "\n}\n"

	GenerateDot(path, "ResportMount.dot", report)
}

func checkNextSpace(start int64, size int64, nextStart int64, index int, tot float64) string {
	var report string = ""
	var m structs_lwh.MBR
	if index != 3 {
		var p1 float64 = float64(start + size)
		var p2 float64 = float64(nextStart)

		if p2 != -1 {

			if p2 != p1 {
				var frag float64 = p2 - p1
				var percentage float64 = frag * 100 / tot
				if percentage == 0 {
					report = ""
				} else {
					report = "\t\t\t<td> LIBRE<br/> "
					report += FloatToString(percentage)
					report += "% del disco</td>"
				}
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
				report = "\t\t\t<td> LIBRE<br/> "
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
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
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

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
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
	}
	return -1
}

//GetTypeFile ...
func GetTypeFile(path string) string {

	index := strings.LastIndex(path, ".")
	if index > -1 {
		path = path[index:]
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

func GetOpenFile(path string) string {
	index := strings.LastIndex(path, ".")
	if index > -1 {
		path = path[index:]
		return path
	}
	return "-Tpng"
}

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
