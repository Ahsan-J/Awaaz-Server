package report

import (
	"awaaz_go_server/responses"
	"net/http"
	"strings"
	// "fmt"
	// "encoding/json"
)

// Submit function to submit the report
func Submit(w http.ResponseWriter, r *http.Request) {
	
	// var email = strings.TrimSpace(r.FormValue("email"))

}

func verifyUser(w http.ResponseWriter, r *http.Request) bool {
	var email = strings.TrimSpace(r.FormValue("email"))
	var name = strings.TrimSpace(r.FormValue("name"))
	var gender = strings.TrimSpace(r.FormValue("gender"))
	var dateOfBirth = strings.TrimSpace(r.FormValue("dateOfBirth"))
	var contact = strings.TrimSpace(r.FormValue("contact"))

	if email == "" || gender == "" || dateOfBirth == "" || contact == "" || name == "" {
		responses.ReportingUserFieldMissing.SendAPI(w, nil)
		return false
	}

	return true
}

func verifyReport(w http.ResponseWriter, r *http.Request) bool {
	var city = strings.TrimSpace(r.FormValue("city"))
	var harassmentType = strings.TrimSpace(r.FormValue("harassmentType"))
	var eventPlace = strings.TrimSpace(r.FormValue("eventPlace"))
	var location = strings.TrimSpace(r.FormValue("location"))
	var dateTime = strings.TrimSpace(r.FormValue("dateTime"))
	var eventDescription = strings.TrimSpace(r.FormValue("eventDescription"))

	if city == "" || harassmentType == "" || eventPlace == "" || location == "" || dateTime == "" || eventDescription == "" {
		responses.ReportingDetailsMissings.SendAPI(w, nil)
		return false
	}

	return true
}

func verifyVictim(w http.ResponseWriter, r *http.Request) bool {
	
	return true
}