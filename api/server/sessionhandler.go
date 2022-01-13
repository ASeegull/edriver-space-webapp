package server

import (
	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RegisterSession() handles process of adding new session to Server.Sessions slice.
func (server *Server) RegisterSession(session *model.Session, c *fiber.Ctx) {

	id := uuid.New().String()
	session.UserIP = c.IP()
	session.Active = true

	server.Sessions[id] = *session
	server.SetCookie(c, "sesid", 60, id)
}

// EndSession(id int) handles process of deleting a session with given id from Server.Sessions slice.
func (server *Server) EndSession(uuid string) {
	bufferSession := server.Sessions[uuid]
	bufferSession.Active = false
	server.Sessions[uuid] = bufferSession
}
