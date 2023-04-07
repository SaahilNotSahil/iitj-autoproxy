package pkg

import (
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"strings"
	"syscall"

	"github.com/spf13/viper"
)

func Login(url string, username string, password string) error {
	var magic string

	token, err := GetToken(url)
	if err != nil {
		return err
	}

	loginPageURL := viper.GetString("login_base_url") + "fgtauth?" + token

	res, err := http.Get(loginPageURL)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	bodyString := string(body)

	magicStr := "name=\"magic\" value=\""
	if strings.Contains(bodyString, magicStr) {
		strs := strings.SplitAfter(bodyString, magicStr)
		strs = strings.Split(strs[1], "\"")
		magic = strs[0]
	}

	referer := viper.GetString("login_base_url") + "login?" + token

	data := u.Values{}
	data.Add("4Tredir", referer)
	data.Add("magic", magic)
	data.Add("username", username)
	data.Add("password", password)

	client := http.Client{}
	defer client.CloseIdleConnections()

	req, err := http.NewRequest("POST", viper.GetString("login_base_url"), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", referer)

	req.Close = true

	res, err = client.Do(req)
	if err != nil {
		log.Println("Error logging in")
		return err
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	bodyString = string(body)

	if strings.Contains(bodyString, "keepalive?") {
		strs := strings.SplitAfter(bodyString, "keepalive?")
		prefix, found := strings.CutSuffix(strs[1], "\";</script></body></html>")

		if found {
			token = prefix
		} else {
			token = ""
		}
	}

	viper.Set("token", token)

	return viper.WriteConfig()
}

func Logout(token string) error {
	url := viper.GetString("login_base_url") + "logout?" + token

	_, err := http.Get(url)
	if err != nil {
		return err
	}

	pid := viper.GetInt("pid")

	return syscall.Kill(pid, syscall.SIGKILL)
}

func GetToken(url string) (string, error) {
	var token string

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	bodyString := string(body)

	if strings.Contains(bodyString, "fgtauth?") {
		strs := strings.SplitAfter(bodyString, "fgtauth?")
		prefix, found := strings.CutSuffix(strs[1], "\";</script></body></html>")

		if found {
			token = prefix
		} else {
			token = ""
		}
	}

	return token, nil
}

func GetCurrentKeepaliveToken() (string, error) {
	var token string
	url := viper.GetString("login_base_url") + "keepalive?" + viper.GetString("token")

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(body)

	if strings.Contains(bodyString, "<p><a href=\"https://gateway.iitj.ac.in:1003/logout?") {
		strs := strings.SplitAfter(bodyString, "<p><a href=\"https://gateway.iitj.ac.in:1003/logout?")
		strs = strings.Split(strs[1], "\"")
		token = strs[0]
	} 

	return token, nil
}
