package swaggoit

import (
  "fmt"
  "github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Use:   "hello-world",
	Short: "Imprime Hello World",
	Run: func(cmd *cobra.Command, args []string) {
   		 fmt.Println(args)

    		fmt.Println("Hello World")
	},
}
