package cmd

import (
	webapi "github.com/liturgiko/doxa/pkg/server/api"
	webapp "github.com/liturgiko/doxa/pkg/server/app"
	"github.com/spf13/cobra"
	"time"
)

var api bool
var app bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs an http server with access to the liturgical database",
	Long: `serve runs an http server with access to the liturgical database.
`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		if api {
			webapi.Serve(Paths.DbPath, "8090")
		}
		if app {
			webapp.Serve(Paths.DbPath,"8080")
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&api, "api", false, "serve using rest api")
	serveCmd.Flags().BoolVar(&app, "app", false, "serve using app. Will also start api.")
}

