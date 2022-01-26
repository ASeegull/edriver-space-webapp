package server

import (
	"fmt"
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/ASeegull/edriver-space-webapp/pkg/api_client"
	"net/http"
	"strconv"

	"github.com/ASeegull/edriver-space-webapp/logger"
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
	//srv := server
	return func(ctx *fiber.Ctx) error {
		input := model.SignInInput{}

		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(err.Error())
		}

		apiResp, err := h.client.Users.SignIn(input)
		if err != nil {
			return ctx.SendString(err.Error())
		}

		h.setCookie(ctx, apiResp.Cookies)

		return ctx.Status(apiResp.StatusCode).JSON(apiResp.Body)
	}
}

func (h *Handler) ClosureSignUp(server *Server) fiber.Handler {
	//srv := server
	return func(ctx *fiber.Ctx) error {
		input := model.SignUpInput{}

		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(err.Error())
		}

		apiResp, err := h.client.Users.SignUp(input)
		if err != nil {
			return ctx.SendString(err.Error())
		}

		h.setCookie(ctx, apiResp.Cookies)

		return ctx.Status(apiResp.StatusCode).JSON(apiResp.Body)
	}
}

func (h *Handler) ClosureSignOut(server *Server) fiber.Handler {
	srv := server
	return func(ctx *fiber.Ctx) error {
		cookieName := srv.Config.CookieName

		apiResp, err := h.client.Users.SignOut(&http.Cookie{
			Name:  cookieName,
			Value: ctx.Cookies(cookieName),
		})
		if err != nil {
			return ctx.SendString(err.Error())
		}

		ctx.ClearCookie(cookieName)

		return ctx.Status(apiResp.StatusCode).JSON(apiResp.Body)
	}
}

func (h *Handler) ClosureRefreshTokens(server *Server) fiber.Handler {
	srv := server
	return func(ctx *fiber.Ctx) error {
		cookieName := srv.Config.CookieName

		ctx.Request().Header.Cookie("refreshToken")

		apiRespWithCookies, err := h.client.Users.RefreshTokens(&http.Cookie{
			Name:  cookieName,
			Value: ctx.Cookies(cookieName),
		})

		if err != nil {
			return ctx.SendString(err.Error())
		}

		h.setCookie(ctx, apiRespWithCookies.Cookies)

		return ctx.Status(apiRespWithCookies.StatusCode).JSON(apiRespWithCookies.Body)
	}
}

func (h *Handler) ClosureAddDriverLicense(server *Server) fiber.Handler {
	//srv := server
	return func(ctx *fiber.Ctx) error {

		input := model.AddDriverLicenceInput{}

		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(err.Error())
		}

		jwtHeader := ctx.Get("Authorization", "Bearer ")

		apiResp, err := h.client.Users.AddDriverLicense(input, jwtHeader)
		if err != nil {
			return ctx.SendString(err.Error())
		}

		return ctx.Status(apiResp.StatusCode).JSON(apiResp.Body)
	}
}

func (h *Handler) ClosureGetFines(server *Server) fiber.Handler {
	//srv := server
	return func(ctx *fiber.Ctx) error {

		jwtHeader := ctx.Get("Authorization", "Bearer ")

		apiResp, err := h.client.Users.GetFines(jwtHeader)
		if err != nil {
			return ctx.SendString(err.Error())
		}

		return ctx.Status(apiResp.StatusCode).JSON(apiResp.Body)
	}
}

// ClosureExit returns a webapp route closure function that handles exit from session proccess
func (h *Handler) ClosureExit(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Getting id for current session from cookies
		sesid, err := strconv.Atoi(c.Cookies("sesid"))
		if err != nil {
			logger.LogErr(err)
		}

		// Marking session as ended
		srv.EndSession(sesid)
		c.ClearCookie()
		return c.Redirect("/")
	}
}

// ClosureMain returns a webapp route closure function that handles requests to base URL
func (h *Handler) ClosureMain(server *Server) fiber.Handler {
	srv := server
	return func(c *fiber.Ctx) error {
		// Redirecting to panel page if user is already logged in. If not - redirecting to login form
		if srv.CheckAuth(c) {
			return c.Redirect("/panel")
		} else {
			return c.Render("index", fiber.Map{
				"Title": srv.Config.MainPageTitle,
				"Error": c.Cookies("LogInErr"),
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
			return c.Render("panel", fiber.Map{
				"Title": srv.Config.PanelPageTitle,
				"Error": c.Cookies("PanelErr"),
			})
		} else {
			return c.Redirect("/")

		}
	}
}

func (h *Handler) setCookie(ctx *fiber.Ctx, cookies []*http.Cookie) {
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
