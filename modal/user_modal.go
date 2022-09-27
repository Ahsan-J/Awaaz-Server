package modal

import (
	"github.com/teris-io/shortid"
	"github.com/jmoiron/sqlx"
	"strings"
	"fmt"
)

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
	Status         int         `db:"status" json:"user_status"`
	CreatedAt      string      `db:"created_at" json:"-"`
	UpdatedAt      interface{} `db:"updated_at" json:"-"`
	DeletedAt      interface{} `db:"deleted_at" json:"-"`
}

// Optimize optimizes the Interface data types to their defined data types
func (user *User) Optimize() {
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
	if user.UpdatedAt != nil {
		user.UpdatedAt = string(user.UpdatedAt.([]uint8))
	}
}


// Save entry to database
func (user *User) Save(db *sqlx.DB) {

	if strings.TrimSpace(user.Name) == "" || user.Status == 0 {
		fmt.Println("Missing Fields from users")
		return
	}

	// Validating Gender
	if strings.ToLower(user.Gender) != "male" && strings.ToLower(user.Gender) != "female" && strings.ToLower(user.Gender) != "other" {
		fmt.Println("Gender unspecified")
		return
	}
	
	if  strings.TrimSpace(user.ID) == "" {
		user.ID, _ = shortid.Generate()
	}
	
	sql := "INSERT INTO user(id,name, password, social_platform, gender,date_of_birth, email, contact, status, created_at, updated_at, deleted_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := db.Exec(sql,
		user.ID,
		user.Name,
		user.Password,
		user.SocialPlatform,
		user.Gender,
		user.DateOfBirth,
		user.Email,
		user.Contact,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	)
	if err != nil {
		fmt.Println(err)
	}
}

// Update the database entry
func (user *User) Update(db *sqlx.DB) {
	
	if strings.TrimSpace(user.Name) == "" || user.Status == 0 {
		fmt.Println("Missing Fields from users")
		return
	}

	// Validating Gender
	if strings.ToLower(user.Gender) != "male" && strings.ToLower(user.Gender) != "female" && strings.ToLower(user.Gender) != "other" {
		fmt.Println("Gender unspecified")
		return
	}
	
	if strings.TrimSpace(user.ID) == "" {
		fmt.Println("Empty ID")
		return
	}

	sql := "UPDATE user SET name=?,password=?,social_platform=?,gender=?,date_of_birth=?,email=?,contact=?,status=?,created_at=?,updated_at=?,deleted_at=? WHERE id='" + user.ID + "'"

	_, err := db.Exec(sql,
		user.Name,
		user.Password,
		user.SocialPlatform,
		user.Gender,
		user.DateOfBirth,
		user.Email,
		user.Contact,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	)
	if err != nil {
		fmt.Println(err)
	}
}
