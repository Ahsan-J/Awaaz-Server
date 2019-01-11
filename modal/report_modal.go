package modal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// Report modal to match report table defined in Database
type Report struct {
	ID               string      `db:"case_id" json:"case_id"`
	UserID           string      `db:"user_id" json:"-"`
	DateTime         string      `db:"date_time" json:"date_time"`
	City             string      `db:"city" json:"city"`
	Location         string      `db:"location" json:"location"`
	EventPlace       string      `db:"event_place" json:"event_place"`
	HarassmentType   string      `db:"harassment_type" json:"harassment-type"`
	EventDescription string      `db:"event_description" json:"event_description"`
	Status           int         `db:"status" json:"case_status"`
	CreatedAt        string      `db:"created_at" json:"-"`
	UpdatedAt        interface{} `db:"updated_at" json:"-"`
	DeletedAt        interface{} `db:"deleted_at" json:"-"`
}

// Optimize optimizes the Interface data types to their defined data types
func (report *Report) Optimize() {
	if report.DeletedAt != nil {
		report.DeletedAt = string(report.DeletedAt.([]uint8))
	}
	if report.UpdatedAt != nil {
		report.UpdatedAt = string(report.UpdatedAt.([]uint8))
	}
}

// Save entry to database
func (report *Report) Save(db *sqlx.DB) {
	if strings.TrimSpace(report.DateTime) == "" || strings.TrimSpace(report.City) == "" || strings.TrimSpace(report.Location) == "" || strings.TrimSpace(report.EventPlace) == "" || strings.TrimSpace(report.EventPlace) == "" || strings.TrimSpace(report.EventDescription) == "" {
		fmt.Println("Entry could not be saved", report)
		return
	}
	if strings.TrimSpace(report.CreatedAt) == "" {
		now := time.Now().Unix()
		report.CreatedAt = strconv.Itoa(int(now))
	}

	sql := "INSERT INTO report(id,user_id,date_time, city, location, event_place,harassment_type,event_description,status,created_at,updated_at,deleted_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := db.Exec(sql,
		report.ID,
		report.UserID,
		report.DateTime,
		report.City,
		report.Location,
		report.EventPlace,
		report.HarassmentType,
		report.EventDescription,
		report.Status,
		report.CreatedAt,
		report.UpdatedAt,
		report.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error from Report", err)
	}
}

// Update entry to database
func (report *Report) Update(db *sqlx.DB) {
	if strings.TrimSpace(report.DateTime) == "" || strings.TrimSpace(report.City) == "" || strings.TrimSpace(report.Location) == "" || strings.TrimSpace(report.EventPlace) == "" || strings.TrimSpace(report.EventPlace) == "" || strings.TrimSpace(report.EventDescription) == "" {
		fmt.Println("Entry could not be saved", report)
		return
	}

	if strings.TrimSpace(report.ID) == "" {
		fmt.Println("Missing Case ID")
		return
	}

	sql := "UPDATE report SET user_id=?,date_time=?, city=?, location=?, event_place=?,harassment_type=?,event_description=?,status=?,created_at=?,updated_at=?,deleted_at=? WHERE id='" + report.ID + "'"

	_, err := db.Exec(sql,
		report.ID,
		report.UserID,
		report.DateTime,
		report.City,
		report.Location,
		report.EventPlace,
		report.HarassmentType,
		report.EventDescription,
		report.Status,
		report.CreatedAt,
		report.UpdatedAt,
		report.DeletedAt,
	)
	if err != nil {
		fmt.Println("Error from Report", err)
	}
}
