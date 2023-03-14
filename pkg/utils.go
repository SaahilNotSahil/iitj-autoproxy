package pkg

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	if strings.Contains(bodyString, viper.GetString("login_base_url")) {
		strs := strings.SplitAfter(bodyString, "window.location=\""+viper.GetString("login_base_url")+"fgtauth?")
		prefix, found := strings.CutSuffix(strs[1], "\";</script></body></html>")

		if found {
			token = prefix
		} else {
			token = ""
		}
	}

	return token, nil
}
