package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/sevren/go-vasttrafik/api"
	"github.com/sevren/go-vasttrafik/auth"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Debug bool `env:"DEBUG"`
	VT    auth.Config
	API   api.Config
}

func checkConfig(cfg *config) error {
	if cfg.VT.Key == "" {
		return fmt.Errorf("VT_KEY can't be empty")
	}

	if cfg.VT.Secret == "" {
		return fmt.Errorf("VT_SECRET can't be empty")
	}

	return nil
}

func main() {
	log.Info("Initalizing go-vasttrafik...")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Could not parse config %s", err)
	}

	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	if err := checkConfig(&cfg); err != nil {
		log.Fatal(err)
	}

	a, err := auth.GetAccessToken(cfg.VT.Key, cfg.VT.Secret)
	if err != nil {
		log.Fatal(err)
	}
	cl := &api.Client{
		Base: cfg.API.BaseURL,
		Auth: *a,
	}

	log.Info(cl.Auth.Token)

	resp, err := cl.GetLocationIds()
	if err != nil {
		log.Error(err)
	}

	log.Infof("Response back from getLocationIds %+v", resp.StopLocations)

}
