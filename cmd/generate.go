/*
Copyright © 2025 Yogesh Bhutkar <yogeshbhutkar3@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

var (
	inputFilePath  string
	outputFilePath string
	fps            int
	quality        int
	size           string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert video files from any format to any format.",
	Long: `Convert video files (MOV, MP4, etc.) from any to any format.

Supports customization of:
- Frame rate (FPS)
- Output quality
- Resolution/size
- Output filename

Example: videoconv convert -i video.mov -o animation.gif -f 15 -q 20 -s 640x480`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" {
			cmd.Help()
			return
		}

		inputExtension := filepath.Ext(inputFilePath)
		if outputFilePath == "" {
			outputFilePath = fmt.Sprintf("output.%s", inputExtension)
		}

		kwargs := ffmpeg_go.KwArgs{
			"r":   fps,
			"q:v": quality,
		}

		if size != "" {
			kwargs["s"] = size
		}

		err := ffmpeg_go.Input(inputFilePath).Output(outputFilePath, kwargs).OverWriteOutput().Run()
		if err != nil {
			cmd.PrintErr(err)
		}

		cmd.Println("Generated file at", outputFilePath)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Add flags for input and output file paths.
	convertCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Input file path")
	convertCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Output file path")
	convertCmd.Flags().IntVarP(&fps, "fps", "f", 10, "Frames per second (default 10). Common values: 10-60 fps.")
	convertCmd.Flags().IntVarP(&quality, "quality", "q", 31, "Quality (default 31). Common values: 1-31.")
	convertCmd.Flags().StringVarP(&size, "size", "s", "", "Size of the output gif file. Common values: 320x240, 640x480, 1280x720, 1920x1080.")
}
