package main

import (
	"github.com/ASeegull/edriver-space-webapp/api/server"
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.LogInit()

	config, err := config.LoadConfig("./config")
	if err != nil {
		logger.LogErr(err)
	}

	webapp := server.Init(config)
	webapp.BuildRoutes()

	go webapp.Start()

	app := fiber.New()
	app.Post("/sign-in", func(c *fiber.Ctx) error {
		c.SendString(`{
			"accesstoken" : "67686ds8676f7fd68766df7"
			"refreshtoken" : "67686ds8676f7fd68766df7"
		}`)
		return nil
	})

	app.Listen(":5050")
}
