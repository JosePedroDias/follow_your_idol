package config

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
)

type Config struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessTokenKey    string `json:"access_token_key"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func FromJSONFile(filePath string) (*Config, error) {
	var config Config

	cfgS, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &config, err
	}

	if err := json.Unmarshal(cfgS, &config); err != nil {
		return &config, err
	}

	return &config, nil
}
