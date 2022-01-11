package archive

import (
	"archive/zip"
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

// Helper function to write the zip file to disk.
func writeZip(files map[string]interface{}) (string, error) {
	// Use hashing on the files map so we can build a deterministic hash for this file.
	name := md5.Sum([]byte(fmt.Sprintf("%v", files)))

	// Store this in the operating systems temporary directory. That way we don't have to worry about cleanup.
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%x.zip", name))

	f, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	for fileName, fileData := range files {
		file, err := writer.Create(fileName)
		if err != nil {
			return path, err
		}

		raw := []byte(fmt.Sprintf("%v", fileData))

		_, err = file.Write(raw)
		if err != nil {
			return path, err
		}
	}

	return path, nil
}
