package server

import (
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/logger"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// Server struct holds all important server info
type Server struct {
	App      *fiber.App
	Sessions []model.Session
	Config   *config.Config
	Handler  ServerHandler
}

// Init() method inizializes server
func Init(config *config.Config) *Server {

	server := new(Server)
	engine := html.New("./public", ".html")
	server.App = fiber.New(fiber.Config{
		Views: engine,
	})
	server.Sessions = []model.Session{}
	server.Config = config

	return server
}

// BuidRoutes() method setting routes for public requests.
func (server *Server) BuildRoutes() {
	server.App.Static("/public", "./public")
	server.App.Get("/cabinet/vehicles", VehiclesPageRoute)
	server.App.Get("/cabinet/fines", FinesPageRoute)
	server.App.Get("/showtokens", ShowTokens)

	server.App.Get("/entry", server.Handler.ClosureMain(server))
	server.App.Get("/", VehiclesPageRoute)
	server.App.Post("/newsession", server.Handler.ClosureLogin(server))
	server.App.Get("/panel", server.Handler.ClosurePanel(server))
	server.App.Get("/exit", server.Handler.ClosureExit(server))

	server.App.Get("/getses", server.Handler.ClosureGetSessions(server))
}

// Start() method starts server.
func (server *Server) Start() {
	logger.LogFatal(server.App.Listen(":3000"))
}
