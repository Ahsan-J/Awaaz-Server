package modal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
)

// Victim table modal definition to be matched up with the Database"
type Victim struct {
	ID        string      `db:"id" json:"victim_id"`
	CaseID    string      `db:"case_id" json:"case_id"`
	Name      interface{} `db:"victim_name" json:"victim_name"`
	Gender    string      `db:"gender" json:"gender"`
	Age       interface{} `db:"age" json:"age"`
	Status    int         `db:"status" json:"status"`
	CreatedAt string      `db:"created_at" json:"-"`
	UpdatedAt interface{} `db:"updated_at" json:"-"`
	DeletedAt interface{} `db:"deleted_at" json:"-"`
}

// Optimize the nil values
func (victim *Victim) Optimize() {
	if victim.Name != nil {
		victim.Name = string(victim.Name.([]uint8))
	}
	if victim.Age != nil {
		victim.Age, _ = strconv.Atoi(string(victim.Age.([]uint8)))
	}
	if victim.UpdatedAt != nil {
		victim.UpdatedAt = string(victim.UpdatedAt.([]uint8))
	}
	if victim.DeletedAt != nil {
		victim.DeletedAt = string(victim.DeletedAt.([]uint8))
	}
}

// Save entry to database
func (victim *Victim) Save(db *sqlx.DB) {
	if strings.TrimSpace(victim.ID) == "" {
		victim.ID, _ = shortid.Generate()
	}

	if strings.TrimSpace(victim.CaseID) == "" {
		fmt.Println("Case ID missing")
		return
	}

	if strings.TrimSpace(victim.Gender) != "male" || strings.TrimSpace(victim.Gender) != "female" || strings.TrimSpace(victim.Gender) != "other" {
		fmt.Println("Invalid gender value")
		return
	}

	if strings.TrimSpace(victim.CreatedAt) == "" {
		now := time.Now().Unix()
		victim.CreatedAt = strconv.Itoa(int(now))
	}

	sql := "INSERT INTO victim(id,case_id, victim_name, gender,age, status, created_at, updated_at, deleted_at) VALUES (?,?,?,?,?,?)"

	_, err := db.Exec(sql,
		victim.ID,
		victim.CaseID,
		victim.Name,
		victim.Gender,
		victim.Age,
		victim.Status,
		victim.CreatedAt,
		victim.UpdatedAt,
		victim.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error exec DB", err)
	}
}
