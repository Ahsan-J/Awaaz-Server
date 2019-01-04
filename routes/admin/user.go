package admin

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func GenerateUserRoutes (router * mux.Router) {
	sub := router.PathPrefix("/user").Subrouter();

	sub.HandleFunc("/",user)
}

func user (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to admin users page")
}