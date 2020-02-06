package user

import (
	"io/ioutil"
	"log"
	"net/url"

	"github.com/anatolyefimov/cf-cli/utils"
)

//Submit solution to archive
func (user *User) Submit(problemID string, source string) {
	resp, err := user.Client.Get(utils.Host + "problemset/submit")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	csrf, err := getCsrf(body)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = user.Client.PostForm(utils.Host+"problemset/submit?csrf_token="+csrf, url.Values{
		"csrf_token":           {csrf},
		"ftaa":                 {user.Ftaa},
		"bfaa":                 {user.Bfaa},
		"action":               {"submitSolutionFormSubmitted"},
		"submittedProblemCode": {problemID},
		"programTypeId":        {"50"},
		"source":               {source},
		"tabSize":              {"4"},
		"sourceFile":           {},
		"_tta":                 {"434"},
	})
	if err != nil {
		log.Fatalln(err)
	}
}
