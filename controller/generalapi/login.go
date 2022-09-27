package generalapi

import (
	"fmt"
	"net/http"
	"strings"

	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"awaaz_go_server/responses"
)

type combinedResponse struct {
	modal.Token
	modal.User
}

// Login path:/login
func Login(w http.ResponseWriter, r *http.Request) {
	var email = strings.TrimSpace(r.FormValue("email"))
	var password = strings.TrimSpace(r.FormValue("password"))
	var platform = strings.TrimSpace(r.FormValue("platform"))
	var deviceID = strings.TrimSpace(r.FormValue("deviceId"))
	var macAddress = strings.TrimSpace(r.FormValue("macAddress"))
	var apiKey = strings.TrimSpace(r.FormValue("apiKey"))
	var socialPlatform = strings.TrimSpace(r.FormValue("socialPlatform"))
	var version = strings.TrimSpace(r.FormValue("version"))

	if len(version) <= 0 || len(platform) <= 0 || len(deviceID) <= 0 || len(deviceID) <= 0 || len(macAddress) <= 0 || len(apiKey) <= 0 {
		responses.FieldMissing.SendAPI(w, nil)
		return
	}

	if len(socialPlatform) > 0 {
		fmt.Println("Signing from Social platform: " + socialPlatform)
	}

	if len(email) <= 0 {
		responses.EmailRequired.SendAPI(w, nil)
		return
	}

	if helpers.ValidateEmail(email) == false {
		responses.EmailInvalid.SendAPI(w, nil)
		return
	}

	db := helpers.GetDBInstance()
	defer db.Close()

	apiKeyData := []modal.APIKey{}
	err := db.Select(&apiKeyData, "SELECT * FROM api_keys WHERE api_key='"+apiKey+"'")
	if err != nil {
		fmt.Println("API Key fetching error", err)
		responses.InternalServerError.SendAPI(w, nil)
		return
	}

	if len(apiKeyData) <= 0 || apiKeyData[0].APIKey != apiKey || apiKeyData[0].Platform != platform || helpers.IsApproved(apiKeyData[0].Status) == false {
		responses.InvalidAPIKey.SendAPI(w, nil)
		return
	}

	if apiKeyData[0].Version != version {
		responses.VersionUpdate.SendAPI(w, nil)
		return
	}

	var user = []modal.User{}

	err = db.Select(&user, "SELECT * FROM user WHERE email='"+email+"'")
	if err != nil {
		fmt.Println(err)
		responses.InternalServerError.SendAPI(w, nil)
		return
	}

	if len(user) <= 0 {
		responses.LoginUserNotFound.SendAPI(w, nil)
		return
	}

	user[0].Optimize()

	if user[0].Password != nil || helpers.IsApproved(user[0].Status) == true {
		// approved password valdation user
		if len(password) <= 0 {
			responses.PasswordRequired.SendAPI(w, nil)
			return
		}

		if helpers.GetSecuredHash(password) != user[0].Password {
			responses.PasswordMismatch.SendAPI(w, nil)
			return
		}
	}

	token := helpers.CreateNewToken(user[0], apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user[0]}

	responses.LoginSuccess.SendAPI(w, res)
}