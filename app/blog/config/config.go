package config

import (
	"encoding/json"
	"os"

	"github.com/yira97/cnworld/lib/base/defawt"
)

type RouteConfig struct {
	Port string `json:"port"`
}

type AppConfig struct {
	Route   RouteConfig           `json:"route"`
	MongoDB defawt.MongoDB_Option `json:"mongodb"`
}

var c = &AppConfig{}

func Setup(configFilePathString string) {
	configFileString, err := os.ReadFile(configFilePathString)
	if err != nil {
		panic("读取配置文件失败")
	}

	if err := json.Unmarshal(configFileString, c); err != nil {
		panic(err)
	}
}

func GetMongoDB() defawt.MongoDB_Option {
	return c.MongoDB
}

func GetPort() string {
	return c.Route.Port
}
