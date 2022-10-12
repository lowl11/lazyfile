package folderapi

import (
	"fmt"
	"strings"
)

func PathByWindows(path string) string {
	return replaceAllDashes(path, "\\")
}

func PathByUnix(path string) string {
	return replaceAllDashes(path, "/")
}

func replaceAllDashes(path, delimiter string) string {
	path = strings.ReplaceAll(path, "/", delimiter)
	path = strings.ReplaceAll(path, "\\", delimiter)
	return path
}

func buildObjectPath(path, name string) string {
	return fmt.Sprintf("%s/%s", path, name)
}

func buildMemoryObjectPath(relativePath, objectName string) string {
	return fmt.Sprintf("%s/%s", relativePath, objectName)
}
