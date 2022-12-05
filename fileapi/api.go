package fileapi

import (
	"errors"
	"io/ioutil"
	"os"
)

/*
	Create file in given path
	If file already exist does nothing
*/
func Create(path string, body []byte) error {
	if Exist(path) {
		return nil
	}

	return ioutil.WriteFile(path, body, os.ModePerm)
}

/*
	CreateFromFile create file
	Takes content from one file and create new with given path
	If source file does not exist returns error
	If destination path already exist does nothing
*/
func CreateFromFile(source, destination string) error {
	if NotExist(source) {
		return errors.New("source file does not exist")
	}

	if Exist(destination) {
		return nil
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
	stat, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return stat == nil
}

// Read get content of file
func Read(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
