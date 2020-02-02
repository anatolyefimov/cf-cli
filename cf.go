package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

var ErrorNotLogged = "Not logged in"

func checkLogin(username string, body []byte) error {
	match, err := regexp.Match(fmt.Sprintf(`handle = "%v"`, username), body)
	if err != nil || !match {
		return errors.New(ErrorNotLogged)
	}
	return nil
}

func getFtaa() string {
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyz"
	ftaa := make([]byte, 18)
	for i := 0; i < 18; i++ {
		ftaa[i] = alphabet[rand.Intn(36)]
	}

	return string(ftaa)
}

func getBfaa() string {
	return "f1b3f18c715565b589b7823cda7448ce"
}

func findCsrf(body []byte) (string, error) {
	reg := regexp.MustCompile(`csrf='(.+?)'`)
	tmp := reg.FindSubmatch(body)
	if len(tmp) < 2 {
		return "", errors.New("Cannot find csrf")
	}
	return string(tmp[1]), nil
}

func main() {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{Jar: jar}

	resp, err := client.Get("https://codeforces.com" + "/enter")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	csrf, err := findCsrf(body)
	if err != nil {
		return
	}

	ftaa := getFtaa()
	bfaa := getBfaa()

	resp, err = client.PostForm("https://codeforces.com"+"/enter", url.Values{
		"csrf_token":    {csrf},
		"action":        {"enter"},
		"ftaa":          {ftaa},
		"bfaa":          {bfaa},
		"handleOrEmail": {"anatolyefimov"},
		"password":      {"575410820u"},
		"_tta":          {"434"},
		"remember":      {"on"},
	})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	// fmt.Print(string(body))
	if err != nil {
		return
	}
	err = checkLogin("anatolyefimov", body)
	if err != nil {
		return
	}

	csrf, err = findCsrf(body)
	if err != nil {
		return
	}
	resp, err = client.PostForm("https://codeforces.com/problemset/submit?csrf_token="+csrf, url.Values{
		"csrf_token":           {csrf},
		"ftaa":                 {ftaa},
		"bfaa":                 {bfaa},
		"action":               {"submitSolutionFormSubmitted"},
		"submittedProblemCode": {"33A"},
		"programTypeId":        {"50"},
		"source":               {"101010010101010100101010101032423451010"},
		"tabSize":              {"4"},
		"sourceFile":           {},
		"_tta":                 {"434"},
	})
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
