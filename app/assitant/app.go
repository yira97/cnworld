package assitant

import (
	"github.com/yira97/cnworld/app/assitant/options"
	"github.com/yira97/cnworld/app/assitant/routes"
	"github.com/yira97/cnworld/app/assitant/storage"
)

func Run(cfg string) {
	options.LoadAppOption(cfg)

	storage.Setup()

	routes.Setup()

	routes.Start()
}
