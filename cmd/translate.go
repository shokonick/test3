package cmd

import (
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"codeberg.org/aryak/mozhi/utils"

	"github.com/spf13/cobra"
)

var (
	engine   string
	query    string
	source   string
	dest     string
	langlist string
)

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate.",
	Long:  `Translate.`,
	Run: func(cmd *cobra.Command, args []string) {
		if langlist == "sl" || langlist == "tl" {
			list, err := utils.LangList(engine, langlist)
			if err != nil {
				fmt.Println(err)
			}
			idxs, err := fuzzyfinder.FindMulti(
				list,
				func(i int) string {
					return list[i].Name
				})
			if err != nil {
				fmt.Println(err)
			}
			for _, idx := range idxs {
				fmt.Println("Selected Language:", list[idx].Id, "("+list[idx].Name+")")
			}
		} else if engine == "all" {
			fmt.Println(utils.TranslateAll(dest, source, query))
		} else {
			fmt.Println(utils.Translate(engine, dest, source, query))
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
