package commands

import (
	"swaggoit/commands/cmd"
)

func init() {
	rootCmd.AddCommand(cmd.Hello)
}
