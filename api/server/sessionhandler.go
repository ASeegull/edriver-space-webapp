package server

import (
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
)

// RegisterSession() handles process of adding new session to Server.Sessions slice.
func (server *Server) RegisterSession(session *model.Session, c *fiber.Ctx) {
	session.ID = len(server.Sessions)
	session.UserIP = c.IP()
	session.Active = true

	server.Sessions = append(server.Sessions, *session)
	server.SetCookie(c, "logged", true)
	server.SetCookie(c, "sesid", session.ID)
}

// EndSession(id int) handles process of deleting a session with given id from Server.Sessions slice.
func (server *Server) EndSession(id int) {
	server.Sessions[id].Active = false
}
