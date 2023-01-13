package main

// rename refactor
import (
	"log"
	"os"
)

func main() {

	original := "main.txt"
	hedef := "./xtest/xxtest.txt"
	err := os.Rename(original, hedef)
	if err != nil {
		log.Fatal()
	}
}
