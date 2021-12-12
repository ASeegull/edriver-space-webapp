package server

import (
	"github.com/ASeegull/edriver-space-webapp/model"
)

func RegisterSession(session *model.Session, ip string) {
	session.ID = len(WebServer.Sessions)
	session.UserIP = ip
	session.Active = true

	WebServer.Sessions = append(WebServer.Sessions, *session)
}

func EndSession(id int) {
	WebServer.Sessions[id-1].Active = false
}
