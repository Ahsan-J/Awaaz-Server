package user

import (
	"net/http"
	"fmt"
	// "encoding/json"
	"../../helpers"
)

func getAll (w http.ResponseWriter, r *http.Request){
	var err error
	users := []User{};
	db := helpers.GetDBInstance();
	defer db.Close();
	// Fetching Data
	err = db.Select(&users, "SELECT * FROM user")
	if err != nil {
		fmt.Println(err);
		return
	}

	for i := range users {
		users[i].optimize();
	}

	// Sending Response
	usersAllSuccess.SendAPI(w,users);
}