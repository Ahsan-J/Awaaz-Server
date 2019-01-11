package responses

import (
	"awaaz_go_server/modal"
	"net/http"
)
// Success
var ReportsSuccess = modal.Response{ Status: http.StatusOK, Message: "Successfully fetched all reports", Code: "SUCCREP001"}

// Error
var ReportsNotFound = modal.Response{ Status : http.StatusNotFound, Message: "No reports found", Code: "ERRREP001"}