package linker

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadTidyConf(t *testing.T) {
	sample := []byte(`{
		"home": [".zshrc", ".zshenv"],
		"config": ["zsh"],
		"bin": ["increasevol.sh"],
		"custom": {
			"~/Notes": ["obsidian"],
			"~/Documents": ["passwords.kbdx"]
		}
	}`)

	expected := tidyConf{
		Home:   []string{".zshrc", ".zshenv"},
		Config: []string{"zsh"},
		Bin:    []string{"increasevol.sh"},
		Custom: map[string][]string{
			"~/Notes":     {"obsidian"},
			"~/Documents": {"passwords.kbdx"},
		},
	}

	r := bytes.NewReader(sample)

	result, err := readTidyConf(r)
	if err != nil {
		t.Fatal("readTidyConf returned error: ", err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatal(fmt.Println("\x1b[0;31mResult:\n", result, "\x1b[0;32m\nExpected:\n", expected, "\x1b[0m"))
	}
}

/* Assumptions:
*  - config is at $HOME/.config
*  - bin is at $HOME/.local/bin
 */
func TestGetLinkPaths(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Failed to setup test - Get home dir")
	}
	tmpDir, err := os.MkdirTemp("", "tidy-test-")
	if err != nil {
		t.Fatal("Failed to setup test - Create /tmp/tidy-test- dir")
	}
	defer os.RemoveAll(tmpDir)

	sample := []byte(`{
		"home": [".zshrc", ".zshenv"],
		"config": ["zsh"],
		"bin": ["increasevol.sh"],
		"custom": {
			"~/Notes": ["obsidian"],
			"~/Documents": ["passwords.kbdx"],
			"/tmp": ["testfile"]
		}
	}`)

	os.WriteFile(filepath.Join(tmpDir, ".tidy.json"), sample, 0600)
	if err != nil {
		t.Fatal("Failed to setup test - Create /tmp/tidy-test-/.tidy.json")
	}

	expected := tidyPaths{
		filepath.Join(homeDir, ".config/zsh"):               filepath.Join(tmpDir, "zsh"),
		filepath.Join(homeDir, ".zshrc"):                    filepath.Join(tmpDir, ".zshrc"),
		filepath.Join(homeDir, ".zshenv"):                   filepath.Join(tmpDir, ".zshenv"),
		filepath.Join(homeDir, ".local/bin/increasevol.sh"): filepath.Join(tmpDir, "increasevol.sh"),
		filepath.Join(homeDir, "Documents/passwords.kbdx"):  filepath.Join(tmpDir, "passwords.kbdx"),
		filepath.Join(homeDir, "Notes/obsidian"):            filepath.Join(tmpDir, "obsidian"),
		filepath.Join("/tmp", "testfile"):                   filepath.Join(tmpDir, "testfile"),
	}

	result, err := getLinkPaths(tmpDir)
	if err != nil {
		t.Fatal("getLinkPaths returned error: ", err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatal(fmt.Println("\x1b[0;31mResult:\n", result, "\x1b[0;32m\nExpected:\n", expected, "\x1b[0m"))
	}
}
