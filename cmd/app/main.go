package main

import (
	"github.com/ASeegull/edriver-space-webapp/api/server"
)

func main() {
	webapp := server.Init()
	webapp.BuildRoutes()

	webapp.Start()
}
