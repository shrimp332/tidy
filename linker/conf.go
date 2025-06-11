package linker

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrNotDirectory = errors.New("path is not a directory")

type tidyPaths map[string]string

type tidyConf struct {
	Config []string            `json:"config"`
	Home   []string            `json:"home"`
	Bin    []string            `json:"bin"`
	Custom map[string][]string `json:"custom"`
}

func readTidyConf(r io.Reader) (tidyConf, error) {
	var conf tidyConf

	data, err := io.ReadAll(r)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(data, &conf)

	return conf, err
}

func getLinkPaths(path string) (tidyPaths, error) {
	fs, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if fs.IsDir() {
		tidyPath := filepath.Join(path, ".tidy.json")

		f, err := os.Open(tidyPath)
		if err != nil {
			return nil, err
		}

		t, err := readTidyConf(f)
		if err != nil {
			return nil, err
		}

		paths := tidyPaths{}
		configDir, err := os.UserConfigDir()
		for _, s := range t.Config {
			source, err := filepath.Abs(filepath.Join(path, s))
			if err != nil {
			}

			if err != nil {
				return nil, err
			}
			dest := filepath.Join(configDir, s)

			paths[dest] = source
		}

		homeDir, err := os.UserHomeDir()
		for _, s := range t.Home {
			source, err := filepath.Abs(filepath.Join(path, s))
			if err != nil {
				return nil, err
			}

			dest := filepath.Join(homeDir, s)

			paths[dest] = source
		}

		binDir := filepath.Join(homeDir, ".local/bin")
		for _, s := range t.Bin {
			source, err := filepath.Abs(filepath.Join(path, s))
			if err != nil {
				return nil, err
			}

			dest := filepath.Join(binDir, s)

			paths[dest] = source
		}

		// custom
		for d, s := range t.Custom {
			if strings.HasPrefix(d, "~") {
				d = filepath.Join(homeDir, d[1:])
			}

			for _, ss := range s {
				source, err := filepath.Abs(filepath.Join(path, ss))
				if err != nil {
					return nil, err
				}

				dest := filepath.Join(d, ss)

				paths[dest] = source
			}
		}

		return paths, nil
	}

	return nil, ErrNotDirectory
}
