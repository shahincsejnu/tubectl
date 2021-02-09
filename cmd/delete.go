package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete is a command for deleting a specific workload resource (API object)",
	Long:  "delete is a command for deleting a specific workload resource (API object)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}
