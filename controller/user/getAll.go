package user

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"fmt" // "encoding/json"
	"net/http"
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
		userNotFound.SendAPI(w, nil)
	}

	for i := range users {
		users[i].Optimize()
	}

	// Sending Response
	usersAllSuccess.SendAPI(w, users)
}
