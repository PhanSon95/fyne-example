package main

import (
	"os"

	"myapp/desktop"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "Desktop app service",
	Long:             `This is a command runner desktop app service.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		desktop.Run()
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
