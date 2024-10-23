package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

type TidyConf struct {
	Config []string `json:"config"`
	Home   []string `json:"home"`
	Bin    []string `json:"bin"`
}

func main() {
	var set bool
	var unset bool

	rootCmd := &cobra.Command{
		Use:   "tidy",
		Short: "Tidy Dotfile Linker",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if set {
				for _, arg := range args {
					err = SetSym(arg)
					if errors.Is(err, os.ErrNotExist) {
						fmt.Fprintln(os.Stderr, arg, "does not have a .tidy.json file")
						err = nil
					} else if err != nil {
						return err
					}
				}
			} else if unset {
				for _, arg := range args {
					err = UnsetSym(arg)
					if errors.Is(err, os.ErrNotExist) {
						fmt.Fprintln(os.Stderr, arg, "does not have a .tidy.json file")
						err = nil
					} else if err != nil {
						return err
					}
				}
			} else {
				cmd.Help()
			}
			return err
		},
		Example: "tidy [-s | -u] [directory | *]",
	}

	rootCmd.Flags().
		BoolVarP(&set, "set", "s", false, "use to create symlinks, mutually exclusive with unset")
	rootCmd.Flags().
		BoolVarP(&unset, "unset", "u", false, "use to remove symlinks, mutually exclusive with set")
	rootCmd.MarkFlagsMutuallyExclusive("set", "unset")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func readTidyConf(path string) (TidyConf, error) {
	var conf TidyConf

	fs, err := os.Stat(path)
	if err != nil {
		return conf, err
	}

	if fs.IsDir() {
		path := filepath.Join(path, ".tidy.json")
		_, err := os.Stat(path)
		if err != nil {
			return conf, err
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return conf, err
		}

		err = json.Unmarshal(contents, &conf)
		if err != nil {
			return conf, err
		}
	}

	return conf, nil
}

func SetSym(arg string) error {
	conf, err := readTidyConf(arg)
	if err != nil {
		return err
	}

	for _, s := range conf.Home {
		absTarget, err := filepath.Abs(filepath.Join(arg, s))
		if err != nil {
			return err
		}
		err = os.Symlink(absTarget, filepath.Join(xdg.Home, s))
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				continue
			}

			return err
		}
	}
	for _, s := range conf.Config {
		absTarget, err := filepath.Abs(filepath.Join(arg, s))
		if err != nil {
			return err
		}
		err = os.Symlink(absTarget, filepath.Join(xdg.ConfigHome, s))
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				continue
			}
			return err
		}
	}
	for _, s := range conf.Bin {
		absTarget, err := filepath.Abs(filepath.Join(arg, s))
		if err != nil {
			return err
		}

		err = os.Symlink(absTarget, filepath.Join(xdg.BinHome, s))
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				continue
			}
			return err
		}
	}

	return nil
}

func UnsetSym(arg string) error {
	conf, err := readTidyConf(arg)
	if err != nil {
		return err
	}

	for _, s := range conf.Home {
		path := filepath.Join(xdg.Home, s)

		f, err := os.Lstat(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		if f.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
	}

	for _, s := range conf.Config {
		path := filepath.Join(xdg.ConfigHome, s)

		f, err := os.Lstat(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		if f.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
	}

	for _, s := range conf.Bin {
		path := filepath.Join(xdg.BinHome, s)

		f, err := os.Lstat(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		if f.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
