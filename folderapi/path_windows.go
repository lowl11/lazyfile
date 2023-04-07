package folderapi

/*
	Path give path with Windows format
	Windows path contains reverse slashes - \
*/
func Path(path string) string {
	return replaceAllDashes(path, "\\")
}
