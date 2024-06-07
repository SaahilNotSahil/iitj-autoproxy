package keystore

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}

	configName := ".autoproxy.config"

	viper.AddConfigPath(home)

	viper.SetConfigType("json")
	viper.SetConfigName(configName)

	err = viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}
}

func Get(key string) (string, error) {
	item := viper.GetString(key)
	if item != "" {
		return "", errors.New("Please set the value for " + key)
	}

	return item, nil
}

func Set(key string, value string) {
	viper.Set(key, value)
}

func Remove(key string) {
	viper.Set(key, "")
}

func Keys() []string {
	return viper.AllKeys()
}
