package folderapi

/*
	Path give path with Unix (Linux & macOS) format
	Unix path contains default slashed - /
*/
func Path(path string) string {
	return replaceAllDashes(path, "/")
}
