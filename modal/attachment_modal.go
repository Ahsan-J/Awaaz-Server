package modal

import (
	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
	"strings"
	"fmt"
	"time"
	"strconv"
)

// Attachment modal to match attachment table defined in Database
type Attachment struct {
	ID         string `db:"id" json:"id"`
	CaseID     string `db:"case_id" json:"case_id"`
	FileName   string `db:"file_name" json:"file_name"`
	MemeType   string `db:"meme_type" json:"meme_type"`
	HostedURL  string `db:"hosted_url" json:"hosted_url"`
	UploadedAt string `db:"uploaded_at" json:"uploaded_at"`
}
// Save in database
func (fileData * Attachment) Save(db *sqlx.DB) {

	if strings.TrimSpace(fileData.CaseID) == "" || strings.TrimSpace(fileData.FileName) == "" || strings.TrimSpace(fileData.HostedURL) == "" || strings.TrimSpace(fileData.FileName) == "" || strings.TrimSpace(fileData.MemeType) == ""{
		fmt.Println("Missing Fields")
		return
	}

	if strings.TrimSpace(fileData.ID) == "" {
		fileData.ID,_ = shortid.Generate()
	}

	if strings.TrimSpace(fileData.UploadedAt) == "" {
		if strings.TrimSpace(fileData.UploadedAt) == "" {
			now := time.Now().Unix()
			fileData.UploadedAt = strconv.Itoa(int(now))
		}
	} 

	sql := "INSERT INTO attachment(id,case_id, file_name, meme_type, hosted_url, uploaded_at) VALUES (?,?,?,?,?,?)"

	_,err := db.Exec(sql,
		fileData.ID,
		fileData.CaseID,
		fileData.FileName,
		fileData.MemeType,
		fileData.HostedURL,
		fileData.UploadedAt,
	)
	if err != nil {
		fmt.Println("Error inserting",err)
	}
}