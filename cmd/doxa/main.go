package main

import webapp "github.com/liturgiko/doxa/pkg/server/app"

func main() {
	webapp.Serve("$HOME/doxa/data/sql/liturgical.db","8080", "8090")
}
