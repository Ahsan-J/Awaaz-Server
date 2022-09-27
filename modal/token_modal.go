package modal

import (
	"github.com/teris-io/shortid"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Token table modal definition to be matched up with the Database"
type Token struct {
	ID                 string `db:"id" json:"-"`
	AccessToken        string `db:"access_token" json:"access_token"`
	AccessTokenTime    string `db:"access_token_time" json:"access_token_time"`
	AccessTokenExpiry  int    `db:"access_token_expiry" json:"access_token_expiry"`
	RefreshToken       string `db:"refresh_token" json:"refresh_token"`
	RefreshTokenTime   string `db:"refresh_token_time" json:"refresh_token_time"`
	RefreshTokenExpiry int    `db:"refresh_token_expiry" json:"refresh_token_expiry"`
	UserID             string `db:"user_id" json:"-"`
	Status             int    `db:"status" json:"token_status"`
	APIKey             string `db:"api_key" json:"-"`
	DeviceID           string `db:"device_id" json:"-"`
	MacAddress         string `db:"mac_address" json:"-"`
}

// Save entry to database
func (token *Token) Save(db *sqlx.DB) {

	if strings.TrimSpace(token.AccessToken) == "" || strings.TrimSpace(token.RefreshToken) == "" || token.RefreshTokenExpiry == 0 || token.AccessTokenExpiry == 0 || strings.TrimSpace(token.DeviceID) == "" || strings.TrimSpace(token.MacAddress) == "" || strings.TrimSpace(token.APIKey) == "" || token.Status == 0{
		fmt.Println("Missing Fields from tokens")
		return
	}
	
	if  strings.TrimSpace(token.ID) == "" {
		token.ID, _ = shortid.Generate()
	}
	
	sql := "INSERT INTO tokens(id,access_token, access_token_time, access_token_expiry, user_id, status, api_key, refresh_token, refresh_token_time, refresh_token_expiry, device_id, mac_address) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

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

// Update the database entry
func (token *Token) Update(db *sqlx.DB) {
	
	if strings.TrimSpace(token.AccessToken) == "" || strings.TrimSpace(token.RefreshToken) == "" || token.RefreshTokenExpiry == 0 || token.AccessTokenExpiry == 0 || strings.TrimSpace(token.DeviceID) == "" || strings.TrimSpace(token.MacAddress) == "" || strings.TrimSpace(token.APIKey) == "" || token.Status == 0{
		fmt.Println("Missing Fields from tokens")
		return
	}
	
	if strings.TrimSpace(token.ID) == "" {
		fmt.Println("Empty ID")
		return
	}

	sql := "UPDATE tokens SET access_token=?,access_token_time=?,access_token_expiry=?,user_id=?,status=?,api_key=?,refresh_token=?,refresh_token_time=?,refresh_token_expiry=?,device_id=?,mac_address=? WHERE id='" + token.ID + "'"

	_, err := db.Exec(sql,
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
