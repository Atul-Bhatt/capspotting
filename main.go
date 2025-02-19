package main

import (
	"os"
	"log"
)

func main() {
	b, err := os.ReadFile("caption.txt")

	if err != nil {
		log.Fatal(err)
	}
	
	os.Stdout.Write(b)
}
