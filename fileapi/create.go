package fileapi

import (
	"errors"
	"io/ioutil"
	"os"
)

func Create(path string, body []byte) error {
	if Exists(path) {
		return errors.New("file already exist")
	}

	return ioutil.WriteFile(path, body, os.ModePerm)
}

func CreateFromFile(source, destination string) error {
	if NotExists(source) {
		return errors.New("source file does not exist")
	}

	if Exists(destination) {
		return errors.New("destination file already exist")
	}

	sourceBody, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return Create(destination, sourceBody)
}
