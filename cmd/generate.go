/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new gif from mov file",
	Long:  "Generate a new gif from mov file using ffmpeg.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Add flags for input and output file paths.
	generateCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Input file path")
	generateCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Output file path")

	// Input path should be mandatory.
	generateCmd.MarkFlagRequired("input")
}
