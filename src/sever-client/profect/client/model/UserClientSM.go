package model

import (
	"net"
	"sever-client/profect/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}