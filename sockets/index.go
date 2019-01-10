package sockets

import (
	"fmt"

	"gopkg.in/googollee/go-socket.io.v1"
)

func InitializeSocket(socket *socketio.Server) {

	socket.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

}
