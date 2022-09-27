package generalapi

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"github.com/teris-io/shortid"
	"awaaz_go_server/responses"
)

// Register path:/register
func Register(w http.ResponseWriter, r *http.Request) {
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
		responses.FieldMissing.SendAPI(w, nil)
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
			responses.PasswordInvalid.SendAPI(w, nil)
			return
		}

		if password != confirmPassword {
			responses.PasswordMismatch.SendAPI(w, nil)
			return
		}

		password = helpers.GetSecuredHash(password)
	}
	// Validating Email
	if len(email) > 0 && helpers.IsApproved(status) == true && helpers.ValidateEmail(email) == false {
		responses.EmailInvalid.SendAPI(w, nil)
		return
	}
	// Validating Gender
	if strings.ToLower(gender) != "male" && strings.ToLower(gender) != "female" && strings.ToLower(gender) != "other" {
		responses.GenderUnspecified.SendAPI(w, nil)
		return
	}

	db := helpers.GetDBInstance()
	defer db.Close()

	// Validating API key
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

	newID, err := shortid.Generate()
	if err != nil {
		fmt.Println("Error generating ID", err)
	}

	now := time.Now().Unix()
	var user = modal.User{
		ID:             newID,
		Name:           name,
		Password:       password,
		SocialPlatform: socialPlatform,
		Gender:         gender,
		DateOfBirth:    dateOfBirth,
		Email:          email,
		Contact:        contact,
		Status:         status,
		CreatedAt:      strconv.Itoa(int(now)),
		UpdatedAt:      nil,
		DeletedAt:      nil,
	}
	
	user.Save(db);

	token := helpers.CreateNewToken(user, apiKeyData[0], deviceID, macAddress)
	res := combinedResponse{token, user}
	responses.RegisterSuccess.SendAPI(w, res)
}