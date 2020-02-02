package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anatolyefimov/cf-cli/user"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	args := os.Args[1:]
	if args[0] == "login" {
		fmt.Println("Enter handle:")
		var handle string
		fmt.Scan(&handle)
		fmt.Println("Enter password:")
		bytePassword, err := terminal.ReadPassword(0)

		if err == nil {
			user := user.New(handle, string(bytePassword))
			user.Login()
			resp, err := user.Client.Get("https://codeforces.com")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			html, _ := ioutil.ReadAll(resp.Body)
			if user.IsLoggedIn(html) {
				fmt.Printf("You are logged in as %s", handle)
			}
		}
	}
	// csrf, err = findCsrf(body)
	// if err != nil {
	// 	return
	// }
	// resp, err = client.PostForm("https://codeforces.com/problemset/submit?csrf_token="+csrf, url.Values{
	// 	"csrf_token":           {csrf},
	// 	"ftaa":                 {ftaa},
	// 	"bfaa":                 {bfaa},
	// 	"action":               {"submitSolutionFormSubmitted"},
	// 	"submittedProblemCode": {"33A"},
	// 	"programTypeId":        {"50"},
	// 	"source":               {"101010010101010100101010101032423451010"},
	// 	"tabSize":              {"4"},
	// 	"sourceFile":           {},
	// 	"_tta":                 {"434"},
	// })
	// if err != nil {
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err = ioutil.ReadAll(resp.Body) //nolint
	// fmt.Print(string(body))
}
