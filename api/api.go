package api

import (
	"encoding/json"
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
	Conn *http.Client
}

//New creates a new default client
func New(apiBase string, authenticator *auth.Auth) *Client {
	cl := &Client{
		Base: apiBase,
		Auth: *authenticator,
		Conn: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	return cl
}

// Config holds the api configurations
type Config struct {
	BaseURL string `env:"BASE_URL" envDefault:"https://api.vasttrafik.se/bin/rest.exe/v2"`
}

func (c *Client) GetLocationIds() (*getLocationResponse, error) {

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

	resp, err := c.Conn.Do(req)
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

func (c *Client) GetTrip() (*getTripResponse, error) {
	//IDs 9021014006462000 == Svarte Mosse, goteborg
	//IDs 9021014001760000 == Brunnsparken

	c.Auth = *c.Auth.RefreshToken()

	bearerHeaderVal := fmt.Sprintf("Bearer %s", c.Auth.Token)

	message := getTripRequest{
		OriginID: "9021014006462000",
		DestID:   "9021014001760000",
		Format:   "json",
	}

	// b, err := json.Marshal(message)
	// if err != nil {
	// 	return nil, err
	// }

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s?originId=%s&destId=%s&format=json", c.Base, "trip", url.QueryEscape(message.OriginID), url.QueryEscape(message.DestID)), nil)
	req.Header.Add("Authorization", bearerHeaderVal)

	resp, err := c.Conn.Do(req)
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

	t := &getTripResponse{}
	err = json.Unmarshal(bs, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
