package attachment

// Attachment table modal definition to be matched up with the Database"
type Attachment struct {
	ID        string `db:"id"`
	CaseID    string `db:"case_id"`
	Name      string `db:"name"`
	MemeType  string `db:"meme_type"`
	HostedURL string `db:"hosted_url"`
}