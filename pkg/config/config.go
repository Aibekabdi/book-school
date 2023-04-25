package config

import (
	"encoding/json"
	"os"
)

type (
	Conf struct {
		Api      Api      `json:"api"`
		Database Database `json:"database"`
	}
	Api struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	}
)

func NewConfig(path string) (*Conf, error) {
	var newConfig Conf
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&newConfig); err != nil {
		return nil, err
	}
	return &newConfig, nil
}
