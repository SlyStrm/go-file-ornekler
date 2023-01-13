package main

// Dosya içeriği Kopyalama - tümünü çekme
import (
	"io"
	"log"
	"os"
)

func main() {
	orijinal, err := os.Open("xmain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer orijinal.Close()

	Dosya, err := os.Create("x")
	if err != nil {
		log.Fatal(err)
	}
	defer Dosya.Close()

	icerik, err := io.Copy(Dosya, orijinal)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d byte kopyalandı", icerik)
	err = Dosya.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
