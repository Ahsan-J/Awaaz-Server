package modal

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"github.com/jmoiron/sqlx"
)

// APIKey modal to match API_KEYS table defined in Database
type APIKey struct {
	APIKey string `db:"api_key" json:"api"`
	Version string `db:"version" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
	DeletedAt interface{} `db:"deleted_at" json:"-"`
	UpdatedAt interface{} `db:"updated_at" json:"-"`
	Platform string `db:"platform" json:"-"`
	Status int `db:"status" json:"-"`
}
// Optimize optimizes the Interface data types to their defined data types
func (key *APIKey) Optimize() {
	if key.DeletedAt != nil {
		key.DeletedAt = string(key.DeletedAt.([]uint8))
	}
	if key.UpdatedAt != nil {
		key.UpdatedAt = string(key.UpdatedAt.([]uint8))
	}
}

// Save entry to database
func (key *APIKey) Save(db *sqlx.DB) {
	if strings.TrimSpace(key.APIKey) == "" || strings.TrimSpace(key.Version) == "" || strings.TrimSpace(key.Platform) == "" || key.Status == 0 {
		fmt.Println("Entry could not be saved", key)
		return
	}
	if strings.TrimSpace(key.CreatedAt) == "" {
	  now := time.Now().Unix()
		key.CreatedAt = strconv.Itoa(int(now))
	}
	sql := "INSERT INTO tokens(api_key,version,platform, status, created_at, updated_at,deleted_at) VALUES (?,?,?,?,?,?,?)"

	_,err := db.Exec(sql,
		key.APIKey,
		key.Version,
		key.Platform,
		key.Status,
		key.CreatedAt,
		key.UpdatedAt,
		key.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error from API Key",err)
	}
}
