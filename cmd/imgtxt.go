package cmd

import (
	"codeberg.org/aryak/simplytranslate/utils"
	"github.com/spf13/cobra"
	"fmt"
)

var file string

var imgtxtCmd = &cobra.Command{
	Use:   "imgtxt",
	Short: "Image -> Text.",
	Long:  `Convert given image (filename) to text using gosseract.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.ImgTxt(file))
	},
}

func init() {
	rootCmd.AddCommand(imgtxtCmd)

	imgtxtCmd.Flags().StringVarP(&file, "file", "f", "", "The query SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_query environment variable.")

	langlist = imgtxtCmd.Flag("file").Value.String()
}
