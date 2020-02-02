package user

import (
	"net/http"
	"net/http/cookiejar"
)

//User class
type User struct {
	handle    string
	password  string
	htaa      string
	bfaa      string
	client    *http.Client
	cookieJar *cookiejar.Jar
}

//New User
func New(handle string, password string) User {
	user := User{
		handle:   handle,
		password: password,
	}
	user.cookieJar, _ = cookiejar.New(nil)
	user.client = &http.Client{Jar: user.cookieJar}

	return user
}
