package folderapi

import "os"

func Exists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return !os.IsNotExist(err)
}

func NotExists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return os.IsNotExist(err)
}
