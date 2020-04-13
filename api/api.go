package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sevren/go-vasttrafik/auth"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Base string
	Auth auth.Auth
}

// Config holds the api configurations
type Config struct {
	BaseURL string `env:"BASE_URL" envDefault:"https://api.vasttrafik.se/bin/rest.exe/v2"`
}

func (c *Client) GetLocationIds() (*getLocationResponse, error) {

	cl := &http.Client{
		Timeout: time.Second * 10,
	}

	c.Auth = *c.Auth.RefreshToken()

	bearerHeaderVal := fmt.Sprintf("Bearer %s", c.Auth.Token)

	message := getLocationRequest{
		Input:  "Svarte Mosse",
		Format: "json",
	}

	// b, err := json.Marshal(message)
	// if err != nil {
	// 	return nil, err
	// }

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s?input=%s&format=json", c.Base, "location.name", url.QueryEscape(message.Input)), nil)
	req.Header.Add("Authorization", bearerHeaderVal)

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

	log.Info(string(bs))

	l := getLocationResponse{}
	// err = json.Unmarshal(bs, &l)
	// if err != nil {
	// 	return nil, err
	// }

	return &l, nil

}
