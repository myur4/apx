package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"github.com/spf13/cobra"
)

func NewSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: "apx search neovim",
		Use:     "search <packages>",
		Short:   "Search for packages in a managed container.",
		RunE:    search,
	}
	return cmd
}

func search(cmd *cobra.Command, args []string) error {

	command := append([]string{}, container.GetPkgCommand("search")...)
	command = append(command, args...)

	container.Run(command...)

	return nil
}
