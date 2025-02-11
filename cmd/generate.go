/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new gif from mov file",
	Long:  "Generate a new gif from mov file using ffmpeg.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
