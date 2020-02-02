package user

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"regexp"
)

func getCsrf(body []byte) (string, error) {
	reg := regexp.MustCompile(`csrf='(.+?)'`)
	tmp := reg.FindSubmatch(body)
	if len(tmp) < 2 {
		return "", errors.New("Cannot find csrf")
	}
	return string(tmp[1]), nil
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
	return "b4020441f030d052e817e828e9aef3ff"
}

// Login handler
func (user *User) Login() {

	resp, err := user.client.Get("https://codeforces.com" + "/enter")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Body reading error")
	}

	csrf, err := getCsrf(body)
	if err != nil {
		log.Fatalln("CSRF not found")
	}

	ftaa := getFtaa()
	bfaa := getBfaa()

	resp, err = user.client.PostForm("https://codeforces.com"+"/enter", url.Values{
		"csrf_token":    {csrf},
		"action":        {"enter"},
		"ftaa":          {ftaa},
		"bfaa":          {bfaa},
		"handleOrEmail": {user.handle},
		"password":      {user.password},
		"_tta":          {"434"},
		"remember":      {"on"},
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
}
