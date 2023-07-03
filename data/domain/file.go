package domain

import (
	"errors"
	"fmt"
	errors2 "github.com/lowl11/lazyfile/data/errors"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/internal/path_helper"
)

type File struct {
	name        string
	content     []byte
	path        string
	isDestroyed bool
}

func NewFile(name string, content []byte, path string) *File {
	return &File{
		name:    name,
		content: content,
		path:    path,
	}
}

func (file *File) Name() string {
	if file.isDestroyed {
		return ""
	}

	return file.name
}

func (file *File) Bytes() []byte {
	if file.isDestroyed {
		return nil
	}

	return file.content
}

func (file *File) String() string {
	if file.isDestroyed {
		return ""
	}

	return string(file.content)
}

func (file *File) Path() string {
	if file.isDestroyed {
		return ""
	}

	return file.path
}

func (file *File) Update(content []byte) error {
	if file.isDestroyed {
		return errors.New("file destroyed")
	}

	if err := fileapi.Update(path_helper.Build(file.path, file.name), content); err != nil {
		return err
	}

	file.content = content
	return nil
}

func (file *File) Delete() error {
	if err := fileapi.Delete(path_helper.Build(file.path, file.name)); err != nil {
		return err
	}

	file.isDestroyed = true

	return nil
}

func (file *File) IsDestroyed() bool {
	return file.isDestroyed
}

func (file *File) Sync() error {
	fmt.Println("sync:", path_helper.Build(file.path, file.name))
	if !fileapi.Exist(path_helper.Build(file.path, file.name)) {
		return errors2.FileNotFound
	}

	return nil
}

func (file *File) Restore() error {
	return fileapi.Create(path_helper.Build(file.path, file.name), file.content)
}
