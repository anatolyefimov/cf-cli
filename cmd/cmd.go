package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anatolyefimov/cf-cli/user"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Login command
func Login() {
	fmt.Println("Enter handle:")
	var handle string
	fmt.Scan(&handle)
	fmt.Println("Enter password:")
	bytePassword, _ := terminal.ReadPassword(0)
	password := string(bytePassword)
	user := user.New(handle, password)
	user.Login()
}

//Submit to archive
func Submit(problemId string, source string) {
	user := user.New("", "")
	user.Fetch()
	if user.IsLoggedIn() {
		err := os.Chdir(".")
		if err != nil {
			log.Fatalln(err)
		}
		file, err := os.Open(source)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		sourceCode, err := ioutil.ReadAll(file)

		if err != nil {
			log.Fatalln(err)
		}
		user.Submit(problemId, string(sourceCode))

	} else {
		color.Red("Not loggen in")
	}
}
