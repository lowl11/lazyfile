package folderapi

import "os"

func Exists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}

func NotExists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if err != nil {
		return false
	}

	return os.IsNotExist(err)
}
