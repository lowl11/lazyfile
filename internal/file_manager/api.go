package file_manager

import (
	"github.com/lowl11/lazyfile/data/domain"
	"github.com/lowl11/lazyfile/data/errors"
	"github.com/lowl11/lazyfile/data/interfaces"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"github.com/lowl11/lazyfile/internal/path_helper"
	"strings"
)

func (manager *Manager) ThreadSafe() interfaces.IManager {
	manager.threadSafe = true
	return manager
}

func (manager *Manager) FileByPath(path string) (interfaces.IFile, error) {
	manager.lock()
	defer manager.unlock()

	filePath := path_helper.Build(manager.path, path)

	content, err := fileapi.Read(filePath)
	if err != nil {
		return nil, err
	}

	pathArray := strings.Split(path, "/")
	var name string
	if len(pathArray) > 1 {
		name = pathArray[len(pathArray)-1]
	} else {
		name = path
	}

	return domain.NewFile(name, content, filePath), nil
}

func (manager *Manager) AddFileByPath(path string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	filePath := path_helper.Build(manager.path, path)

	if fileapi.Exist(filePath) {
		return errors.FileAlreadyExist
	}

	return fileapi.Create(filePath, content)
}

func (manager *Manager) UpdateFileByPath(path string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Update(path, content)
}

func (manager *Manager) DeleteFileByPath(path string) error {
	manager.lock()
	defer manager.unlock()
	return fileapi.Delete(path_helper.Build(manager.path, path))
}

func (manager *Manager) FolderByPath(path string) (interfaces.IFolder, error) {
	manager.lock()
	defer manager.unlock()

	folderPath := path_helper.Build(manager.path, path)

	if !folderapi.Exist(folderPath) {
		return nil, errors.FolderNotExist
	}

	newFolder := New(folderPath)
	if manager.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}

func (manager *Manager) AddFolderByPath(path, name string) error {
	manager.lock()
	defer manager.unlock()

	folderPath := path_helper.Build(manager.path, path)

	if folderapi.Exist(folderPath) {
		return errors.FolderAlreadyExist
	}

	return folderapi.Create(folderPath, name)
}

func (manager *Manager) DeleteFolderByPath(path string, force bool) error {
	manager.lock()
	defer manager.unlock()

	return folderapi.Delete(path, force)
}

func (manager *Manager) Path() string {
	return manager.path
}

func (manager *Manager) Name() string {
	pathArray := strings.Split(manager.path, "/")
	if len(pathArray) == 1 {
		return manager.path
	}

	return pathArray[len(pathArray)-1]
}

func (manager *Manager) List() ([]domain.Object, error) {
	manager.lock()
	defer manager.unlock()

	return folderapi.Objects(manager.path)
}

func (manager *Manager) File(name string) (interfaces.IFile, error) {
	manager.lock()
	defer manager.unlock()

	content, err := fileapi.Read(path_helper.Build(manager.path, name))
	if err != nil {
		return nil, err
	}

	return domain.NewFile(name, content, manager.path), nil
}

func (manager *Manager) AddFile(name string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Create(path_helper.Build(manager.path, name), content)
}

func (manager *Manager) UpdateFile(name string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Update(path_helper.Build(manager.path, name), content)
}

func (manager *Manager) DeleteFile(name string) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Delete(path_helper.Build(manager.path, name))
}

func (manager *Manager) Folder(name string) (interfaces.IFolder, error) {
	manager.lock()
	defer manager.unlock()

	if !folderapi.Exist(name) {
		return nil, errors.FolderNotExist
	}

	newFolder := New(path_helper.Build(manager.path, name))
	if manager.threadSafe {
		newFolder.ThreadSafe()
	}
	return newFolder, nil
}

func (manager *Manager) AddFolder(name string) (interfaces.IFolder, error) {
	manager.lock()
	defer manager.unlock()

	if err := folderapi.Create(manager.path, name); err != nil {
		return nil, err
	}

	newFolder := New(path_helper.Build(manager.path, name))
	if manager.threadSafe {
		newFolder.ThreadSafe()
	}

	return newFolder, nil
}

func (manager *Manager) DeleteFolder(name string, force bool) error {
	manager.lock()
	defer manager.unlock()

	return folderapi.Delete(path_helper.Build(manager.path, name), force)
}
