package file_manager

import (
	"github.com/lowl11/lazyfile/data/domain"
	"github.com/lowl11/lazyfile/data/errors"
	"github.com/lowl11/lazyfile/data/interfaces"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"strings"
)

func (manager *Manager) ThreadSafe() interfaces.IManager {
	manager.threadSafe = true
	return manager
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

	content, err := fileapi.Read(manager.path + "/" + name)
	if err != nil {
		return nil, err
	}
	return domain.NewFile(name, content, manager.path), nil
}

func (manager *Manager) AddFile(name string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Create(manager.path+"/"+name, content)
}

func (manager *Manager) UpdateFile(name string, content []byte) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Update(manager.path+"/"+name, content)
}

func (manager *Manager) DeleteFile(name string) error {
	manager.lock()
	defer manager.unlock()

	return fileapi.Delete(manager.path + "/" + name)
}

func (manager *Manager) Folder(name string) (interfaces.IFolder, error) {
	if !folderapi.Exist(name) {
		return nil, errors.FolderNotExist
	}

	return New(manager.path + "/" + name), nil
}

func (manager *Manager) AddFolder(name string) (interfaces.IFolder, error) {
	manager.lock()
	defer manager.unlock()

	if err := folderapi.Create(manager.path, name); err != nil {
		return nil, err
	}

	return New(manager.path + "/" + name), nil
}

func (manager *Manager) DeleteFolder(name string) error {
	manager.lock()
	defer manager.unlock()

	return folderapi.Delete(manager.path+"/"+name, false)
}

func (manager *Manager) DeleteFolderForce(name string) error {
	manager.lock()
	defer manager.unlock()

	return folderapi.Delete(manager.path+"/"+name, true)
}
