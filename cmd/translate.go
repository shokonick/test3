package cmd

import (
	"fmt"
	"codeberg.org/aryak/mozhi/utils"

	"github.com/spf13/cobra"
)

var (
	engine   string
	query    string
	source   string
	dest     string
)

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate.",
	Run: func(cmd *cobra.Command, args []string) {
		if engine == "all" {
			data := utils.TranslateAll(dest, source, query)
			for i := 0; i < len(data); i++ {
				fmt.Println("-----------------------------------")
				fmt.Println("Engine: "+data[i].Engine)
				fmt.Println("Translated Text: "+data[i].OutputText)
				if source == "auto" {
					fmt.Println("Detected Language: "+data[i].AutoDetect)
				}
				fmt.Println("Source Language: "+data[i].SourceLang)
				fmt.Println("Target Language: "+data[i].TargetLang)
			}
		} else {
			data, err := utils.Translate(engine, dest, source, query)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Translated Text: "+data.OutputText)
			if source == "auto" {
				fmt.Println("Detected Language: "+data.AutoDetect)
			}
			fmt.Println("Source Language: "+data.SourceLang)
			fmt.Println("Target Language: "+data.TargetLang)
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
	translateCmd.Flags().SortFlags = false

	translateCmd.Flags().StringVarP(&engine, "engine", "e", "", "[google|libre|reverso|deepl|watson|yandex|mymemory|duckduckgo]")
	translateCmd.Flags().StringVarP(&source, "source", "s", "", "Source language. Use langlist command to get code for your language")
	translateCmd.Flags().StringVarP(&dest, "dest", "t", "", "Target language. Use langlist command to get code for your language")
	translateCmd.Flags().StringVarP(&query, "query", "q", "", "Text to be translated")

	translateCmd.MarkFlagRequired("engine")
	translateCmd.MarkFlagRequired("source")
	translateCmd.MarkFlagRequired("dest")
	translateCmd.MarkFlagRequired("query")
}
