package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetCookie() handles process of creating a cookie.
// Appart from name and value of cookie, func accepts pointer to fiber context.
func (server *Server) SetCookie(c *fiber.Ctx, name string, value interface{}) {

	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprint(value)
	cookie.Expires = time.Now().Add(24 * time.Hour)

	// Setting cookie
	c.Cookie(cookie)
}

// CheckAuth() checks if user is already logged in, and returns bool value as a result.
func (server *Server) CheckAuth(c *fiber.Ctx) bool {
	res := false
	a := c.Cookies("authtoken")
	r := c.Cookies("refreshtoken")

	if a == "" && r == "" {
		res = false
	}

	if logged := c.Cookies("logged"); logged == "true" {
		res = true
	}

	return res
}
