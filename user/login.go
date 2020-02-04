package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"regexp"

	"github.com/anatolyefimov/cf-cli/utils"
	"github.com/fatih/color"
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

//IsLoggedIn check whether the user is logged in
func (user *User) IsLoggedIn(html []byte) bool {
	re := regexp.MustCompile(fmt.Sprintf(`var handle = "%s"`, user.Handle))
	return re.Match(html)
}

// Login handler
func (user *User) Login() {

	resp, err := user.Client.Get(utils.Host + "/enter")
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
		log.Fatalln("CSRF not found")
	}

	ftaa := getFtaa()
	bfaa := getBfaa()

	resp, err = user.Client.PostForm(utils.Host+"/enter", url.Values{
		"csrf_token":    {csrf},
		"action":        {"enter"},
		"ftaa":          {ftaa},
		"bfaa":          {bfaa},
		"handleOrEmail": {user.Handle},
		"password":      {user.Password},
		"_tta":          {"434"},
		"remember":      {"on"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)

	if user.IsLoggedIn(html) {
		user.Bfaa = bfaa
		user.Ftaa = ftaa
		color.Green("You are logged in as %s\n", user.Handle)
		user.dump()
	} else {
		re := regexp.MustCompile(`error for__password`)
		if re.Match(html) {
			color.Red(`Invalid handle/email or password`)
		} else {
			fmt.Println(`Not logged in`)
		}
	}
}
