package interfaces

import "github.com/lowl11/lazyfile/data/domain"

type IManager interface {
	IFolder

	ThreadSafe() IManager

	FileByPath(path string) (IFile, error)
	AddFileByPath(path string, content []byte) error
	UpdateFileByPath(path string, content []byte) error
	DeleteFileByPath(path string) error

	FolderByPath(path string) (IFolder, error)
	AddFolderByPath(path, name string) error
	DeleteFolderByPath(path string, force bool) error
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
	DeleteFolder(name string, force bool) error
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
