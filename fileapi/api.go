package fileapi

import (
	"errors"
	"io/ioutil"
	"os"
)

// Create file in given path
func Create(path string, body []byte) error {
	if Exist(path) {
		return errors.New("file already exist")
	}

	return ioutil.WriteFile(path, body, os.ModePerm)
}

/*
	CreateFromFile create file
	Takes content from one file and create new with given path
*/
func CreateFromFile(source, destination string) error {
	if NotExist(source) {
		return errors.New("source file does not exist")
	}

	if Exist(destination) {
		return errors.New("destination file already exist")
	}

	sourceBody, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return Create(destination, sourceBody)
}

// Exist check if file exist
func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}

// NotExist like Exist but opposite
func NotExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return os.IsNotExist(err)
}

// Read get content of file
func Read(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
