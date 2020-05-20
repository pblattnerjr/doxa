package cmd

import (
	webapi "github.com/liturgiko/doxa/pkg/server/api"
	webapp "github.com/liturgiko/doxa/pkg/server/app"
	"github.com/spf13/cobra"
)

var api bool
var app bool
var site bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs an http server with access to the liturgical database",
	Long: `serve runs an http server with access to the liturgical database.
`,
	Run: func(cmd *cobra.Command, args []string) {
		appPort := "8080"
		sitePort := "8085"
		apiPort := "8090"
		if api { // serve only the api
			webapi.Serve(Paths.DbPath, apiPort)
		} else if site {
			webapp.ServeGeneratedSite(Paths.SitePath, sitePort)
		} else { // serve both the web app and the api, locally
			webapp.ServeLocal(Paths.DbPath,Paths.SitePath, appPort, apiPort)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&api, "api", false, "serve using only the rest api")
	serveCmd.Flags().BoolVar(&site, "site", false, "serve generated site only")
}

