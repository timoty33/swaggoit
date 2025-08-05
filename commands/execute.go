package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaggoit",
	Short: "Swaggoit is a CLI tool for managing Swagger documentation.",
	Long:  `Swaggoit is a command-line interface tool that helps you manage and generate Swagger documentation for your APIs.`,
}

func Execute(version string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
