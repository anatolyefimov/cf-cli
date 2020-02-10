package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"time"

	"github.com/anatolyefimov/cf-cli/utils"
	"github.com/fatih/color"
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
	resp, err = user.Client.PostForm(utils.Host+"problemset/submit?csrf_token="+csrf, url.Values{
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

	defer resp.Body.Close()
	doc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`error for__source`)
	if re.Match(doc) {
		color.Red("You have submitted exactly the same code before")
	} else {

		reVerdict := regexp.MustCompile(`"verdict":"([A-Z_]*)"`)
		reNumberTest := regexp.MustCompile(`"passedTestCount":(\d*)`)

		for {

			time.Sleep(time.Second)
			resp, err := user.Client.Get("http://codeforces.com/api/user.status?handle=" + user.Handle + "&from=1&count=10")
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()
			bodyByte, _ := ioutil.ReadAll(resp.Body)

			body := string(bodyByte)
			resv := reVerdict.FindStringSubmatch(body)
			resnt := reNumberTest.FindStringSubmatch(body)

			g := color.New(color.FgGreen)
			r := color.New(color.FgHiRed)

			if resv == nil || len(resv) <= 1 {
				color.Red("Error: Verdict not found")
			} else {
				nt := resnt[1]
				verdict := resv[1]
				if verdict != "TESTING" {
					if verdict == "OK" {
						g.Printf(utils.ReplaceOutput("OK"))
					}
					if verdict == "WRONG_ANSWER" {
						r.Printf(utils.ReplaceOutput("Wront answer on " + nt))
					}
					if verdict == "TIME_LIMIT_EXCEEDED" {
						r.Printf(utils.ReplaceOutput("Time limit exceeded on" + nt))
					}
					if verdict == "COMPILATION_ERROR" {
						r.Printf(utils.ReplaceOutput("Compilation error"))
					}

					break
				}
				if verdict == "TESTING" {
					t := "Testing"

					fmt.Printf("%s", utils.ReplaceOutput(t))
				}
			}

		}
	}

}
