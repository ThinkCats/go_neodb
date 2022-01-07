package store

import (
	"log"
	"os"
)

func OpenFile(fileName string) *os.File {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func CreateFixedSizeFile(fileName string, mbSize int64) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Truncate(mbSize * 1024 * 1024); err != nil {
		log.Fatal(err)
	}
}
