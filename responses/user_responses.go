package responses

import (
	"awaaz_go_server/modal"
	"net/http"
)

// Success
var UsersAllSuccess = modal.Response{Status: http.StatusOK, Message: "Successfully fetched all users", Code: "SUCCUSR001"}
var UserSuccess = modal.Response{Status: http.StatusOK, Message: "Success", Code: "SUCCUSR002"}

// Errors
var UserNotFound = modal.Response{Status: http.StatusNotFound, Message: "Requested User was not found", Code: "ERRUSR001"}
var UserNotEnoughParams = modal.Response{Status: http.StatusBadRequest, Message: "Not Enough Parameters provided"}
