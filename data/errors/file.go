package errors

import "errors"

var (
	FileAlreadyExist   = errors.New("file already exist")
	FileSourceNotFound = errors.New("source file not found")
	FileNotFound       = errors.New("file not found")
	FileIsFolder       = errors.New("this is folder")
)
