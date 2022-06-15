package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}
