package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Simulate some processing
	random := 6 * time.Second
	fmt.Println("Wicked process is running. Writing flag in", random, "seconds...")
	time.Sleep(random)

	// Write the flag to the specified file
	flag := "HD{Yomonlar_uchun_dam_yo'q}\n"
	err := os.WriteFile("C:\\Users\\Haady\\Desktop\\flag.txt", []byte(flag), 0644)
	if err != nil {
		fmt.Println("Error writing flag to file:", err)
		return
	} else {
		fmt.Println("Bayroq C:\\Users\\Haady\\Desktop\\flag.txt faylga yozildi")
		return
	}

}
