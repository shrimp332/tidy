package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/shrimp332/tidy/linker"
)

var (
	version = "v1.2.1"
	set     bool
	unset   bool
	force   bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "tidy",
		Short: "Tidy Dotfile Linker " + version,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if set {
				for _, arg := range args {
					err = linker.SetSym(arg, force)
					if errors.Is(err, os.ErrNotExist) {
						fmt.Fprintln(os.Stderr, arg, "does not have a .tidy.json file")
						fmt.Fprintln(os.Stderr, err)
						err = nil
					} else if err != nil {
						return err
					}
				}
			} else if unset {
				for _, arg := range args {
					err = linker.UnsetSym(arg)
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
	rootCmd.Flags().
		BoolVarP(&force, "force", "f", false, "overwrite existing files")
	rootCmd.MarkFlagsMutuallyExclusive("set", "unset")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
