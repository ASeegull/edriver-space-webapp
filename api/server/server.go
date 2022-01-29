package server

import (
	"github.com/ASeegull/edriver-space-webapp/config"
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// Server struct holds all important server info
type Server struct {
	App      *fiber.App
	Sessions map[string]model.Session
	Config   *config.Config
	Handler  *Handler
}

// Init method inizializes server
func Init(config *config.Config) *Server {

	server := new(Server)
	engine := html.New("./public", ".html")
	server.App = fiber.New(fiber.Config{
		Views: engine,
	})
	server.Sessions = make(map[string]model.Session)
	server.Config = config
	server.Handler = NewHandler(config)

	return server
}

// BuildRoutes method setting routes for public requests.
func (server *Server) BuildRoutes() {
	server.App.Static("/public", "./public")
	server.App.Get("/cabinet/vehicles", VehiclesPageRoute)
	server.App.Get("/cabinet/fines", FinesPageRoute)
	// Technical route, temporary
	server.App.Get("/showtokens", ShowTokens)

	server.App.Get("/", server.Handler.ClosureMain(server))

	server.App.Post("/sign-in", server.Handler.ClosureSignIn(server))
	server.App.Post("/sign-up", server.Handler.ClosureSignUp(server))
	server.App.Post("/sign-out", server.Handler.ClosureSignOut(server))
	server.App.Get("/refresh-tokens", server.Handler.ClosureRefreshTokens(server))
	server.App.Post("/add-driver-licence", server.Handler.ClosureAddDriverLicense(server))
	server.App.Get("/fines", server.Handler.ClosureGetFines(server))

	server.App.Post("/newuser", server.Handler.ClosureNewUser(server))

	server.App.Get("/panel", server.Handler.ClosurePanel(server))

	server.App.Get("/add-info", server.Handler.ClosureAddInfo(server))
	server.App.Get("/vehicles", server.Handler.ClosureVehicles(server))
	server.App.Get("/vehiclefines", server.Handler.ClosureVehicleFineList(server))
	server.App.Get("/fines", server.Handler.ClosureFineList(server))
	server.App.Get("/fine", server.Handler.ClosureFineSingle(server))
	server.App.Get("/exit", server.Handler.ClosureExit(server))

	server.App.Get("/getses", server.Handler.ClosureGetSessions(server))
}

// Start method starts server.
func (server *Server) Start() error {
	return server.App.Listen(":3000")
}
