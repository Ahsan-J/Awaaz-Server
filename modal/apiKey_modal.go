package modal
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
