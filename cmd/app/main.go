package main

import (
	"github.com/ASeegull/edriver-space-webapp/api/server"
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
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
		logInData := new(model.SingInData)
		c.BodyParser(logInData)
		if logInData.Email == "skskuzan" {
			if logInData.Password == "oqooq1" {
				return c.SendString(`{
					"accesstoken" : "67686ds8676f7fd68766df7",
					"refreshtoken" : "67686ds8676f7fd68766df7"
				}`)
			} else {

				return c.SendString("Wrong Password")
			}
		} else {
			return c.SendString("User doesn't excist")
		}
	})

	app.Listen(":5050")
}
