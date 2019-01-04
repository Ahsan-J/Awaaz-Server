package sockets

import (
	"github.com/googollee/go-socket.io"
	"fmt"
)

func InitializeSocket (socket *socketio.Server) {
	
	socket.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

}