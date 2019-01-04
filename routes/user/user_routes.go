package user

import (
	"github.com/gorilla/mux"
)

// GenerateUserAPIRoutes generates the required user API to serve
func GenerateUserAPIRoutes (router * mux.Router) {

	sub := router.PathPrefix("/user").Subrouter();
	
	// All of the Registered User API routes
	sub.HandleFunc("/",getAll);
}