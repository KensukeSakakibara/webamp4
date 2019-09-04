/*
config.go
@import github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/config
@author Kensuke Sakakibara
@since 2019.08.28
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note 設定情報をconfig/*.jsonより読み込んで保持しておく処理です。シングルトンで動作します。
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Db  *DbConfig
	App *AppConfig
}

var config *Config = new(Config)

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

type AppConfig struct {
	ApplicationPort string `json:"application_port"`
	SessionName     string `json:"session_name"`
	RedisHost       string `json:"redis_host"`
	RedisPassword   string `json:"redis_password"`
	AdminEmail      string `json:"admin_email"`
	AdminPassword   string `json:"admin_password"`
	AdminName       string `json:"admin_name"`
}

// コンストラクタ
func NewConfig() *Config {
	if config.Db != nil && config.App != nil {
		return config
	}

	// DBコンフィグの読み込み
	dbRaw, err := ioutil.ReadFile("./config/db.json")
	if err != nil {
		log.Fatal(err)
	}
	var dbConfig DbConfig
	json.Unmarshal(dbRaw, &dbConfig)

	// APPコンフィグの読み込み
	appRaw, err := ioutil.ReadFile("./config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	var appConfig AppConfig
	json.Unmarshal(appRaw, &appConfig)

	// コンフィグ
	config.Db = &dbConfig
	config.App = &appConfig
	return config
}
