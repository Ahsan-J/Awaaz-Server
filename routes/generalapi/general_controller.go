package generalapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"strconv"
	"github.com/teris-io/shortid"

	"../../helpers"
	"../../modal"
)

type combinedResponse struct {
	modal.Token
	modal.User
}

func login(w http.ResponseWriter, r *http.Request) {
	var email = strings.TrimSpace(r.FormValue("email"))
	var password = strings.TrimSpace(r.FormValue("password"))
	var platform = strings.TrimSpace(r.FormValue("platform"))
	var deviceID = strings.TrimSpace(r.FormValue("deviceId"))
	var macAddress = strings.TrimSpace(r.FormValue("macAddress"))
	var apiKey = strings.TrimSpace(r.FormValue("apiKey"))
	var socialPlatform = strings.TrimSpace(r.FormValue("socialPlatform"))
	var version = strings.TrimSpace(r.FormValue("version"))

	if len(version) <= 0 || len(platform) <= 0 || len(deviceID) <= 0 || len(deviceID) <= 0 || len(macAddress) <= 0 || len(apiKey) <= 0 {
		fieldMissing.SendAPI(w, nil)
		return
	}

	if len(socialPlatform) > 0 {
		fmt.Println("Signing from Social platform: " + socialPlatform)
	}

	if len(email) <= 0 {
		emailRequired.SendAPI(w, nil)
		return
	}

	if helpers.ValidateEmail(email) == false {
		emailInvalid.SendAPI(w, nil)
		return
	}

	db := helpers.GetDBInstance()
	defer db.Close()

	apiKeyData := []modal.APIKey{}
	err := db.Select(&apiKeyData, "SELECT * FROM api_keys WHERE api_key='"+apiKey+"'")
	if err != nil {
		fmt.Println("API Key fetching error", err)
		internalServerError.SendAPI(w, nil)
		return
	}

	if len(apiKeyData) <= 0 || apiKeyData[0].APIKey != apiKey || apiKeyData[0].Platform != platform || helpers.IsApproved(apiKeyData[0].Status) == false {
		invalidAPIKey.SendAPI(w, nil)
		return
	}

	if apiKeyData[0].Version != version {
		versionUpdate.SendAPI(w, nil)
		return
	}

	var user = []modal.User{}

	err = db.Select(&user, "SELECT * FROM user WHERE email='"+email+"'")
	if err != nil {
		fmt.Println(err)
		internalServerError.SendAPI(w, nil)
		return
	}

	if len(user) <= 0 {
		loginUserNotFound.SendAPI(w, nil)
		return
	}

	user[0].Optimize()

	if user[0].Password != nil || helpers.IsApproved(user[0].Status) == true {
		// approved password valdation user
		if len(password) <= 0 {
			passwordRequired.SendAPI(w, nil)
			return
		}

		if helpers.GetSecuredHash(password) != user[0].Password {
			passwordMismatch.SendAPI(w, nil)
			return
		}
	}

	token := helpers.CreateNewToken(user[0], apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user[0]}

	loginSuccess.SendAPI(w, res)
}

func register(w http.ResponseWriter, r *http.Request) {
	var name = strings.TrimSpace(r.FormValue("name"))
	var gender = strings.TrimSpace(r.FormValue("gender"))
	var dateOfBirth = strings.TrimSpace(r.FormValue("dateOfBirth"))         // optional field
	var email = strings.TrimSpace(r.FormValue("email"))                     // optional field
	var contact = strings.TrimSpace(r.FormValue("contact"))                 // optional field
	var password = strings.TrimSpace(r.FormValue("password"))               // optional field
	var confirmPassword = strings.TrimSpace(r.FormValue("confirmPassword")) // optional field
	var socialPlatform = strings.TrimSpace(r.FormValue("socialPlatform"))   //optional field

	var platform = strings.TrimSpace(r.FormValue("platform"))
	var deviceID = strings.TrimSpace(r.FormValue("deviceId"))
	var macAddress = strings.TrimSpace(r.FormValue("macAddress"))
	var apiKey = strings.TrimSpace(r.FormValue("apiKey"))
	var version = strings.TrimSpace(r.FormValue("version"))
	var status = 1

	if len(version) <= 0 || len(platform) <= 0 || len(deviceID) <= 0 || len(macAddress) <= 0 || len(apiKey) <= 0 || len(name) <= 0 || len(gender) <= 0 {
		fieldMissing.SendAPI(w, nil)
		return
	}

	if len(password) <= 0 || len(email) <= 0 || len(contact) <= 0 || len(dateOfBirth) <= 0 {
		// Setting to be anonymous
		status = 8
	}

	// Password Valdation
	if len(password) > 0 && helpers.IsApproved(status) == true {
		// approved password valdation user
		if helpers.ValidatePassword(password) == false {
			passwordInvalid.SendAPI(w, nil)
			return
		}

		if password != confirmPassword {
			passwordMismatch.SendAPI(w, nil)
			return
		}

		password = helpers.GetSecuredHash(password)
	} 
	// Validating Email
	if len(email) > 0 && helpers.IsApproved(status) == true && helpers.ValidateEmail(email) == false {
		emailInvalid.SendAPI(w, nil)
		return
	}
	// Validating Gender
	if strings.ToLower(gender) != "male" && strings.ToLower(gender) != "female" && strings.ToLower(gender) != "other" {
		genderUnspecified.SendAPI(w, nil)
		return
	}

	db := helpers.GetDBInstance()
	defer db.Close()

	// Validating API key
	apiKeyData := []modal.APIKey{}
	err := db.Select(&apiKeyData, "SELECT * FROM api_keys WHERE api_key='"+apiKey+"'")
	if err != nil {
		fmt.Println("API Key fetching error", err)
		internalServerError.SendAPI(w, nil)
		return
	}

	if len(apiKeyData) <= 0 || apiKeyData[0].APIKey != apiKey || apiKeyData[0].Platform != platform || helpers.IsApproved(apiKeyData[0].Status) == false {
		invalidAPIKey.SendAPI(w, nil)
		return
	}

	if apiKeyData[0].Version != version {
		versionUpdate.SendAPI(w, nil)
		return
	}

	newID,err := shortid.Generate()
	if err != nil {
		fmt.Println("Error generating ID",err)
	}

	now := time.Now().Unix()
	var user = modal.User {
		ID: newID,
		Name: name,
		Password: password,
		SocialPlatform: socialPlatform,
		Gender: gender,
		DateOfBirth: dateOfBirth,
		Email: email,
		Contact: contact,
		Status : status,
		CreatedAt : strconv.Itoa(int(now)),
		UpdatedAt : nil,
		DeletedAt : nil,
	}

	token := helpers.CreateNewToken(user, apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user}

	registerSuccess.SendAPI(w, res)
}
