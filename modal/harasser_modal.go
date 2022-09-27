package modal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
)

// Harasser table modal definition to be matched up with the Database"
type Harasser struct {
	ID          string      `db:"id" json:"harasser_id"`
	CaseID      string      `db:"case_id" json:"case_id"`
	Name        interface{} `db:"harasser_name" json:"harasser_name"`
	Gender      string      `db:"gender" json:"gender"`
	Age         interface{} `db:"age" json:"age"`
	Description string      `db:"harasser_description" json:"harasser_description"`
	Relation    interface{} `db:"relation" json:"relation"`
	Frequency   interface{}	`db:"frequency" json:"frequency"`
	Status      int         `db:"status" json:"status"`
	CreatedAt   string      `db:"created_at" json:"-"`
	UpdatedAt   interface{} `db:"updated_at" json:"-"`
	DeletedAt   interface{} `db:"deleted_at" json:"-"`
}

// Optimize the nil values
func (harasser *Harasser) Optimize() {
	if harasser.Name != nil {
		harasser.Name = string(harasser.Name.([]uint8))
	}
	if harasser.Age != nil {
		harasser.Age = string(harasser.Age.([]uint8))
	}
	if harasser.Relation != nil {
		harasser.Relation = string(harasser.Relation.([]uint8))
	}
	if harasser.Frequency != nil {
		harasser.Frequency, _ = strconv.Atoi(string(harasser.Frequency.([]uint8)))
	}
	if harasser.UpdatedAt != nil {
		harasser.UpdatedAt = string(harasser.UpdatedAt.([]uint8))
	}
	if harasser.DeletedAt != nil {
		harasser.DeletedAt = string(harasser.DeletedAt.([]uint8))
	}
}

// Save entry to database
func (harasser *Harasser) Save(db *sqlx.DB) {

	if strings.TrimSpace(harasser.CaseID) == "" {
		fmt.Println("Case ID missing")
		return
	}

	if strings.TrimSpace(harasser.Description) == "" {
		fmt.Println("Missing Required Fields")
		return
	}

	if strings.TrimSpace(harasser.Gender) != "male" || strings.TrimSpace(harasser.Gender) != "female" || strings.TrimSpace(harasser.Gender) != "other" {
		fmt.Println("Invalid gender value")
		return
	}

	if strings.TrimSpace(harasser.ID) == "" {
		harasser.ID, _ = shortid.Generate()
	}

	if strings.TrimSpace(harasser.CreatedAt) == "" {
		now := time.Now().Unix()
		harasser.CreatedAt = strconv.Itoa(int(now))
	}

	sql := "INSERT INTO harasser(id,case_id, harasser_name, gender,age,harasser_description,relation,frequency, status, created_at, updated_at, deleted_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := db.Exec(sql,
		harasser.ID,
		harasser.CaseID,
		harasser.Name,
		harasser.Gender,
		harasser.Age,
		harasser.Description,
		harasser.Relation,
		harasser.Frequency,
		harasser.Status,
		harasser.CreatedAt,
		harasser.UpdatedAt,
		harasser.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error exec DB", err)
	}
}
