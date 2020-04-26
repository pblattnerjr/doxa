package main

import "github.com/liturgiko/doxa/pkg/server/app"

func main() {
	app.Serve("$HOME/doxa/data/sql/liturgical.db","8080", "8090")
}
