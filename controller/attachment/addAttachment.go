package attachment

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/modal"
	"awaaz_go_server/responses"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/teris-io/shortid"
)

// AddAttachment adds the attachment to the database
func AddAttachment(w http.ResponseWriter, r *http.Request) {
	var UserID = r.FormValue("id")
	var CaseID = r.FormValue("caseId")

	if strings.TrimSpace(UserID) == "" || strings.TrimSpace(CaseID) == "" {
		responses.FieldMissing.SendAPI(w, nil)
		return
	}

	var Buf bytes.Buffer
	defer Buf.Reset()

	file, header, err := r.FormFile("file")
	if err != nil {
		responses.InternalServerError.SendAPI(w, err)
		return
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")

	name[0] = strings.TrimSpace(strings.Replace(name[0], " ", "", -1))
	name[1] = strings.TrimSpace(name[1])
	fileProcessedName := UserID + "_" + CaseID + "_" + name[0] + "." + name[1]

	io.Copy(&Buf, file)
	err = ioutil.WriteFile("./uploads/"+fileProcessedName, Buf.Bytes(), 0777)
	if err != nil {
		fmt.Println("Error uploading", err)
	}

	db := helpers.GetDBInstance()
	defer db.Close()

	newID, _ := shortid.Generate()
	now := time.Now().Unix()
	var attachment = modal.Attachment{
		ID:         newID,
		CaseID:     CaseID,
		FileName:   fileProcessedName,
		HostedURL:  "https://" + r.Host + "/uploads/" + fileProcessedName,
		MemeType:   header.Header["Content-Type"][0],
		UploadedAt: strconv.Itoa(int(now)),
	}
	attachment.Save(db)
	responses.AttachmentSuccess.SendAPI(w, attachment)
	return
}
