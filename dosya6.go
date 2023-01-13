package main

// Dosya Silme
import (
	"log"
	"os"
)

func main() {
	err := os.Remove("./x/y")
	if err != nil {
		log.Fatal(err)
	}
}
