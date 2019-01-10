package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"awaaz_go_server/modal"
	"github.com/teris-io/shortid"
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
func GetSecuredHash(text string) string {
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
	now := time.Now().Unix()
	db := GetDBInstance()
	defer db.Close()
	k, _ := base64.StdEncoding.DecodeString(secretkey)
	c, _ := aes.NewCipher(k)
	gcm, _ := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	id, _ := shortid.Generate()

	accTokenString := "Awaaz-App.AccessToken." + id + "." + user.ID + "." + key.APIKey + "." + key.Platform + "." + deviceID + "." + macAddress + "." + strconv.Itoa(accessTokenTime) + "." + strconv.Itoa(int(now))
	refTokenString := "Awaaz-App.RefreshToken." + id + "." + user.ID + "." + key.APIKey + "." + key.Platform + "." + deviceID + "." + macAddress + "." + strconv.Itoa(refreshTokenTime) + "." + strconv.Itoa(int(now))

	token := modal.Token{
		ID:                 id,
		AccessToken:        hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(accTokenString), nil)),
		AccessTokenExpiry:  accessTokenTime,
		AccessTokenTime:    strconv.Itoa(int(now)),
		RefreshToken:       hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(refTokenString), nil)),
		RefreshTokenTime:   strconv.Itoa(int(now)),
		RefreshTokenExpiry: refreshTokenTime,
		UserID:             user.ID,
		Status:             1,
		APIKey:             key.APIKey,
		DeviceID:           deviceID,
		MacAddress:         macAddress,
	}

	SaveByID(&token, false)
	return token
}

// GetRefreshedToken generates the updated token from refresh token
func GetRefreshedToken(refreshToken string, user modal.User) modal.Token {
	now := time.Now().Unix()
	k, _ := base64.StdEncoding.DecodeString(secretkey)
	t, _ := hex.DecodeString(refreshToken)
	c, _ := aes.NewCipher(k)
	gcm, err := cipher.NewGCM(c)
	nonceSize := gcm.NonceSize()
	if len(t) < nonceSize {
		fmt.Println(err)
	}
	nonce, t := t[:nonceSize], t[nonceSize:]
	text, err := gcm.Open(nil, nonce, t, nil)
	if err != nil {
		fmt.Println(err)
	}

	newNonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, newNonce); err != nil {
		fmt.Println(err)
	}

	s := strings.Split(string(text), ".")

	accTokenString := "Awaaz-App.AccessToken." + s[2] + "." + s[3] + "." + s[4] + "." + s[5] + "." + s[6] + "." + s[7] + "." + strconv.Itoa(accessTokenTime) + "." + strconv.Itoa(int(now))
	refTokenString := "Awaaz-App.RefreshToken." + s[2] + "." + s[3] + "." + s[4] + "." + s[5] + "." + s[6] + "." + s[7] + "." + strconv.Itoa(refreshTokenTime) + "." + strconv.Itoa(int(now))

	db := GetDBInstance()
	defer db.Close()

	apiKeyData := []modal.APIKey{}
	err = db.Select(&apiKeyData, "SELECT * FROM api_keys WHERE api_key='"+s[4]+"'")
	if err != nil {
		fmt.Println("API Key fetching error", err)
	}

	if len(apiKeyData) <= 0 || apiKeyData[0].APIKey != s[4] {
		fmt.Println("APIKey Invalid")
	}

	var token = modal.Token{
		ID:                 s[2],
		AccessToken:        hex.EncodeToString(gcm.Seal(newNonce, newNonce, []byte(accTokenString), nil)),
		AccessTokenExpiry:  accessTokenTime,
		AccessTokenTime:    strconv.Itoa(int(now)),
		RefreshToken:       hex.EncodeToString(gcm.Seal(newNonce, newNonce, []byte(refTokenString), nil)),
		RefreshTokenTime:   strconv.Itoa(int(now)),
		RefreshTokenExpiry: refreshTokenTime,
		UserID:             user.ID,
		Status:             1,
		APIKey:             s[4],
		DeviceID:           s[6],
		MacAddress:         s[7],
	}

	SaveByID(&token, true)
	return token
}

// SaveByID saves token by its ID
func SaveByID(token *modal.Token, update bool) {

	db := GetDBInstance()
	defer db.Close()

	if strings.TrimSpace(token.ID) == "" {
		fmt.Println("Empty ID")
		return
	}

	sql := "INSERT INTO tokens(id,access_token, access_token_time, access_token_expiry, user_id, status, api_key, refresh_token, refresh_token_time, refresh_token_expiry, device_id, mac_address) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	if update == true {
		sql = "UPDATE tokens SET id=?,access_token=?,access_token_time=?,access_token_expiry=?,user_id=?,status=?,api_key=?,refresh_token=?,refresh_token_time=?,refresh_token_expiry=?,device_id=?,mac_address=? WHERE id='" + token.ID + "'"
	}

	_, err := db.Exec(sql,
		token.ID,
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
}
