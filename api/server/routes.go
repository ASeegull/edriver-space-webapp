package server

import (
	"fmt"

	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
)

// IndexPageRoute() func is basic webapp route that delivers to user welcome web page
func IndexPageRoute(c *fiber.Ctx) error {
	c.SendString("Hello Lv-644")
	return nil
}

func GetAllSessionsRoute(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprint(WebServer.Sessions))

}

// LoginRoute() func is a webapp route that proceeds user authorization data and starts login session
func LoginRoute(c *fiber.Ctx) error {

	tempSession := new(model.Session)

	err := c.BodyParser(tempSession)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	RegisterSession(tempSession, c.IP())

	return c.Redirect("/panel.html")

}

// MainPageRoute() func is a webapp route that delivers to user "Personal Cabinet" web page
func MainPageRoute(c *fiber.Ctx) error {
	c.SendString("Hello Lv-644")
	return nil
}

// VehiclesPageRoute() func is a webapp route that delivers to user "My Vehicles" web page
func VehiclesPageRoute(c *fiber.Ctx) error {
	c.SendString("Hello Lv-644")
	return nil
}

// FinesPageRoute() func is a webapp route that delivers to user "My Fines" web page
func FinesPageRoute(c *fiber.Ctx) error {
	c.SendString("Hello Lv-644")
	return nil
}

// ExitRoute() func is a webapp route that handles exit from session proccess
func ExitRoute(c *fiber.Ctx) error {
	EndSession(len(WebServer.Sessions))
	return c.Redirect("/")
}
