package user

import (
	"net/http"

	"../../modal"
)

// Success
var usersAllSuccess = modal.Response{Status: http.StatusOK, Message: "Successfully fetched all users", Code: "SUCCUSR001"}
var userSuccess = modal.Response{Status: http.StatusOK, Message: "Success", Code: "SUCCUSR002"}

// Errors
var userNotFound = modal.Response{Status: http.StatusNotFound, Message: "Requested User was not found", Code: "ERRUSR001"}
var userNotEnoughParams = modal.Response{Status: http.StatusBadRequest, Message: "Not Enough Parameters provided"}
