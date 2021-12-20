package server

import (
	"fmt"
	"strconv"

	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
)

type ServerHandler struct {
}

// ClosureGetSessions() returns a webapp route closure function that proceeds technical requests for seeing all sessions
func (ServerHandler) ClosureGetSessions(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprint(&srv.Sessions))
	}
}

// ClosureRoute() returns a webapp route closure function that proceeds user authorization data and starts login session
func (ServerHandler) ClosureLogin(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {

		tempSession := new(model.Session)

		// Parsing login data from POST request (via html <form>) to new session info
		err := c.BodyParser(tempSession)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		srv.RegisterSession(tempSession, c)

		return c.Redirect("./public/panel.html")

	}
}

// ClosureExit() returns a webapp route closure function that handles exit from session proccess
func (ServerHandler) ClosureExit(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Getting id for current session from cookies
		sesid, err := strconv.Atoi(c.Cookies("sesid"))
		if err != nil {
			logger.LogErr(err)
		}

		srv.EndSession(sesid)
		c.ClearCookie()
		return c.Redirect("/public")
	}
}

// ClosureMain() returns a webapp route closure function that handles requests to base URL
func (ServerHandler) ClosureMain(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			c.Redirect("./public/panel.html")
		} else {
			c.Redirect("./public")
		}
		return nil
	}
}
