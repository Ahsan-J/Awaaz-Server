package modal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
)

// Witness table modal definition to be matched up with the Database"
type Witness struct {
	ID        string      `db:"id" json:"witness_id"`
	CaseID    string      `db:"case_id" json:"case_id"`
	Name      interface{} `db:"witness_name" json:"witness_name"`
	Statement string      `db:"statement" json:"statement"`
	Gender    string      `db:"gender" json:"gender"`
	Contact   interface{} `db:"contact" json:"contact"`
	Address   interface{} `db:"address" json:"address"`
	Age       interface{} `db:"age" json:"age"`
	Status    int         `db:"status" json:"status"`
	CreatedAt string      `db:"created_at" json:"-"`
	UpdatedAt interface{} `db:"updated_at" json:"-"`
	DeletedAt interface{} `db:"deleted_at" json:"-"`
}

// Optimize the nil values
func (witness *Witness) Optimize() {
	if witness.Name != nil {
		witness.Name = string(witness.Name.([]uint8))
	}
	if witness.Contact != nil {
		witness.Contact = string(witness.Contact.([]uint8))
	}
	if witness.Address != nil {
		witness.Address = string(witness.Address.([]uint8))
	}
	if witness.Age != nil {
		witness.Age, _ = strconv.Atoi(string(witness.Age.([]uint8)))
	}
	if witness.UpdatedAt != nil {
		witness.UpdatedAt = string(witness.UpdatedAt.([]uint8))
	}
	if witness.DeletedAt != nil {
		witness.DeletedAt = string(witness.DeletedAt.([]uint8))
	}
}

// Save entry to database
func (witness *Witness) Save(db *sqlx.DB) {

	if strings.TrimSpace(witness.CaseID) == "" {
		fmt.Println("Case ID missing")
		return
	}

	if strings.TrimSpace(witness.Statement) == "" {
		fmt.Println("Missing Required Fields")
		return
	}

	if strings.TrimSpace(witness.Gender) != "male" || strings.TrimSpace(witness.Gender) != "female" || strings.TrimSpace(witness.Gender) != "other" {
		fmt.Println("Invalid gender value")
		return
	}

	if strings.TrimSpace(witness.ID) == "" {
		witness.ID, _ = shortid.Generate()
	}

	if strings.TrimSpace(witness.CreatedAt) == "" {
		now := time.Now().Unix()
		witness.CreatedAt = strconv.Itoa(int(now))
	}

	sql := "INSERT INTO witness(id,case_id, witness_name,statement, gender,contact,address,age, status, created_at, updated_at, deleted_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := db.Exec(sql,
		witness.ID,
		witness.CaseID,
		witness.Name,
		witness.Statement,
		witness.Gender,
		witness.Contact,
		witness.Address,
		witness.Age,
		witness.Status,
		witness.CreatedAt,
		witness.UpdatedAt,
		witness.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error exec DB", err)
	}
}
