package interfaces

import "github.com/lowl11/lazyfile/data/domain"

type IManager interface {
	IFolder

	ThreadSafe() IManager
	FolderByPath(path string) (IFolder, error)
	FileByPath(path string) (IFile, error)
}

type IFolder interface {
	Name() string
	Path() string

	List() ([]domain.Object, error)

	File(name string) (IFile, error)
	AddFile(name string, content []byte) error
	UpdateFile(name string, content []byte) error
	DeleteFile(name string) error

	Folder(name string) (IFolder, error)
	AddFolder(name string) (IFolder, error)
	DeleteFolder(name string) error
	DeleteFolderForce(name string) error
}

type IFile interface {
	Name() string
	Bytes() []byte
	String() string
	Path() string
	Update(content []byte) error
	Delete() error
	IsDestroyed() bool
}
