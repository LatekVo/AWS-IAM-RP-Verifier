package utilities

import (
	"log"
	"os"
)

func LoadFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
