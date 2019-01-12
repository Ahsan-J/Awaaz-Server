package main

import (
	"awaaz_go_server/routes"
	"awaaz_go_server/sockets"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/googollee/go-socket.io.v1"
	// "./helpers"
)

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":7000"
	}
	return ":" + port
}

func main() {
	addr := determineListenAddress()

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
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	log.Println("Serving at port = " + addr + "...")
	log.Fatal(http.ListenAndServe(addr, r))
}
