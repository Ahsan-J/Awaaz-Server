package user

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"fmt" // "encoding/json"
	"net/http"
	"awaaz_go_server/responses"
)

// GetUser path:/user
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query() // Get GET Query Paramaters

	if _, ok := params["id"]; !ok {
		responses.UserNotEnoughParams.SendAPI(w, nil)
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
		responses.UserSuccess.SendAPI(w, user[0])
		return
	}
	responses.UserNotFound.SendAPI(w, nil)
}
