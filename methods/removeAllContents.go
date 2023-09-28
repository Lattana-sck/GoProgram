package methods

import (
	"os"
	"path/filepath"
)

func removeAllContents(path string) error {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}