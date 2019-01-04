package user

import (
	"net/http"
	"fmt"
	"encoding/json"
	"../../helpers"
)

func getAll (w http.ResponseWriter, r *http.Request){
	fmt.Println(r);
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err);
		return
	}
}