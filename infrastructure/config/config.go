package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Adapter struct {
	Dialect  string `json:"dialect"`
	Username string `json:"username"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Database string `json:"database"`
}

type DbConfig struct {
	Master Adapter   `json:"master"`
	Slaves []Adapter `json:"slaves"`
}

func GetConfig() *DbConfig {
	// コンフィグの読み込み
	raw, err := ioutil.ReadFile("./config/db.json")
	if err != nil {
		log.Fatal(err)
	}

	var dbconfig DbConfig
	json.Unmarshal(raw, &dbconfig)

	return &dbconfig;
}
