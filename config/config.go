package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	AppName string `json:"app_name"`
	Port    string `json:"port"`
	Mode    string `json:"mode"`
}

func InitAppConfig() *AppConfig {
	file, ee := os.Open("/Users/kzcai/goProject/appMarket/config.json")

	if ee != nil {
		fmt.Println(ee)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}

	return &conf
}
