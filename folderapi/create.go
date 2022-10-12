package folderapi

import (
	"github.com/lowl11/lazyfile/filemodels"
	"os"
)

func Copy(objectList []filemodels.Object, destination string) error {
	if NotExists(destination) {
		if err := os.Mkdir(destination, os.ModePerm); err != nil {
			return err
		}
	}

	for _, objectItem := range objectList {
		_ = objectItem
	}
	return nil
}

func Create(path, name string) error {
	return os.Mkdir(path+"/"+name, os.ModePerm)
}
