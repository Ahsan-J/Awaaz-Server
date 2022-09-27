package generalapi

import (
	"fmt"
	"net/http"
	"strings"

	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"awaaz_go_server/responses"
)

// Refresh path:/refresh
func Refresh(w http.ResponseWriter, r *http.Request) {
	var id = strings.TrimSpace(r.FormValue("id"))

	db := helpers.GetDBInstance()
	defer db.Close()

	if len(id) <= 0 {
		responses.FieldMissing.SendAPI(w,nil)
		return
	}		
	
	var user = []modal.User{}
	
	err := db.Select(&user, "SELECT * FROM user WHERE id='"+id+"'")
	
	
	if err != nil {
		fmt.Println("Error fetching DB", err)
		responses.InternalServerError.SendAPI(w, nil)
		return
	}
	user[0].Optimize()
	s := strings.Split(r.Header["Authorization"][0], " ")
	token := helpers.GetRefreshedToken(s[1],user[0]);
	// token := helpers.CreateNewToken(user[0], apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user[0]}
	responses.RefreshSuccess.SendAPI(w,res)
}
