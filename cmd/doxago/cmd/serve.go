package cmd

import (
	"github.com/liturgiko/doxa/pkg/server"
	"github.com/spf13/cobra"
	"time"
)

var ages bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs an http server with access to the liturgical database",
	Long: `serve runs an http server with access to the liturgical database.
`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()
		server.Ages(Paths.DbPath, "8080")
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&ages, "ages", true, "serve using AGES database")
}

