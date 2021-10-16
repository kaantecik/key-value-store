package ioTools

import (
	"github.com/kaantecik/key-value-store/internal/logging"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetFiles(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}

	return files
}

func CreateFolder(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}
}

func ReadFileAndProcess(path string, inner func(param []byte)) {
	f, err := os.Open(path)
	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logging.ErrorLogger.Fatal(err)
		}
	}(f)

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}

	inner(byteValue)
}
