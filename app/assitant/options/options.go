package options

import (
	"encoding/json"
	"os"

	"github.com/yira97/cnworld/lib/base/defawt"
)

type Session struct {
	Secret string `json:"secret"`
}

var app *Application

type Application struct {
	Port    string                `json:"port"`
	MongoDB defawt.MongoDB_Option `json:"mongodb"`
	Session Session               `json:"session"`
}

func LoadAppOption(configFilePathString string) {
	app = new(Application)

	configFileString, err := os.ReadFile(configFilePathString)
	if err != nil {
		panic("读取配置文件失败")
	}

	if err := json.Unmarshal(configFileString, app); err != nil {
		panic(err)
	}
}

func GetPort() string {
	return app.Port
}

func GetSession() Session {
	return app.Session
}

func GetMongoDB() defawt.MongoDB_Option {
	return app.MongoDB
}
