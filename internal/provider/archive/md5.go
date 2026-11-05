package archive

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Package to get the md5 sum of the file.
func getMD5(path string) (sum string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
