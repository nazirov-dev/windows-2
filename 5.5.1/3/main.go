package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("C:\\Users\\Haady\\Desktop\\OSProc\\OSLib.dll")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dastur ishlamaydi, chunki OSLib.dll fayli sizning kompyuterida yo'q. Dasturni qayta o'rnatish yordamida bu muammo yechiladi.")
			os.Exit(1)
		}
		fmt.Printf("Faylni tekshirishda xatolik: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Sizning bayrog'ingiz:\nHD{Sizning_kutubxona_shu_qal'ada!}\n")
}
