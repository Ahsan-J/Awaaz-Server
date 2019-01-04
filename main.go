package main
import (
	"log"
	"net/http"
	"strconv"
	"./routes"
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"./sockets"
	"./helpers"
)
// var config = DB_config {Host : "db4free.net",User:"awaaz_admin",Password:"qwerty12345",Port:3306,Database : "awaaz_case_store"}

var port = 7000;

func main() {

	r := mux.NewRouter()
	// // DBConn = InitiateConnection(&config)
	// // SelectAllUsers()
	
	// fmt.Println("Serving at port 3000")
	routes.GenerateRoutes(r);
	// http.ListenAndServe(":3000",router)
	
	
	socketServer, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	sockets.InitializeSocket(socketServer);
	go socketServer.Serve() 
	defer socketServer.Close();
	r.Handle("/socket.io/", socketServer); // Handling Sockets connection url
	// Adding Middlewares
	r.Use(helpers.SetHeaders);

	log.Println("Serving at port = "+strconv.Itoa(port)+"...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}