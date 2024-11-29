package linker

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func UnsetSym(arg string) error {
	conf, err := readTidyConf(arg)
	if err != nil {
		return err
	}

	for _, s := range conf.Home {
		if err := unlink(s, xdg.Home); err != nil {
			return err
		}
	}

	for _, s := range conf.Config {
		if err := unlink(s, xdg.ConfigHome); err != nil {
			return err
		}
	}

	for _, s := range conf.Bin {
		if err := unlink(s, xdg.BinHome); err != nil {
			return err
		}
	}

	for k, c := range conf.Custom {
		for _, s := range c {

			var dest string
			if strings.HasPrefix(k, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return err
				}
				dest = filepath.Join(homeDir, k[1:])
			} else {
				dest = k
			}

			if err := unlink(s, dest); err != nil {
				return err
			}
		}
	}

	return nil
}

func unlink(s, dest string) error {
	path := filepath.Join(dest, s)

	f, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if f.Mode()&os.ModeSymlink != 0 {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}
