package attachment

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func GenerateAttachmentApiRoutes (router * mux.Router) {
	sub := router.PathPrefix("/attachment").Subrouter();
	sub.HandleFunc("/",index);
}

func index (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to api Attachment page")
}