package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read spectrum.txt file
	content, err := os.ReadFile("C:\\Users\\Haady\\Desktop\\Colors\\spectrum.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert content to string and trim whitespace
	text := strings.TrimSpace(string(content))

	// Check if content is "flag"
	if text == "flag" {
		fmt.Println("Juda yaxshi! Sizning bayrog'ingiz:\nHD{Tungi_vaqtda_ko'rinadigan_yoritishni_ko'rish_mumkinmi?}")
	} else {
		fmt.Println("spectrum.txt faylini qiymatini o'zgartirgandan keyin ishga tushirish kerak!\nFayl o'zgartirilmagan!")
	}
}
