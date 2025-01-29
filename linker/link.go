package linker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func SetSym(path string, force bool) error {
	conf, err := readTidyConf(path)
	if err != nil {
		return err
	}

	for _, s := range conf.Home {
		err := link(path, s, xdg.Home, force)
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(xdg.ConfigHome)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(xdg.ConfigHome, 0755)
		if err != nil {
			return err
		}
	}
	for _, s := range conf.Config {
		err := link(path, s, xdg.ConfigHome, force)
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(xdg.BinHome)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(xdg.BinHome, 0755)
		if err != nil {
			return err
		}
	}
	for _, s := range conf.Bin {
		err := link(path, s, xdg.BinHome, force)
		if err != nil {
			return err
		}
	}

	for k, c := range conf.Custom {
		for _, s := range c {

			var destPath string
			if strings.HasPrefix(k, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return err
				}
				destPath = filepath.Join(homeDir, k[1:])
			} else {
				destPath = k
			}

			_, err = os.Stat(destPath)
			if errors.Is(err, os.ErrNotExist) {
				err := os.MkdirAll(destPath, 0755)
				if err != nil {
					return err
				}
			}

			err = link(path, s, destPath, force)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func link(arg, s, dest string, force bool) error {
	absTarget, err := filepath.Abs(filepath.Join(arg, s))
	if err != nil {
		return err
	}
	err = os.Symlink(absTarget, filepath.Join(dest, s))
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			if force {
				err := os.RemoveAll(filepath.Join(dest, s))
				if err != nil {
					return err
				}
				err = os.Symlink(absTarget, filepath.Join(dest, s))
				if err != nil {
					return err
				}
			} else {
				fmt.Println(err)
			}
			return nil
		}

		return err
	}

	return nil
}
