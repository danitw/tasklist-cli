package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sunflower",
	Short: "A Cute Tasklist CLI application",
	Long:  "A simple CLI application for tasklists built in Go using Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the root command is called
		fmt.Println("Welcome to tasklist cli")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
