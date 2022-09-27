package responses

import (
	"net/http"
	"awaaz_go_server/modal"
)

// Login Success
var LoginSuccess = modal.Response{Status: http.StatusOK, Message: "Login Success", Code: "SUCCLOG001"}

// Login Error
var LoginUserNotFound = modal.Response{Status: http.StatusNotFound, Message: "Login user not Found", Code: "ERRLOG001"}

// Register Success
var RegisterSuccess = modal.Response{Status: http.StatusOK, Message: "Register Successfully", Code: "SUCCREG001"}

// Success Token
var RefreshSuccess = modal.Response{Status: http.StatusOK, Message: "Refreshed your access token", Code: "SUCCRT001"}

// General Error
var EmailRequired = modal.Response{Status: http.StatusBadRequest, Message: "Email is required", Code: "ERRGEN001"}
var EmailInvalid = modal.Response{Status: http.StatusBadRequest, Message: "Not a Valid Email address", Code: "ERRGEN002"}
var EmailAlreadyInUse = modal.Response{Status: http.StatusBadRequest, Message: "Email is already in use", Code: "ERRGEN003"}
var PasswordRequired = modal.Response{Status: http.StatusBadRequest, Message: "Password is required", Code: "ERRGEN004"}
var PasswordMismatch = modal.Response{Status: http.StatusBadRequest, Message: "Password mismatch", Code: "ERRGEN005"}
var PasswordInvalid = modal.Response{Status: http.StatusBadRequest, Message: "Password is not valid", Code: "ERRGEN006"}
var InternalServerError = modal.Response{Status: http.StatusInternalServerError, Message: "Something went wrong", Code: "ERRGEN007"}
var FieldMissing = modal.Response{Status: http.StatusBadRequest, Message: "Missing required paramaters", Code: "ERRGEN008"}
var InvalidAPIKey = modal.Response{Status: http.StatusForbidden, Message: "Invalid api key provided", Code: "ERRGEN009"}
var BlockedAPIKey = modal.Response{Status: http.StatusForbidden, Message: "The API key is blocked", Code: "ERRGEN010"}
var VersionUpdate = modal.Response{Status: http.StatusUpgradeRequired, Message: "Kindly Update to the latest version", Code: "ERRGEN011"}
var GenderUnspecified = modal.Response{Status: http.StatusForbidden, Message: "Gender Unspecified", Code: "ERRGEN012"}

var UserIsBlocked = modal.Response{Status: http.StatusForbidden, Message: "User is unauthorized to access", Code: "ERRUSR010"}
