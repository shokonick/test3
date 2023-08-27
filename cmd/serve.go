package cmd

import (
	"github.com/spf13/cobra"

	"codeberg.org/aryak/mozhi/serve"
)

var port string = "3000"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server.",
	Long:  `Start the web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&port, "port", "p", "", "The port Mozhi will listen to. Defaults to 3000, and overrides the MOZHI_PORT environment variable.")

	// set port variable to the value of the port flag
	port = serveCmd.Flag("port").Value.String()
}
