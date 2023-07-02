package file_manager

import (
	"github.com/lowl11/lazyfile/data/interfaces"
	"sync"
)

type Manager struct {
	path string

	threadSafe bool
	mutex      sync.Mutex
}

func New(path string) interfaces.IManager {
	return &Manager{
		path: path,
	}
}
