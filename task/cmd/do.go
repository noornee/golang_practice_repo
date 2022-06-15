package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task as done and deletes from database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}
