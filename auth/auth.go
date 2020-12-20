package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Key    string `env:"VT_KEY"`
	Secret string `env:"VT_SECRET"`
}

type token struct {
	Scope       string `json:"scope"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` // in seconds
	Accesstoken string `json:"access_token"`
}

// Auth the authentication structure containing the token required to make calls to Vasttrafiks API
type Auth struct {
	creds string
	//The access token used to call the Vasttrafik api
	Token string
	// Access token validity in seconds epoch
	ExpiresAt int64
}

func credEncode(key string, secret string) string {
	credentials := fmt.Sprintf("%s:%s", key, secret)
	credentialsEncoded := base64.URLEncoding.EncodeToString([]byte(credentials))
	log.Debugf("Credentials encoded %s", credentialsEncoded)
	return credentialsEncoded
}

func calculateExpiry(s int) int64 {
	return time.Now().Add(time.Duration(s) * time.Second).Unix()
}

//refreshes the Bearer token if invalid
func (a *Auth) RefreshToken() *Auth {
	if time.Now().Unix() >= a.ExpiresAt {

		t, err := generateAccessToken(a.creds)
		if err != nil {
			log.Fatal(err)
		}
		a.Token = t.Accesstoken
		a.ExpiresAt = calculateExpiry(t.ExpiresIn)

	}
	return a
}

func generateAccessToken(creds string) (*token, error) {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	cl := &http.Client{
		Timeout: time.Second * 10,
	}

	basicHeaderVal := fmt.Sprintf("Basic %s", creds)

	req, err := http.NewRequest("POST", "https://api.vasttrafik.se:443/token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Authorization", basicHeaderVal)

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	t := token{}
	err = json.Unmarshal(bs, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil

}

func GetAccessToken(key string, secret string) (*Auth, error) {

	creds := credEncode(key, secret)

	t, err := generateAccessToken(creds)
	et := calculateExpiry(t.ExpiresIn)
	if err != nil {
		return nil, err
	}

	return &Auth{
		creds:     creds,
		Token:     t.Accesstoken,
		ExpiresAt: et,
	}, nil
}
