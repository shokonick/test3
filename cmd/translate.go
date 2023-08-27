package cmd

import (
	"codeberg.org/aryak/mozhi/utils"
	"fmt"
	"github.com/spf13/cobra"
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
		} else if engine == "deepl" {
			if langlist == "sl" {
				fmt.Println(utils.LangListDeepl(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListDeepl(langlist))
			} else {
				fmt.Println(utils.TranslateDeepl(dest, source, query))
			}
		} else if engine == "libretranslate" {
			if langlist == "sl" {
				fmt.Println(utils.LangListLibreTranslate(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListLibreTranslate(langlist))
			} else {
				fmt.Println(utils.TranslateLibreTranslate(dest, source, query))
			}
		} else if engine == "watson" {
			if langlist == "sl" {
				fmt.Println(utils.LangListWatson(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListWatson(langlist))
			} else {
				fmt.Println(utils.TranslateWatson(dest, source, query))
			}
		} else if engine == "yandex" {
			if langlist == "sl" {
				fmt.Println(utils.LangListYandex(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListYandex(langlist))
			} else {
				fmt.Println(utils.TranslateYandex(dest, source, query))
			}
		} else if engine == "duckduckgo" {
			if langlist == "sl" {
				fmt.Println(utils.LangListDuckDuckGo(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListDuckDuckGo(langlist))
			} else {
				fmt.Println(utils.TranslateDuckDuckGo(dest, source, query))
			}
		} else if engine == "mymemory" {
			if langlist == "sl" {
				fmt.Println(utils.LangListMyMemory(langlist))
			} else if langlist == "tl" {
				fmt.Println(utils.LangListMyMemory(langlist))
			} else {
				fmt.Println(utils.TranslateMyMemory(dest, source, query))
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

	translateCmd.Flags().StringVarP(&engine, "engine", "e", "", "The source Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_source environment variable.")
	translateCmd.Flags().StringVarP(&source, "source", "s", "", "The source Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_source environment variable.")
	translateCmd.Flags().StringVarP(&dest, "dest", "t", "", "The dest Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_dest environment variable.")
	translateCmd.Flags().StringVarP(&query, "query", "q", "", "The query Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_query environment variable.")
	translateCmd.Flags().StringVarP(&langlist, "langlist", "l", "", "The query Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_query environment variable.")

	engine = translateCmd.Flag("engine").Value.String()
	dest = translateCmd.Flag("dest").Value.String()
	source = translateCmd.Flag("source").Value.String()
	query = translateCmd.Flag("query").Value.String()
	langlist = translateCmd.Flag("query").Value.String()
}
