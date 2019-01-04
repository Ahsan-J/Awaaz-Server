package user

// User table modal definition to be matched up with the Database"
type User struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	SocialPlatform string `db:"social_platform"`
	Gender         string `db:"gender"`
	DateOfBirth    string `db:"date_of_birth"`
	Email          string `db:"email"`
	Contact        string `db:"contact"`
	DeviceID       string `db:"device_id"`
	MacAddress     string `db:"mac_address"`
	CreatedAt      string `db:"created_at"`
	UpdateddAt     string `db:"updated_at"`
	DeletedAt      string `db:"deleted_at"`
}
