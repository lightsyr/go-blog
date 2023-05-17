package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(path string, file *multipart.FileHeader) (string, error) {
	// reads file
	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// create file destiny
	dst, err := os.Create(filepath.Join(dir, "/public/images/", file.Filename))

	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return "", err
	}

	return file.Filename, nil
}
