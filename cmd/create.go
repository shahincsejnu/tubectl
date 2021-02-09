package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create is a command for creating a specific workload resource (API object)",
	Long:  "create is a command for creating a specific workload resource (API object)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}
