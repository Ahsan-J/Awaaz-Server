package responses

import (
	"net/http"
	"awaaz_go_server/modal"
)
// Success
var AttachmentSuccess = modal.Response{ Status: http.StatusOK , Message: "Successfully added the attachment", Code:"SUCCEVD001"}

// Errors
var RequiredCaseID = modal.Response{ Status: http.StatusBadRequest, Message: "Missing Case ID", Code: "ERREVD001"}