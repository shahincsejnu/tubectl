package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get is a command to see a specific resource of a resource type",
	Long:  "get is a command to see a specific resource of a resource type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}
