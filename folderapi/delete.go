package folderapi

import "os"

func Delete(path string, withContent bool) error {
	if withContent {
		if Exists(path) {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		}
	}

	return os.Remove(path)
}
