package domain

import (
	"errors"
	"io/ioutil"
)

type Object struct {
	Name         string
	ObjectCount  int
	Path         string
	RelativePath string
	IsFolder     bool

	Children []Object
}

func (o *Object) Read() ([]byte, error) {
	if o.IsFolder {
		return nil, errors.New("object is directory")
	}

	return ioutil.ReadFile(o.Path)
}
