package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks added by user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}
