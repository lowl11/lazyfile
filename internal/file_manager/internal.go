package file_manager

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
