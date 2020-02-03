package main

import (
	"os"

	"github.com/anatolyefimov/cf-cli/cmd"
)

func main() {
	args := os.Args[1:]
	if args[0] == "login" {
		cmd.Login()
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
