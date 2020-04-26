package cmd

import (
	webapi "github.com/liturgiko/doxa/pkg/server/api"
	webapp "github.com/liturgiko/doxa/pkg/server/app"
	"github.com/spf13/cobra"
)

var api bool
var app bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs an http server with access to the liturgical database",
	Long: `serve runs an http server with access to the liturgical database.
`,
	Run: func(cmd *cobra.Command, args []string) {
		appPort := "8080"
		apiPort := "8090"
		if api {
			webapi.Serve(Paths.DbPath, apiPort)
		} else {
			webapp.Serve(Paths.DbPath,appPort, apiPort)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&api, "api", false, "serve using only the rest api")
}

