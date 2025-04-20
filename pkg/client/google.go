package client

import (
	"resty.dev/v3"
)

type GoogleUser struct {
	Email   string `json:"email" copier:"Username"`
	Name    string `json:"name" copier:"Name"`
	Picture string `json:"picture" copier:"Avatar"`
}

func GetGoogleUserInfo(token string) (GoogleUser, error) {
	c := resty.New()
	defer c.Close()

	result := GoogleUser{}
	_, err := c.R().
		SetQueryParam("access_token", token).
		SetResult(&result).
		Get("https://www.googleapis.com/oauth2/v1/userinfo")
	return result, err
}
