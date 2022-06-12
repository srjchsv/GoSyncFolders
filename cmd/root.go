/*
Copyright © 2022 srjchsv@gmail.com

*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wwww",
	Short: "Let GO sync folders for you",
	Long:  `A program that syncs source and destination folders of your input.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
