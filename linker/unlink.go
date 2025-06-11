package linker

import (
	"errors"
	"os"
)

func UnsetSym(path string) error {
	paths, err := getLinkPaths(path)
	if err != nil {
		return err
	}

	for dest := range paths {
		f, err := os.Lstat(dest)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		if f.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(dest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
