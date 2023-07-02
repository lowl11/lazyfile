package domain

import (
	"errors"
	"github.com/lowl11/lazyfile/fileapi"
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

	if err := fileapi.Update(file.path+"/"+file.name, content); err != nil {
		return err
	}

	file.content = content
	return nil
}

func (file *File) Delete() error {
	if err := fileapi.Delete(file.path + "/" + file.name); err != nil {
		return err
	}
	file.destroy()

	return nil
}

func (file *File) IsDestroyed() bool {
	return file.isDestroyed
}

func (file *File) destroy() {
	file.name = ""
	file.path = ""
	file.content = nil
	file.isDestroyed = true
}
