package report

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"fmt" // "encoding/json"
	"net/http"
	"awaaz_go_server/responses"
)

// GetAll path:/reports
func GetAll(w http.ResponseWriter, r *http.Request) {
	var err error
	reports := []modal.Report{}
	db := helpers.GetDBInstance()
	defer db.Close()
	// Fetching Data
	err = db.Select(&reports, "SELECT * FROM report")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(reports) <= 0 {
		responses.ReportsNotFound.SendAPI(w, nil)
		return
	}

	for i := range reports {
		reports[i].Optimize()
	}

	// Sending Response
	responses.ReportsSuccess.SendAPI(w, reports)
}
