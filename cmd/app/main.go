package main

import (
	"github.com/ASeegull/edriver-space-webapp/api/server"
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
)

var Users []model.SingInData

func FindUser(data model.SingInData) (bool, int) {
	for id, user := range Users {
		if user.Email == data.Email {
			return true, id
		}
	}
	return false, -1
}

// Temporary plug for handling API-requests.
func LoginPlug() {
	Users = append(Users, model.SingInData{Email: "skskuzan",
		Password: "oqooq1"}, model.SingInData{Email: "student",
		Password: "oqooq1"})
	app := fiber.New()
	app.Post("/sign-in", func(c *fiber.Ctx) error {
		logInData := new(model.SingInData)
		c.BodyParser(logInData)
		if found, id := FindUser(*logInData); found {
			if logInData.Password == Users[id].Password {
				return c.SendString(`{
					"accesstoken" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxIiwiUm9sZSI6InVzZXIiLCJleHAi0jE2Mzk3NDQxNDcsInN1YiI6IjEifQ.LpvZ0d9zP2R0q04kldFQEgUcHyjGZZyjkUZBegm0D9A",
					"refreshtoken" : "f8e9ecc6b5f0086220224ce285ea6fd22f77f62b20cf4335baa2b110a9bcce2f"
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

func main() {
	// Initializing logger
	logger.LogInit()

	// Loading config vals
	config, err := config.LoadConfig("./config")
	if err != nil {
		logger.LogErr(err)
	}

	// Initializing server and passing config to it
	webapp := server.Init(config)
	webapp.BuildRoutes()

	// Starting server
	go webapp.Start()

	LoginPlug()
}
