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
	force   bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "tidy",
		Short: "Tidy Dotfile Linker " + version,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Example: "tidy [set | unset] [directory | *]",
	}

	setCommand := &cobra.Command{
		Use:   "set [directory | *]",
		Short: "Create symlinks",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				err := linker.SetSym(arg, force)
				if errors.Is(err, os.ErrNotExist) {
					fmt.Fprintln(os.Stderr, arg, "does not have a .tidy.json file")
					err = nil
				} else if err != nil {
					return err
				}
			}
			return nil
		},
	}

	setCommand.Flags().
		BoolVarP(&force, "force", "f", false, "overwrite existing files")
	rootCmd.AddCommand(setCommand)


	unsetCommand := &cobra.Command{
		Use:   "unset [directory | *]",
		Short: "Delete symlinks",
		Aliases: []string{"u"},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				err := linker.UnsetSym(arg)
				if errors.Is(err, os.ErrNotExist) {
					fmt.Fprintln(os.Stderr, arg, "does not have a .tidy.json file")
					err = nil
				} else if err != nil {
					return err
				}
			}
			return nil
		},
	}

	rootCmd.AddCommand(unsetCommand)

	completionCmd := &cobra.Command{
		Use:    "completion [bash|zsh]",
		Short:  "Generate completion script",
		Args:   cobra.ExactArgs(1),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				return rootCmd.GenZshCompletion(os.Stdout)
			}
			return nil
		},
	}

	rootCmd.AddCommand(completionCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
