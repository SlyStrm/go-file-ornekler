package main

import (
	"fmt"
	"log"
	"os"
)

var (
	dosyaBilgi *os.FileInfo
)

func main() {
	dosyaBilgi, err := os.Stat("main.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dosya ismi :", dosyaBilgi.Name())
	fmt.Println("dosya boyutu: ", dosyaBilgi.Size())
	fmt.Println("izinler: ", dosyaBilgi.Mode())
	fmt.Println("güncelleme tarihi: ", dosyaBilgi.ModTime())
}
