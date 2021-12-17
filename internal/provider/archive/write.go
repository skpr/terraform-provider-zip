package archive

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"io"
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
		header := &zip.FileHeader{
			Name:   fileName,
			Method: zip.Deflate,
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return path, err
		}

		var buf bytes.Buffer

		enc := gob.NewEncoder(&buf)

		err = enc.Encode(fileData)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(headerWriter, &buf)
		if err != nil {
			return path, err
		}
	}

	return path, nil
}
