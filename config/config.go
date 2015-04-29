package config

import (
	"encoding/json"
	"io/ioutil"
)

type TwitterConfig struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessTokenKey    string `json:"access_token_key"`
	AccessTokenSecret string `json:"access_token_secret"`
}

type PostgresqlConfig struct {
	DbName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	SslMode  string `json:"sslmode"`
}

func TwitterFromFile(filePath string) (*TwitterConfig, error) {
	var config TwitterConfig

	cfgS, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &config, err
	}

	if err := json.Unmarshal(cfgS, &config); err != nil {
		return &config, err
	}

	return &config, nil
}

func PostgresqlFromFile(filePath string) (*PostgresqlConfig, error) {
	var config PostgresqlConfig

	cfgS, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &config, err
	}

	if err := json.Unmarshal(cfgS, &config); err != nil {
		return &config, err
	}

	return &config, nil
}
