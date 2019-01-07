package generalapi

import (
	"github.com/gorilla/mux"
	"net/http"
	"../../helpers"
)

// GenerateGeneralAPIRoutes generates the required user API to serve
func GenerateGeneralAPIRoutes(router *mux.Router) {
	router.HandleFunc("/login",login).Methods(http.MethodPost)
	router.HandleFunc("/register",register).Methods(http.MethodPost)
	refreshSub := router.PathPrefix("/refresh").Subrouter()
	refreshSub.Use(helpers.ValidateRefreshToken)
	refreshSub.HandleFunc("",refresh).Methods(http.MethodPost)
}
