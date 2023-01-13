package main

import (
	"log"
	"os"
)

func main() {
	dosyaBilgisi1, err := os.Stat("x.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("baba! akü yok")
		}
	}
	log.Println("dosya bulundu", dosyaBilgisi1)
	log.Println("dosya bilgileri: ", dosyaBilgisi1)
}
