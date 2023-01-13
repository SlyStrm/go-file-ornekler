package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dosyaBilgi2, err := os.Stat("main.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dosya ismi :", dosyaBilgi2.Name())
	fmt.Println("dosya boyutu: ", dosyaBilgi2.Size())
	fmt.Println("izinler: ", dosyaBilgi2.Mode())
	fmt.Println("güncelleme tarihi: ", dosyaBilgi2.ModTime())
}
