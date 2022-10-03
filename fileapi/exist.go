package fileapi

import "os"

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}

func NotExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return os.IsNotExist(err)
}
