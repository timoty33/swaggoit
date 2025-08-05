package commands

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Swagger documentation",

	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
