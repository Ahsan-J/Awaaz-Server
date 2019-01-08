package generalapi

import (
	"net/http"
	"../../modal"
)

type combinedResponse struct {
	modal.Token
	modal.User
}

// Login Success
var loginSuccess = modal.Response{Status: http.StatusOK, Message: "Login Success", Code: "SUCCLOG001"}

// Login Error
var loginUserNotFound = modal.Response{Status: http.StatusNotFound, Message: "Login user not Found", Code: "ERRLOG001"}

// Register Success
var registerSuccess = modal.Response{Status: http.StatusOK, Message: "Register Successfully", Code: "SUCCREG001"}

// Success Token
var refreshSuccess = modal.Response{Status: http.StatusOK, Message: "Refreshed your access token", Code: "SUCCRT001"}

// General Error
var emailRequired = modal.Response{Status: http.StatusBadRequest, Message: "Email is required", Code: "ERRGEN001"}
var emailInvalid = modal.Response{Status: http.StatusBadRequest, Message: "Not a Valid Email address", Code: "ERRGEN002"}
var emailAlreadyInUse = modal.Response{Status: http.StatusBadRequest, Message: "Email is already in use", Code: "ERRGEN003"}
var passwordRequired = modal.Response{Status: http.StatusBadRequest, Message: "Password is required", Code: "ERRGEN004"}
var passwordMismatch = modal.Response{Status: http.StatusBadRequest, Message: "Password mismatch", Code: "ERRGEN005"}
var passwordInvalid = modal.Response{Status: http.StatusBadRequest, Message: "Password is not valid", Code: "ERRGEN006"}
var internalServerError = modal.Response{Status: http.StatusInternalServerError, Message: "Something went wrong", Code: "ERRGEN007"}
var fieldMissing = modal.Response{Status: http.StatusBadRequest, Message: "Missing required paramaters", Code: "ERRGEN008"}
var invalidAPIKey = modal.Response{Status: http.StatusForbidden, Message: "Invalid api key provided", Code: "ERRGEN009"}
var blockedAPIKey = modal.Response{Status: http.StatusForbidden, Message: "The API key is blocked", Code: "ERRGEN010"}
var versionUpdate = modal.Response{Status: http.StatusUpgradeRequired, Message: "Kindly Update to the latest version", Code: "ERRGEN011"}
var genderUnspecified = modal.Response{Status: http.StatusForbidden, Message: "Gender Unspecified", Code: "ERRGEN012"}

var userIsBlocked = modal.Response{Status: http.StatusForbidden, Message: "User is unauthorized to access", Code: "ERRUSR010"}
