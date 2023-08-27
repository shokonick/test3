package cmd

import (
	"fmt"

	"codeberg.org/aryak/mozhi/utils"

	"github.com/spf13/cobra"
)

var file string

var imgtxtCmd = &cobra.Command{
	Use:   "imgtxt",
	Short: "Image -> Text.",
	Long:  `Convert given image (filename) to text using gosseract.`,
	Run: func(cmd *cobra.Command, args []string) {
		text, err := utils.ImgTxt(file)
		if err != nil {
			fmt.Println("Failed to convert image to text: ", err)
		} else {
			fmt.Println(text)
		}
	},
}

func init() {
	rootCmd.AddCommand(imgtxtCmd)

	imgtxtCmd.Flags().StringVarP(&file, "file", "f", "", "The query Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_query environment variable.")

	langlist = imgtxtCmd.Flag("file").Value.String()
}
