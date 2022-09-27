package user

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"fmt" // "encoding/json"
	"net/http"
	"awaaz_go_server/responses"
)

// GetAll path:/users
func GetAll(w http.ResponseWriter, r *http.Request) {
	var err error
	users := []modal.User{}
	db := helpers.GetDBInstance()
	defer db.Close()
	// Fetching Data
	err = db.Select(&users, "SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(users) <= 0 {
		responses.UserNotFound.SendAPI(w, nil)
		return
	}

	for i := range users {
		users[i].Optimize()
	}

	// Sending Response
	responses.UsersAllSuccess.SendAPI(w, users)
}
