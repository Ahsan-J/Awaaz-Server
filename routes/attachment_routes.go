package routes

import (
	"github.com/gorilla/mux"
	// "net/http"
	// "awaaz_go_server/helpers"
	"awaaz_go_server/controller/attachment"
)

// GenerateAttachmentAPIRoutes generates the required user API to serve
func generateAttachmentAPIRoutes(router *mux.Router) {

	sub := router.PathPrefix("/attachment").Subrouter()
	// sub.Use(helpers.LoggingMiddleware)
	
	sub.HandleFunc("/add", attachment.AddAttachment).Methods("POST")
}
