package main

import (
	"log"
	"net/http"
	"strconv"

	"./routes"
	"./sockets"
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
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
	// Adding Middlewares
	// r.Use(helpers.SetHeaders);

	log.Println("Serving at port = " + strconv.Itoa(port) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
