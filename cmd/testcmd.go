package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `new command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new command")
	},
}
