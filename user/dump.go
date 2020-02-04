package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/anatolyefimov/cf-cli/utils"
)

func (user *User) dump() {
	user.Password = string(utils.Encrypt([]byte(user.Password), user.Handle+"666"))
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

// func (user *User) fetch() {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	file, err := os.Open(homeDir + "/" + utils.DumpName)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer file.Close()

// 	d, err := ioutil.ReadAll(file)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = json.Unmarshal(d, user)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
