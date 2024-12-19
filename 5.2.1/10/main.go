package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Check if user "Pianist" is member of "Strings" group
	stringsCmd := exec.Command("net", "localgroup", "Strings")
	stringsOut, err := stringsCmd.Output()
	if err != nil {
		fmt.Println("Error checking Strings group:", err)
		return
	}

	// Check if user "Pianist" is member of "Percussion" group 
	percCmd := exec.Command("net", "localgroup", "Percussion")
	percOut, err := percCmd.Output()
	if err != nil {
		fmt.Println("Error checking Percussion group:", err)
		return
	}

	// Convert outputs to string and check if "Pianist" is present in both
	stringsStr := strings.ToLower(string(stringsOut))
	percStr := strings.ToLower(string(percOut))
	pianist := strings.ToLower("Pianist")

	if strings.Contains(stringsStr, pianist) && strings.Contains(percStr, pianist) {
		fmt.Println("Nice")
	} else {
		fmt.Println("Bad")
	}
}
