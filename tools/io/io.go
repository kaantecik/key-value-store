package iotools

import (
	"github.com/kaantecik/key-value-store/internal/logging"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetFiles function returns all file in given root.
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

// CreateFolder function creates new folder.
func CreateFolder(path string) {
	logrus.Info(path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		logging.ErrorLogger.Fatal(err)
	}
}

// ReadFileAndProcess function reads file and does process.
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
