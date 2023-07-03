package file_manager

import (
	"github.com/lowl11/lazyfile/data/errors"
	"github.com/lowl11/lazyfile/data/interfaces"
	"github.com/lowl11/lazyfile/folderapi"
	"github.com/lowl11/lazyfile/internal/path_helper"
)

func (manager *Manager) lock() {
	if !manager.threadSafe {
		return
	}

	manager.mutex.Lock()
}

func (manager *Manager) unlock() {
	if !manager.threadSafe {
		return
	}

	manager.mutex.Unlock()
}

func (manager *Manager) getFolder(name string) (interfaces.IFolder, error) {
	folderName := path_helper.Build(manager.path, name)

	if !folderapi.Exist(folderName) {
		return nil, errors.FolderNotExist
	}

	newFolder := New(folderName)
	if manager.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}
