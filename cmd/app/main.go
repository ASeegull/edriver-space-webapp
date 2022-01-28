package main

import (
	"github.com/ASeegull/edriver-space-webapp/api/server"
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
)

func main() {
	// Initializing logger
	logger.LogInit()

	// Loading config vals
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		logger.LogErr(err)
	}

	// Initializing server and passing config to it
	webapp := server.Init(cfg)
	webapp.BuildRoutes()

	//Starting server
	if err := webapp.Start(); err != nil {
		logger.LogFatal(err)
	}
}
