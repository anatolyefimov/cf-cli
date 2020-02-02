package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anatolyefimov/cf-cli/user"
)

// Login command
func Login(handle string, password string) {
	user := user.New(handle, password)
	user.Login()
	resp, err := user.Client.Get("https://codeforces.com")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)
	if user.IsLoggedIn(html) {
		log.Println(fmt.Sprintf("You logged in as %s", handle))
	}
}
