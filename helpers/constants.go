package helpers

// DateLayout is the standard layout used throughout the program for Date only
var DateLayout = "2006-01-02"

// DateTimeLayout is the standard layout used throughout the program for DateTime
var DateTimeLayout = ""

// TimeLayout is the standard layout used throughout the program for Time only
var TimeLayout = "03:04 PM"

var accessTokenTime = 900       // 15 minutes in seconds
var refreshTokenTime = 15780000 // 6 months in seconds

// Status to determine data validity
var approved = 1
var pending = 2
var blocked = 4
var anonymous = 8
var admin = 16
var superAdmin = 32

var secretkey = "YXdhYXotdGhlLXJlcG9ydGluZy1hcHBsaWNhdGlvbi4="

// Validation regex patterns
var emailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// Minimum six characters:
var passwordPattern = "([a-zA-Z0-9]+){6,}"

var dbConfig = "xcalaiberz:illusionist@tcp(localhost:3306)/awaaz_case_store" // development
// var dbConfig = "awaaz_admin:qwerty12345@tcp(db4free.net:3306)/awaaz_case_store" // production
