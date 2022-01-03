package server

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/ASeegull/edriver-space-webapp/pkg/auth"
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
		signInData := new(model.SingInData)
		tempSession := new(model.Session)

		// Parsing login data from POST request (via html <form>) to a variable
		err := c.BodyParser(signInData)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// Sending login request to main app
		res := auth.LoginProceed(*signInData, srv.Config)
		token := new(model.AuthData)
		json.Unmarshal([]byte(res), token)

		fmt.Println(res)

		// Saving tokens to cookies
		srv.SetCookie(c, "accesstoken", token.AccessToken)
		srv.SetCookie(c, "refreshtoken", token.RefreshToken)

		// Registring new session
		tempSession.UserLogin = signInData.Email
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
