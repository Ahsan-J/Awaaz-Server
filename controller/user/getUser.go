package user

import (
	"fmt" // "encoding/json"
	"net/http"

	"../../helpers"
	"../../modal"
)

// GetUser path:/user
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query() // Get GET Query Paramaters

	if _, ok := params["id"]; !ok {
		userNotEnoughParams.SendAPI(w, nil)
		return
	}
	fmt.Println("Fetching for user:", params["id"][0], "From data", params)
	user := []modal.User{}

	db := helpers.GetDBInstance()
	defer db.Close()
	err := db.Select(&user, "SELECT * FROM user WHERE id='"+params["id"][0]+"'")
	if err != nil {
		fmt.Println("From Querying DB error: ", err)
		return
	}

	if len(user) > 0 {
		user[0].Optimize()
		userSuccess.SendAPI(w, user[0])
		return
	}
	userNotFound.SendAPI(w, nil)
}
