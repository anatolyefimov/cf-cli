package user

import (
	"net/http"
	"net/http/cookiejar"
)

//User class
type User struct {
	handle    string
	password  string
	ftaa      string // nolint
	bfaa      string // nolint
	Client    *http.Client
	cookieJar *cookiejar.Jar
}

//New User
func New(handle string, password string) User {
	user := User{
		handle:   handle,
		password: password,
	}
	user.cookieJar, _ = cookiejar.New(nil)
	user.Client = &http.Client{Jar: user.cookieJar}

	return user
}
