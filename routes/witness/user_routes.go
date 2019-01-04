package user

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func GenerateUserApiRoutes (router * mux.Router) {
	sub := router.PathPrefix("/user").Subrouter();
	sub.HandleFunc("/",index);
}

func index (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to api users page")
}