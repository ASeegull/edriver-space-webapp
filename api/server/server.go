package server

import (
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
)

// Server struct holds all important server info
type Server struct {
	App      *fiber.App
	Sessions []model.Session
}

var WebServer Server

func Init() {

	WebServer.App = fiber.New()
	WebServer.Sessions = []model.Session{}

	BuildRoutes()

}

func BuildRoutes() {
	WebServer.App.Static("/", "./public")
	WebServer.App.Post("/newsession", LoginRoute)
	WebServer.App.Get("/cabinet", MainPageRoute)
	WebServer.App.Get("/cabinet/vehicles", VehiclesPageRoute)
	WebServer.App.Get("/cabinet/fines", FinesPageRoute)

	WebServer.App.Get("/getses", GetAllSessionsRoute)

	WebServer.App.Get("/exit", ExitRoute)
}

func Start() {
	logger.LogFatal(WebServer.App.Listen(":3000"))
}
