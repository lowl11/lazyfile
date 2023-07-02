package fmanager

import (
	"github.com/lowl11/lazyfile/data/interfaces"
	"github.com/lowl11/lazyfile/internal/file_manager"
)

func New(root string) interfaces.IManager {
	return file_manager.New(root)
}
