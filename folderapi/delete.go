package folderapi

import "os"

func Delete(path string, withContent bool) error {
	if NotExists(path) {
		return nil
	}

	if withContent {
		if err := os.RemoveAll(path); err != nil {
			return err
		}

		return nil
	}

	return os.Remove(path)
}
