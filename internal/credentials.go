package internal

import (
	"github.com/zalando/go-keyring"
)

const service = "ghws"
const user = "GHWS_GITHUB_PAT"

func GetAuthToken() (string, bool) {
	token, err := keyring.Get(service, user)

	if err != nil || token == "" {
		return "", false
	}

	return token, true
}

func SaveAuthToken(token string) error {
	err := keyring.Set(service, user, token)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAuthToken() error {
	err := keyring.Delete(service, user)
	if err != nil {
		return err
	}
	return nil
}
