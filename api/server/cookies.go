package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetCookie() handles process of creating a cookie.
func (*Server) SetCookie(ctx *fiber.Ctx, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		ctx.Cookie(&fiber.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     cookie.Path,
			Domain:   cookie.Domain,
			MaxAge:   cookie.MaxAge,
			Expires:  cookie.Expires,
			Secure:   cookie.Secure,
			HTTPOnly: cookie.HttpOnly,
		})
	}
}

// SetTimedCookie() handles process of creating a cookie.
// Appart from name, value and expiring time (in minutes), func also accepts pointer to fiber context.
func (*Server) SetTimedCookie(c *fiber.Ctx, name string, exp_minutes int, value interface{}) {

	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprint(value)
	cookie.Expires = time.Now().Add(time.Duration(exp_minutes) * time.Minute)

	// Setting cookie
	c.Cookie(cookie)
}
