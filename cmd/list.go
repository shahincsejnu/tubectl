package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is a command to get all the resources of your specified type",
	Long:  "list is a command to get all the resources of your specified type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}
