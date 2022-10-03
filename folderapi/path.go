package folderapi

import "fmt"

func buildObjectPath(path, name string) string {
	return fmt.Sprintf("%s/%s", path, name)
}

func buildMemoryObjectPath(relativePath, objectName string) string {
	return fmt.Sprintf("%s/%s", relativePath, objectName)
}
