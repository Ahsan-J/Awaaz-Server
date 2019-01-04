package attachment

import (
	"net/http"
	"fmt"
)

func getAll (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to api users page")
}