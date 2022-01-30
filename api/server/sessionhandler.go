package server

import (
	"fmt"

	"github.com/ASeegull/edriver-space-webapp/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RegisterSession() handles process of adding new session to Server.Sessions slice.
func (server *Server) RegisterSession(session *model.Session, c *fiber.Ctx) {

	id := uuid.New().String()
	session.UserIP = c.IP()
	server.Sessions[id] = *session
	server.SetTimedCookie(c, "sesid", 60, id)
	server.SetTimedCookie(c, "refreshTime", 8, "no")
}

// EndSession(id int) handles process of deleting a session with given id from Server.Sessions slice.
func (server *Server) EndSession(uuid string) {
	delete(server.Sessions, uuid)
}

// CreateSessionFromApiResponse() checks whether response from main app contains tokens and returns session data if it is or error if it's not.
func (srv *Server) CreateSessionFromApiResponse(ctx *fiber.Ctx, email string, apiResp model.ApiResponseWithCookies) (*model.Session, error) {
	session := new(model.Session)
	session.UserLogin = email

	tokens, ok := apiResp.Body.(model.Tokens)

	if !ok {
		return new(model.Session), fmt.Errorf(fmt.Sprint(apiResp.Body))
	}
	session.AccessToken = tokens.AccessToken

	return session, nil
}

// CheckAuth() checks if user is already logged in, and returns bool value as a result.
func (server *Server) CheckAuth(c *fiber.Ctx) bool {
	res := false
	id := c.Cookies("sesid")
	if _, excist := server.Sessions[id]; excist {
		res = true
	}

	return res
}

// IsRefreshTime() checks if access token needs to be refreshed.
func (server *Server) IsRefreshTime(c *fiber.Ctx) bool {
	res := false
	if c.Cookies("refreshTime") != "no" {
		res = true
	}
	return res
}
