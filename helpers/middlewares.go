package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../modal"
)

var accessTokenExpired = modal.Response{Status: http.StatusForbidden, Message: "Access Token is expired", Code: "ERRAT001"}
var accessTokenInvalid = modal.Response{Status: http.StatusForbidden, Message: "Invalid Access token", Code: "ERRAT002"}
var accessTokenRequired = modal.Response{Status: http.StatusBadRequest, Message: "Access token is required", Code: "ERRAT003"}

// LoggingMiddleware is copied to form a basic structure of middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := []modal.Token{}

		if _, ok := r.Header["Authorization"]; !ok {
			accessTokenRequired.SendAPI(w, nil)
			return
		}

		s := strings.Split(r.Header["Authorization"][0], " ")
		db := GetDBInstance()
		defer db.Close()
		err := db.Select(&token, "SELECT * FROM tokens WHERE access_token='"+s[1]+"'")
		if err != nil {
			fmt.Println("db error", err)
		}

		if len(token) <= 0 {
			accessTokenInvalid.SendAPI(w, nil)
			return
		}

		ATime := ParseTimestamp(token[0].AccessTokenTime)
		accessDuration, _ := time.ParseDuration(strconv.Itoa(token[0].AccessTokenExpiry) + "s")

		if ATime.Add(accessDuration).Unix() < time.Now().Unix() {
			accessTokenExpired.SendAPI(w, nil)
			return
		}

		// accessTokenExpired.SendAPI(w,token);
		next.ServeHTTP(w, r)
	})
}

// SetHeaders injects the necessary header contents for an API call
func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		next.ServeHTTP(w, r)
	})
}
