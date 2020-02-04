package user

import (
	"net/http"
	"github.com/anatolyefimov/cf-cli/cookiejar"
)

//User class
type User struct {
	Handle    string         `json:"handle"`
	Password  string         `json:"password"`
	Ftaa      string         `json:"ftaa"`
	Bfaa      string         `json:"bfaa"`
	Client    *http.Client   `json:"-"`
	CookieJar *cookiejar.Jar `json:"cookie"`
}

//New User
func New(handle string, password string) User {
	user := User{
		Handle:   handle,
		Password: password,
	}
	user.CookieJar, _ = cookiejar.New(nil)
	user.Client = &http.Client{Jar: user.CookieJar}

	return user
}
