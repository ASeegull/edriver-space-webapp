package server

import (
	"encoding/json"
	"fmt"
	"time"

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

// ClosureLogin() returns a webapp route closure function that proceeds user authorization data and starts login session
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

		if res == srv.Config.WrongPassMsg || res == srv.Config.UsrNotFoundMsg {
			srv.SetCookie(c, "LogInErr", 1, res)
			return c.Redirect("/")
		} else {
			// token := new(model.AuthData)
			json.Unmarshal([]byte(res), tempSession)

			// Registring new session
			tempSession.UserLogin = signInData.Email
			srv.RegisterSession(tempSession, c)
			c.ClearCookie("LogInErr")
			return c.Redirect("/panel")
		}

	}
}

// ClosureSignUp() returns a webapp route closure function that proceeds user authorization data and starts login session
func (ServerHandler) ClosureSignUp(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			return c.Redirect("/panel")
		} else {
			signUpErr := c.Cookies("SignUpErr")
			c.ClearCookie("SignUpErr")
			return c.Render("sign-up", fiber.Map{
				"Title": srv.Config.MainPageTitle,
				"Error": signUpErr,
			})
		}
	}
}

// ClosureNewUser() returns a webapp route closure function that proceeds user authorization data and starts login session
func (ServerHandler) ClosureNewUser(server *Server) fiber.Handler {
	//srv := server
	return func(c *fiber.Ctx) error {
		signInData := new(model.SingInData)

		// Parsing login data from POST request (via html <form>) to a variable
		err := c.BodyParser(signInData)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fmt.Printf("New User: %s, Pass: %s", signInData.Email, signInData.Password)
		return c.Redirect("/")

	}
}

// ClosureExit() returns a webapp route closure function that handles exit from session proccess
func (ServerHandler) ClosureExit(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Marking session as ended and clearing cookies
		srv.EndSession(c.Cookies("sesid"))
		c.ClearCookie()
		return c.Redirect("/")
	}
}

// ClosureMain() returns a webapp route closure function that handles requests to base URL
func (ServerHandler) ClosureMain(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			return c.Redirect("/panel")
		} else {
			logErr := c.Cookies("LogInErr")
			c.ClearCookie("LogInErr")
			return c.Render("index", fiber.Map{
				"Title": srv.Config.MainPageTitle,
				"Error": logErr,
			})
		}
	}
}

// ClosureMain() returns a webapp route closure function that handles requests to base URL
func (ServerHandler) ClosurePanel(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to panel page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			return c.Render("panel", fiber.Map{
				"Title": srv.Config.PanelPageTitle,
				"Error": c.Cookies("PanelErr"),
			})
		} else {
			return c.Redirect("/")

		}
	}
}

func (ServerHandler) ClosureAddInfo(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to panel page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			currentTime := time.Now()
			return c.Render("add-info", fiber.Map{
				"Title": srv.Config.PanelPageTitle,
				"Date":  currentTime.Format("2006-01-02"),
			})
		} else {
			return c.Redirect("/")

		}
	}
}

func (ServerHandler) ClosureVehicles(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to vehicles page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			cars := []*model.Car{
				{
					VIN:              "VF3HURHCHFUUDJK206785",
					RegistrationNum:  "AA6666AA",
					VehicleCategory:  "",
					Make:             "Mazda",
					Type:             "2",
					Year:             "2013",
					RegistrationDate: "20/01/2022",
				},
				{
					VIN:              "IHFHIHOIDHOIFIOD456454",
					RegistrationNum:  "",
					VehicleCategory:  "",
					Make:             "Mazda",
					Type:             "CX-5",
					Year:             "2019",
					RegistrationDate: "21/02/2022",
				},
			}
			return c.Render("vehicles", fiber.Map{
				"Title": srv.Config.PanelPageTitle,
				"Cars":  cars,
			})
		} else {
			return c.Redirect("/")

		}
	}
}

func (ServerHandler) ClosureFineSingle(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to fine list page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			return c.Render("single-fine", fiber.Map{
				"Title":       srv.Config.PanelPageTitle,
				"VIN":         c.Query("VIN"),
				"NumberPlate": c.Query("NumberPlate"),
				"IssueDate":   c.Query("IssueDate"),
				"Place":       c.Query("Place"),
				"Violation":   c.Query("Violation"),
				"Ammount":     c.Query("Ammount"),
			})
		} else {
			return c.Redirect("/")

		}
	}
}

func (ServerHandler) ClosureFineList(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to fine list page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			fines := []*model.Fine{
				{
					VIN:         "VF3HURHCHFUUDJK206785",
					NumberPlate: "AA6666AA",
					Date:        "20-01-2022",
					Place:       "Uhorska 22, Lviv",
					Violation:   "Speeding",
					Ammount:     "250",
				},
				{
					VIN:         "HKFKJDGSHLJFSKJ78767",
					NumberPlate: "BC3066KP",
					Date:        "22-01-2022",
					Place:       "Lypnytska 2, Lviv",
					Violation:   "Speeding",
					Ammount:     "500",
				},
			}
			return c.Render("fine-list", fiber.Map{
				"Title": srv.Config.PanelPageTitle,
				"Fines": fines,
			})
		} else {
			return c.Redirect("/")

		}
	}
}
