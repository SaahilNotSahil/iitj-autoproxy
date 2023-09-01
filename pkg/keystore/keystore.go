package keystore

import (
	"github.com/99designs/keyring"
)

var kr keyring.Keyring

func init() {
	kr, _ = keyring.Open(keyring.Config{
		ServiceName: "iitj-autoproxy",
		KeyCtlScope: "user",
		FileDir:    "~/.config/iitj-autoproxy",
	})
}

func Get(key string) (string, error) {
	item, err := kr.Get(key)
	if err != nil {
		return "", err
	}

	return string(item.Data), nil
}

func Set(key string, value string) error {
	return kr.Set(keyring.Item{
		Key:  key,
		Data: []byte(value),
	})
}

func Remove(key string) error {
	return kr.Remove(key)
}

func Keys() ([]string, error) {
	return kr.Keys()
}

func Reset() error {
	keys, err := kr.Keys()
	if err != nil {
		return err
	}

	for _, key := range keys {
		err := kr.Remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}
