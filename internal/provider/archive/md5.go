package archive

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// Package to get the md5 sum of the file.
func getMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
