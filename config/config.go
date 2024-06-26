package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
	IP              string
	Port            string
}

var Cfg *tomlConfig

func init() {
	//初始化方法
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "rm-go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("E:\\code\\go\\rm-go-blog\\config\\config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
