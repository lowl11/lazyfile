package errors

import "errors"

var (
	FolderNotExist     = errors.New("folder not exist")
	FolderAlreadyExist = errors.New("folder already exist")
	FolderIsFile       = errors.New("this if file")
)
