package main

import (
	"awaaz_go_server/routes"
	"awaaz_go_server/sockets"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/googollee/go-socket.io.v1"
	// "./helpers"
)

var port = 7000

func main() {

	r := mux.NewRouter()
	routes.GenerateRoutes(r)

	socketServer, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	sockets.InitializeSocket(socketServer)
	go socketServer.Serve()
	defer socketServer.Close()
	r.Handle("/socket.io/", socketServer) // Handling Sockets connection url

	log.Println("Serving at port = " + strconv.Itoa(port) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
