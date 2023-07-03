package interfaces

import "github.com/lowl11/lazyfile/data/domain"

type IManager interface {
	IFolder

	/*
		ThreadSafe turns on thread safe mode.
		Created folder objects by IManager inherit thread safe mode
	*/
	ThreadSafe() IManager

	/*
		FileByPath get IFile object by given path.
		Path is path inside given root path
	*/
	FileByPath(path string) (IFile, error)

	/*
		AddFileByPath creates new file by given path.
		Path is path inside given root path
	*/
	AddFileByPath(path string, content []byte) error

	/*
		UpdateFileByPath update file content by given path.
		Path is path inside given root path
	*/
	UpdateFileByPath(path string, content []byte) error

	/*
		DeleteFileByPath removes file by given path.
		Path is path inside given root path
	*/
	DeleteFileByPath(path string) error

	/*
		FolderByPath get IFolder object by give path.
		Path is path inside given root path
	*/
	FolderByPath(path string) (IFolder, error)

	/*
		AddFolderByPath	creates path by given path.
		Path is path inside given root path
	*/
	AddFolderByPath(path, name string) error

	/*
		DeleteFolderByPath removes folder by given path.
		Path is path inside given root path
	*/
	DeleteFolderByPath(path string, force bool) error
}

type IFolder interface {
	Name() string
	Path() string

	Sync() error
	Restore() error

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

	Sync() error
	Restore() error
	IsDestroyed() bool
}
