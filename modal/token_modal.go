package modal

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