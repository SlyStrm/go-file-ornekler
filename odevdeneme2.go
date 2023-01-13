package main

//Dosyadan  bilgi çekme

import (
	"fmt"
	"log"
	"os"
)

var (
	dosyaBilgi *os.FileInfo
)

func main() {
	dosyaBilgi, err := os.Stat("xmain.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dosya ismi :", dosyaBilgi.Name())
	fmt.Println("dosya boyutu: ", dosyaBilgi.Size())
	fmt.Println("izinler: ", dosyaBilgi.Mode())
	fmt.Println("güncelleme tarihi: ", dosyaBilgi.ModTime())
}
