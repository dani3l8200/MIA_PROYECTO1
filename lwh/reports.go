package lwh

import (
	"MIA-PROYECTO1/structs_lwh"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func ReportMBR(path string, m structs_lwh.MBR) {

	var report string = ""
	var namePartition string = ""
	size := strconv.Itoa(int(m.MbrSize))
	MbrDiskSignature := strconv.Itoa(int(m.MbrDiskSignature))
	testTime := string(m.MbrTime[:19])
	report += "digraph MBR{\n"
	report += "\tgraph[label=\"REPORT MBR\"];\n"
	report += "\trandir=TB;\n\n"

	report += "\tnode0[shape=plaintext, label=<\n"
	report += "\t\t<table border='0' cellborder='1' cellspacing='0' cellpadding='4'>\n"

	report += "\t\t\t<tr><td colspan='2'>MBR " + "ReportDISK" + "</td></tr>\n"
	report += "\t\t\t<tr>  <td>Nombre</td>  <td>Valor</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_tamaño</td>  <td>" + size + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_fecha_creacion</td>  <td>" + testTime + "</td>  </tr>\n"
	report += "\t\t\t<tr>  <td>mbr_disk_signature</td>  <td>" + MbrDiskSignature + "</td>  </tr>\n"

	for j, i := range m.Partition {
		auxID := j + 1
		x := strconv.Itoa(auxID)
		PartStart := strconv.Itoa(int(i.PartStart))
		PartSize := strconv.Itoa(int(i.PartSize))
		if i.PartStart != -1 && i.PartStatus != '1' {
			for _, k := range i.PartName {
				if k != 0 {
					namePartition += string(k)
				}

			}
			report += "\t\t\t<tr>  <td>part_status_" + x + "</td>  <td>" + string(i.PartStatus) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_type_" + x + "</td>  <td>" + string(i.PartType) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_fit_" + x + "</td>  <td>" + string(i.PartFit) + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_start_" + x + "</td>  <td>" + PartStart + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_size_" + x + "</td>  <td>" + PartSize + "</td>  </tr>\n"
			report += "\t\t\t<tr>  <td>part_name_" + x + "</td>  <td>" + namePartition + "</td>  </tr>\n"
			namePartition = ""
		}
	}

	report += "\t\t</table>\n"
	report += "\t>];\n\n"

	report += "\n}\n"
	writeFileReport(path, report)
	execProcess()
}

func writeFileReport(path string, content string) {
	// open file using READ & WRITE permission
	createFile(path)
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
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

func execProcess() int {
	app := "dot"

	arg0 := "-Tjpg"
	arg1 := "archivo.dot"
	arg2 := "-o"
	arg3 := "reportMBR.jpg"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd.Dir = "/home/dani3l8200/Escritorio/MisDiscos/"
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
