package server

import (
	"github.com/gofiber/fiber/v2"
)

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
