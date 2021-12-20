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
	Handler  ServerHandler
}

// Init() method inizializes server
func Init() *Server {

	server := new(Server)
	server.App = fiber.New()
	server.Sessions = []model.Session{}

	return server
}

// BuidRoutes() method setting routes for public requests.
func (server *Server) BuildRoutes() {
	server.App.Static("/public", "./public")
	server.App.Get("/cabinet/vehicles", VehiclesPageRoute)
	server.App.Get("/cabinet/fines", FinesPageRoute)

	server.App.Get("/", server.Handler.ClosureMain(server))
	server.App.Post("/newsession", server.Handler.ClosureLogin(server))
	server.App.Get("/exit", server.Handler.ClosureExit(server))

	server.App.Get("/getses", server.Handler.ClosureGetSessions(server))
}

// Start() method starts server.
func (server *Server) Start() {
	logger.LogFatal(server.App.Listen(":3000"))
}
