package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
	"strings"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(viper.ReadInConfig())

		username := viper.GetString("username")
		password := viper.GetString("password")

		if username == "" || password == "" {
			cobra.CheckErr("Please configure the application using the config command")
		}

		pkg.Scheduler.Clear()

		_, err := pkg.Scheduler.Every(10).Seconds().Do(func() {
			file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			cobra.CheckErr(err)
			defer file.Close()

			logger := log.New(file, "", log.LstdFlags)

			Login(viper.GetString("base_url"), username, password)

			logger.Println("Login was run")
		})
		cobra.CheckErr(err)

		log.Println("Starting Scheduler")
		go pkg.Scheduler.StartBlocking()

		select {}
	},
}

func Login(url string, username string, password string) error {
	var magic string

	token, err := pkg.GetToken(url)
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

	magicStr := "<input type=\"hidden\" name=\"magic\" value=\""
	if strings.Contains(bodyString, magicStr) {
		strs := strings.SplitAfter(bodyString, magicStr)
		strs = strings.Split(strs[1], "\"")
		magic = strs[0]
	}

	data := u.Values{}
	data.Add("4Tredir", viper.GetString("login_base_url")+"login?"+token)
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
	req.Header.Add("Referer", viper.GetString("login_base_url")+"login?"+token)

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

	if strings.Contains(bodyString, viper.GetString("login_base_url")) {
		strs := strings.SplitAfter(bodyString, "window.location=\""+viper.GetString("login_base_url")+"keepalive?")
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
