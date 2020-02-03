package cmd

import (
	"fmt"

	"github.com/anatolyefimov/cf-cli/user"
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
