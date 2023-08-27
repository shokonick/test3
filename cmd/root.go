package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mozhi",
	Short: "An alternative front-end for many Translation Engines.",
	Long:  "An alternative front-end for many Translation Engines, rewritten in Gofiber+colly by AryaK.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(os.Stderr, "Error running CLI: %v", err)
		os.Exit(1)
	}
}
