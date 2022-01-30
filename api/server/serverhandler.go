package server

import (
	"fmt"
	"time"

	"net/http"

	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/ASeegull/edriver-space-webapp/pkg/api_client"
	"github.com/ASeegull/edriver-space-webapp/pkg/sorts"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	client *api_client.ApiClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{client: api_client.NewApiClient(cfg)}
}

// ClosureGetSessions returns a webapp route closure function that proceeds technical requests for seeing all sessions
func (h *Handler) ClosureGetSessions(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprint(&srv.Sessions))
	}
}

// ClosureSignIn returns a webapp route closure function that proceeds user authorization data and starts login session
func (h *Handler) ClosureSignIn(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		input := model.SignInInput{}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		apiResp, err := h.client.Users.SignIn(input)
		if err != nil {
			return c.SendString(err.Error())
		}

		srv.SetCookie(c, apiResp.Cookies)

		session, err := srv.CreateSessionFromApiResponse(c, input.Email, apiResp)
		if err != nil {
			srv.SetTimedCookie(c, "SignInErr", 2, err)
			return c.Redirect("/")
		}

		srv.RegisterSession(session, c)
		return c.Redirect("/panel")

	}
}

func (h *Handler) ClosureSignUp(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		input := model.SignUpInput{}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		apiResp, err := h.client.Users.SignUp(input)
		if err != nil {
			return c.SendString(err.Error())
		}
		srv.SetCookie(c, apiResp.Cookies)

		session, err := srv.CreateSessionFromApiResponse(c, input.Email, apiResp)
		if err != nil {
			srv.SetTimedCookie(c, "SignUpErr", 2, err)
			return c.Redirect("/register")
		}

		srv.RegisterSession(session, c)
		return c.Redirect("/panel")

	}
}

func (h *Handler) ClosureSignOut(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		cookieName := srv.Config.CookieName

		_, err := h.client.Users.SignOut(&http.Cookie{
			Name:  cookieName,
			Value: c.Cookies(cookieName),
		})
		if err != nil {
			return c.SendString(err.Error())
		}

		srv.EndSession(c.Cookies("sesid"))
		c.ClearCookie()
		return c.Redirect("/")
	}
}

func (h *Handler) ClosureRefreshTokens(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		cookieName := srv.Config.CookieName

		c.Request().Header.Cookie("refreshToken")

		apiRespWithCookies, err := h.client.Users.RefreshTokens(&http.Cookie{
			Name:  cookieName,
			Value: c.Cookies(cookieName),
		})

		if err != nil {
			return c.SendString(err.Error())
		}

		srv.SetCookie(c, apiRespWithCookies.Cookies)
		srv.SetTimedCookie(c, "refreshTime", 8, "no")

		return c.Redirect("/panel")
	}
}

func (h *Handler) ClosureAddDriverLicense(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				input := model.AddDriverLicenceInput{}

				if err := c.BodyParser(&input); err != nil {
					return c.Status(http.StatusBadRequest).JSON(err.Error())
				}

				id := c.Cookies("sesid")
				jwtHeader := "Bearer " + srv.Sessions[id].AccessToken

				apiResp, err := h.client.Users.AddDriverLicense(input, jwtHeader)
				if err != nil {
					return c.SendString(err.Error())
				}

				return c.Status(apiResp.StatusCode).JSON(apiResp.Body)
			}
		} else {
			return c.Redirect("/")
		}
	}
}

func (h *Handler) ClosureGetFines(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				id := c.Cookies("sesid")
				jwtHeader := "Bearer " + srv.Sessions[id].AccessToken

				apiResp, err := h.client.Users.GetFines(jwtHeader)
				if err != nil {
					return c.SendString(err.Error())
				}

				fines, ok := apiResp.Body.(model.Fines)

				if !ok {
					return c.Status(apiResp.StatusCode).JSON(apiResp.Body)
				}

				return c.Render("fine-list", fiber.Map{
					"Title":       srv.Config.PanelPageTitle,
					"ListName":    server.Sessions[id].UserLogin,
					"ReturnURL":   "/panel",
					"CarFines":    fines.CarsFines,
					"DriverFines": fines.DriversFines,
				})

			}
		} else {
			return c.Redirect("/")
		}
	}
}

// ClosureSignUp() returns a webapp route closure function that proceeds user authorization data and starts login session
func (Handler) ClosureRegisterPage(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				return c.Redirect("/panel")
			}
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

// ClosureMain returns a webapp route closure function that handles requests to base URL
func (h *Handler) ClosureMain(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				return c.Redirect("/panel")
			}
		} else {
			logErr := c.Cookies("SignInErr")
			c.ClearCookie("SignInErr")
			return c.Render("index", fiber.Map{
				"Title": srv.Config.MainPageTitle,
				"Error": logErr,
			})
		}
	}
}

// ClosurePanel returns a webapp route closure function that handles requests to base URL
func (h *Handler) ClosurePanel(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to panel page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				return c.Render("panel", fiber.Map{
					"Title": srv.Config.PanelPageTitle,
					"Error": c.Cookies("PanelErr"),
				})
			}
		} else {
			return c.Redirect("/")
		}
	}
}

func (Handler) ClosureAddInfo(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to panel page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				currentTime := time.Now()
				return c.Render("add-info", fiber.Map{
					"Title": srv.Config.PanelPageTitle,
					"Date":  currentTime.Format("2006-01-02"),
				})
			}
		} else {
			return c.Redirect("/")

		}
	}
}

func (h *Handler) ClosureVehicles(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to vehicles page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				id := c.Cookies("sesid")
				jwtHeader := "Bearer " + srv.Sessions[id].AccessToken

				apiResp, err := h.client.Users.GetFines(jwtHeader)
				if err != nil {
					return c.SendString(err.Error())
				}

				fines, ok := apiResp.Body.(model.Fines)

				if !ok {
					return c.Status(apiResp.StatusCode).JSON(apiResp.Body)
				}

				cars := sorts.GetCarListFromFines(fines)

				return c.Render("vehicles", fiber.Map{
					"Title": srv.Config.PanelPageTitle,
					"Cars":  cars,
				})
			}
		} else {
			return c.Redirect("/")

		}
	}
}

func (Handler) ClosureFineSingle(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to fine list page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				return c.Render("single-fine", fiber.Map{
					"Title":         srv.Config.PanelPageTitle,
					"LicenceNumber": c.Query("LicenceNumber"),
					"NumberPlate":   c.Query("NumberPlate"),
					"IssueDate":     c.Query("IssueDate"),
					"Amount":        c.Query("Amount"),
				})
			}
		} else {
			return c.Redirect("/")

		}
	}
}

func (h *Handler) ClosureVehicleFineList(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Allowing access to fine list page if user is logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			if srv.IsRefreshTime(c) {
				return c.Redirect("/refresh-tokens")
			} else {
				id := c.Cookies("sesid")
				jwtHeader := "Bearer " + srv.Sessions[id].AccessToken

				apiResp, err := h.client.Users.GetFines(jwtHeader)
				if err != nil {
					return c.SendString(err.Error())
				}

				fines, ok := apiResp.Body.(model.Fines)

				if !ok {
					return c.Status(apiResp.StatusCode).JSON(apiResp.Body)
				}

				numberplate := c.Query("NumberPlate")
				carfines := sorts.SearchFinesByNumberPlate(fines, numberplate)

				return c.Render("fine-list", fiber.Map{
					"Title":     srv.Config.PanelPageTitle,
					"ListName":  numberplate,
					"ReturnURL": "/vehicles",
					"CarFines":  carfines,
				})
			}
		} else {
			return c.Redirect("/")

		}
	}
}
