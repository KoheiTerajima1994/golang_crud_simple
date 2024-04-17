package config

import (
	"log"

	"github.com/go-ini/ini"
)

// Configリストの構造体
type ConfigList struct {
	Port    string
	LogFile string
}

// グローバルなconfigリストを変数に設定
var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	// config.iniファイルを読み込む
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	// 変数Configにconfig.iniの値を入れていく
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").MustString("8080"),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}
