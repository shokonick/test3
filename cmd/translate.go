package cmd

import (
	"github.com/spf13/cobra"
	"codeberg.org/aryak/simplytranslate/utils"
	"fmt"
)

var engine string
var query string
var source string
var dest string
var langlist string

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate.",
	Long:  `Translate.`,
	Run: func(cmd *cobra.Command, args []string) {
		if engine == "reverso" {
			if langlist == "sl" {
				fmt.Println(utils.LangListReverso(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListReverso(langlist))
			} else {
				fmt.Println(utils.TranslateReverso(dest, source, query))
			}
		} else if engine == "all" {
			fmt.Println(utils.TranslateAll(dest, source, query))
		} else {
			if langlist == "sl" {
				fmt.Println(utils.LangListGoogle(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListGoogle(langlist))
			} else {
				fmt.Println(utils.TranslateGoogle(dest, source, query))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringVarP(&engine, "engine", "e", "", "The source SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_source environment variable.")
	translateCmd.Flags().StringVarP(&source, "source", "s", "", "The source SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_source environment variable.")
	translateCmd.Flags().StringVarP(&dest, "dest", "t", "", "The dest SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_dest environment variable.")
	translateCmd.Flags().StringVarP(&query, "query", "q", "", "The query SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_query environment variable.")
	translateCmd.Flags().StringVarP(&langlist, "langlist", "l", "", "The query SimplyTranslate will listen to. Defaults to 3000, and overrides the SIMPLYTRANSLATE_query environment variable.")

	engine = translateCmd.Flag("engine").Value.String()
	dest = translateCmd.Flag("dest").Value.String()
	source = translateCmd.Flag("source").Value.String()
	query = translateCmd.Flag("query").Value.String()
	langlist = translateCmd.Flag("query").Value.String()
}
