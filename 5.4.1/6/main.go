package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Check if Calculator.exe is running
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq Calculator.exe")
	output, _ := cmd.Output()
	// If Calculator.exe is not found in tasklist
	if strings.Contains(string(output), "No tasks are running which match the specified criteria.") {
		fmt.Println("Vauv! Sizning bayrog'ingiz:\nHD{Ushbu_jarayon_sizga_nima_qildi?!}\n")
	} else {
		fmt.Println("Xato! Calculator.exe hali ham ishlamoqda!")
	}
}
