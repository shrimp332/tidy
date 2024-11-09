package linker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func SetSym(arg string) error {
	conf, err := readTidyConf(arg)
	if err != nil {
		return err
	}

	for _, s := range conf.Home {
		err := link(arg, s, xdg.Home)
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
		err := link(arg, s, xdg.ConfigHome)
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
		err := link(arg, s, xdg.BinHome)
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

			err = link(arg, s, destPath)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func link(arg, s, dest string) error {
	absTarget, err := filepath.Abs(filepath.Join(arg, s))
	if err != nil {
		return err
	}
	err = os.Symlink(absTarget, filepath.Join(dest, s))
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println(err)
			return nil
		}

		return err
	}

	return nil
}
