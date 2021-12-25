package blog

import (
	"github.com/yira97/cnworld/app/blog/config"
	"github.com/yira97/cnworld/app/blog/route"
	"github.com/yira97/cnworld/app/blog/storage"
)

func Run(path string) {
	config.Setup(path)
	storage.Setup()
	route.Setup()
	route.Start()
}
