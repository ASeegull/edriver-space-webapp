package server

import (
	"github.com/gofiber/fiber/v2"
)

// VehiclesPageRoute() func is a webapp route that delivers to user "My Vehicles" web page
func VehiclesPageRoute(c *fiber.Ctx) error {
	return c.SendString("Hello Lv-644")
}

// FinesPageRoute() func is a webapp route that delivers to user "My Fines" web page
func FinesPageRoute(c *fiber.Ctx) error {
	return c.SendString("Hello Lv-644")
}

// ShowTokens() - technical,temporary
func ShowTokens(c *fiber.Ctx) error {
	return c.SendString(c.Cookies("accesstoken") + " // " + c.Cookies("refreshtoken"))
}
