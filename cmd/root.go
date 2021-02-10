package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oka",
	Short: "Main Command",
	Long:  "This is the main command, rest of them are child of it, basically the binary file name (tubectl) represents this, don't need to put <oka> command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello From tubectl")
	},
}

// It will call from the main.go
func Execute() {
	//fmt.Println("in root execute")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println("in root execute 2")
}
