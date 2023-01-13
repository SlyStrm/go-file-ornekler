package main

import (
	"log"
	"os"
)

var (
	yeniDosya1 *os.File
	err        error
)

func main() {
	yeniDosya1, err = os.Create("main.txt")
	if err != nil {
		log.Fatal(err)
	}
}
