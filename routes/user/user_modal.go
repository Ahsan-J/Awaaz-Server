package user

// User table modal definition to be matched up with the Database"
type User struct {
	ID             string      `db:"id" json:"id"`
	Name           string      `db:"name" json:"name"`
	Password       interface{} `db:"password" json:"-"`
	SocialPlatform interface{} `db:"social_platform" json:"social_platform"`
	Gender         string      `db:"gender" json:"gender"`
	DateOfBirth    interface{} `db:"date_of_birth" json:"date_of_birth"`
	Email          interface{} `db:"email" json:"email"`
	Contact        interface{} `db:"contact" json:"contact"`
	DeviceID       string      `db:"device_id" json:"devce_id"`
	MacAddress     string      `db:"mac_address" json:"mac_address"`
	CreatedAt      string      `db:"created_at" json:"created_at"`
	UpdateddAt     string      `db:"updated_at" json:"updated_at"`
	DeletedAt      interface{} `db:"deleted_at" json:"deleted_at"`
}

func (user *User) optimize () {
	if user.Password != nil {
		user.Password = string(user.Password.([]uint8))
	}
	if user.SocialPlatform != nil {
		user.SocialPlatform = string(user.SocialPlatform.([]uint8))
	}
	if user.DateOfBirth != nil {
		user.DateOfBirth = string(user.DateOfBirth.([]uint8))
	}
	if user.Email != nil {
		user.Email = string(user.Email.([]uint8))
	}
	if user.Contact != nil {
		user.Contact = string(user.Contact.([]uint8))
	}
	if user.DeletedAt != nil {
		user.DeletedAt = string(user.DeletedAt.([]uint8))
	}
}