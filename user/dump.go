package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/anatolyefimov/cf-cli/utils"
)

//Dump user info to disk
func (user *User) Dump() {
	user.Password = utils.Encrypt(user.Password, user.Handle+"666")
	d, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(homeDir+"/"+utils.DumpName, d, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

//Fetch user info from disk
func (user *User) Fetch() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Open(homeDir + "/" + utils.DumpName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	d, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(d, user)
	user.Password = utils.Decrypt(user.Password, user.Handle+"666")
	if err != nil {
		log.Fatalln(err)
	}
}
