package user

import (
	"../../helpers"
	"net/http"
)

var usersAllSuccess = helpers.Response{ Status : http.StatusOK, Message:"Successfully fetched all users", Code : "SUCCUSR001" }