package generalapi

import (
	"fmt"
	"net/http"
	"strings"

	"../../helpers"
	"../../modal"
)

// Refresh path:/refresh
func Refresh(w http.ResponseWriter, r *http.Request) {
	var id = strings.TrimSpace(r.FormValue("id"))

	db := helpers.GetDBInstance()
	defer db.Close()

	if len(id) <= 0 {
		fieldMissing.SendAPI(w,nil)
		return
	}		
	
	var user = []modal.User{}
	
	err := db.Select(&user, "SELECT * FROM user WHERE id='"+id+"'")
	
	
	if err != nil {
		fmt.Println("Error fetching DB", err)
		internalServerError.SendAPI(w, nil)
		return
	}
	user[0].Optimize()
	s := strings.Split(r.Header["Authorization"][0], " ")
	token := helpers.GetRefreshedToken(s[1],user[0]);
	// token := helpers.CreateNewToken(user[0], apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user[0]}
	refreshSuccess.SendAPI(w,res)
}
