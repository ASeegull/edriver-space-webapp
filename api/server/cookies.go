package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetCookie() handles process of creating a cookie.
// Appart from name, value and expiring time (in minutes), func also accepts pointer to fiber context.
func (*Server) SetCookie(c *fiber.Ctx, name string, exp_minutes int, value interface{}) {

	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprint(value)
	cookie.Expires = time.Now().Add(time.Duration(exp_minutes) * time.Minute)

	// Setting cookie
	c.Cookie(cookie)
}

// CheckAuth() checks if user is already logged in, and returns bool value as a result.
func (server *Server) CheckAuth(c *fiber.Ctx) bool {
	res := false
	id := c.Cookies("sesid")
	if server.Sessions[id].Active {
		res = true
	}

	return res
}
