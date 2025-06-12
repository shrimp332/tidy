package linker

import (
	"errors"
	"os"
	"strings"
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
				if force {
					err := os.RemoveAll(dest)
					if err != nil {
						return err
					}
					err = os.Symlink(source, dest)
					if err != nil {
						return err
					}
				} else {
					continue
				}
			} else if errors.Is(err, os.ErrNotExist) {
				lastSlash := strings.LastIndex(dest, "/")
				dirPath := dest[:lastSlash]

				err := os.MkdirAll(dirPath, 0755)
				if err != nil {
					return err
				}

				err = os.Symlink(source, dest)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
