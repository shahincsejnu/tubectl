package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "for updating a specific k8s workload resources",
	Long:  "This command is used to update a kubernetes workload resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
