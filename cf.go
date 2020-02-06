package main

import (
	"os"

	"github.com/anatolyefimov/cf-cli/cmd"
)

func main() {

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// file, err := os.Open(homeDir + "/" + utils.DumpName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()
	// user := user.New("", "")

	// js, err := ioutil.ReadAll(file)
	// _ = json.Unmarshal(js, user)
	// // fmt.Println(utils.Decrypt([]byte(user.Password), user.Handle+"666"))
	// fmt.Println(user.Handle + "666")
	args := os.Args[1:]
	if args[0] == "login" {
		cmd.Login()
	} else if args[0] == "submit" {
		cmd.Submit(args[1], args[2])
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
