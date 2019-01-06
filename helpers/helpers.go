package helpers

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strconv"
	"time"
	"encoding/base64"
	"fmt"

	"../modal"
)

// ParseTimestamp parses the String representation for unix timestamp and returns Time
func ParseTimestamp(timestamp string) time.Time {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}

// GetSecuredHash generates and returns the SHA256 hashcode of the string
func GetSecuredHash(text string) string{
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// IsApproved Checks the Approved status
func IsApproved(status int) bool {
	if status&approved == 1 {
		return true
	}
	return false
}

// IsBlocked Checks the Blocked status
func IsBlocked(status int) bool {
	if status&blocked == 1 {
		return true
	}
	return false
}

// IsPending Checks the Pending status
func IsPending(status int) bool {
	if status&pending == 1 {
		return true
	}
	return false
}

// IsAnonymous Checks the Anonymous status
func IsAnonymous(status int) bool {
	if status&anonymous == 1 {
		return true
	}
	return false
}

// ValidateEmail uses regex to validate email address
func ValidateEmail(email string) bool {
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

// ValidatePassword validates password pattern for user
func ValidatePassword(password string) bool {
	match, _ := regexp.MatchString(passwordPattern, password)
	return match
}

// CreateNewToken for the user session
func CreateNewToken(user modal.User, key modal.APIKey, deviceID string, macAddress string) modal.Token {
	now := time.Now().Unix();
	db := GetDBInstance()
	defer db.Close()
	
	k,_ := base64.StdEncoding.DecodeString(secretkey)
	enc, _ := aes.NewCipher([]byte(string(k)))

	accTokenString := "Awaaz-App.AccessToken." + user.ID + "." + key.APIKey + "." + key.Version + "." + key.Platform + "." + "." + deviceID + "." + macAddress + "." + string(accessTokenTime) + "." + strconv.Itoa(int(now))
	refTokenString := "Awaaz-App.RefreshToken." + user.ID + "." + key.APIKey + "." + key.Version + "." + key.Platform + "." + "." + deviceID + "." + macAddress + "." + string(refreshTokenTime) + "." + strconv.Itoa(int(now));
	var t = make([]byte,16)
	var r = make([]byte,16)
	enc.Encrypt(t,[]byte(accTokenString))
	enc.Encrypt(r,[]byte(refTokenString))
		
	token := modal.Token{
		AccessToken : hex.EncodeToString(t[:]),
		AccessTokenExpiry:accessTokenTime,
		AccessTokenTime : strconv.Itoa(int(now)),
		RefreshToken : hex.EncodeToString(r[:]),
		RefreshTokenTime : strconv.Itoa(int(now)),
		RefreshTokenExpiry : refreshTokenTime,
		UserID : user.ID,
		Status : 1,
		APIKey : key.APIKey,
		DeviceID: deviceID, 
		MacAddress: macAddress, 
	}

	sql := "INSERT INTO tokens(access_token, access_token_time, access_token_expiry, user_id, status, api_key, refresh_token, refresh_token_time, refresh_token_expiry, device_id, mac_address) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	_ , err := db.Exec(sql,
		token.AccessToken,
		token.AccessTokenTime,
		token.AccessTokenExpiry,
		token.UserID,
		token.Status,
		token.APIKey,
		token.RefreshToken,
		token.RefreshTokenTime,
		token.RefreshTokenExpiry,
		token.DeviceID,
		token.MacAddress,
	)
	if err != nil {
		fmt.Println(err)
	}
	return token
}
