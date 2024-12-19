package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Internet Explorer is running...\nEnter 'quit' to exit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "quit" {
			break
		}
	}
}
