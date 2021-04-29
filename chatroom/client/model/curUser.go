package model

import (
	"net"
	"github.com/tcp/chatroom/common/message"
)
//因为在客户端很多地方用到curUser做成全局的

type CurUser struct {
	Conn net.Conn
	message.User
}