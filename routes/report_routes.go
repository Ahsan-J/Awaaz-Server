package routes

import (
	"github.com/gorilla/mux"
	// "net/http"
	// "awaaz_go_server/helpers"
	"awaaz_go_server/controller/report"
)

// GenerateReportAPIRoutes generates the required user API to serve
func generateReportAPIRoutes(router *mux.Router) {
	reports := router.PathPrefix("/reports").Subrouter()
	// reports.Use(helpers.LoggingMiddleware)
	reports.HandleFunc("", report.GetAll).Methods("GET") // GET Request to handle all data present in the Database

}
