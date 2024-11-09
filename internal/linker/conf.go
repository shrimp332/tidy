package linker

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type TidyConf struct {
	Config []string            `json:"config"`
	Home   []string            `json:"home"`
	Bin    []string            `json:"bin"`
	Custom map[string][]string `json:"custom"`
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
