package responses

import (
	"awaaz_go_server/modal"
	"net/http"
)

// Success
var ReportsSuccess = modal.Response{Status: http.StatusOK, Message: "Successfully fetched all reports", Code: "SUCCREP001"}

// Error
var ReportsNotFound = modal.Response{Status: http.StatusNotFound, Message: "No reports found", Code: "ERRREP001"}
var ReportingUserFieldMissing = modal.Response{Status: http.StatusBadGateway, Message: "User Information is incomplete", Code: "ERRREP002"}
var ReportingDetailsMissings = modal.Response{Status: http.StatusBadGateway, Message: "Reporting deials missing", Code: "ERRREP003"}
