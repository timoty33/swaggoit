package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Hello = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Long:  `This command prints a hello message to the user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello, World!")

		return nil
	},
}
