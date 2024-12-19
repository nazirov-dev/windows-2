package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("7-Zip File Manager is running...\nEnter 'quit' to exit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "quit" {
			break
		}
	}
}
