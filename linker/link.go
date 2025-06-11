package linker

import (
	"errors"
	"os"
)

func SetSym(path string, force bool) error {
	paths, err := getLinkPaths(path)
	if err != nil {
		return err
	}

	for dest, source := range paths {
		err := os.Symlink(source, dest)
		if err != nil {
			if errors.Is(err, os.ErrExist) && force {
				err := os.RemoveAll(dest)
				if err != nil {
					return err
				}
				err = os.Symlink(source, dest)
				if err != nil {
					return err
				}
			} else if !force {
				continue
			} else {
				return err
			}
		}
	}
	return nil
}
