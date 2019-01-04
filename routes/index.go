package routes

import (
	"github.com/gorilla/mux"
	"./user"
	"./admin"
	"./attachment"
	"net/http"
	"html/template"
)
// GenerateRoutes combines all of the functions and generate them to be used
func GenerateRoutes (router * mux.Router) {
	admin.GenerateAdminRoutes(router)
	// Api
	user.GenerateUserAPIRoutes(router)
	attachment.GenerateAttachmentApiRoutes(router)
	
	serveStatic(router);
	router.HandleFunc("/",index);
}

func serveStatic (router *mux.Router) {
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./html/public"))))
}

func index (w http.ResponseWriter, r *http.Request){
	// Redirects to index.html inside the html folder
	template.
	Must(template.ParseFiles("./html/index.html")).
	Execute(w,nil);
}