/*
Copyright © 2025 Yogesh Bhutkar <yogeshbhutkar3@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "videoconv",
	Short: "Converts mov file type to gif using ffmpeg",
	Long:  "videoconv is a simple CLI tool to convert mov file type to gif using ffmpeg.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
