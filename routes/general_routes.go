package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"awaaz_go_server/helpers"
	"awaaz_go_server/controller/generalapi"
)

// GenerateGeneralAPIRoutes generates the required user API to serve
func generateGeneralAPIRoutes(router *mux.Router) {
	
	router.HandleFunc("/login",generalapi.Login).Methods(http.MethodPost)
	router.HandleFunc("/register",generalapi.Register).Methods(http.MethodPost)
	
	refreshSub := router.PathPrefix("/refresh").Subrouter()
	refreshSub.Use(helpers.ValidateRefreshToken)
	refreshSub.HandleFunc("",generalapi.Refresh).Methods(http.MethodPost)
}
