package main

import (
	"log"
	"os"
)

func main() {

	original := "main.txt"
	hedef := "./test/test.txt"
	err := os.Rename(original, hedef)
	if err != nil {
		log.Fatal()
	}
}
