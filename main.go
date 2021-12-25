package main

import (
	"github.com/yira97/cnworld/app/assitant"
	"github.com/yira97/cnworld/app/blog"
)

func main() {

	go assitant.Run("./app/assitant/config.json")
	go blog.Run("./app/blog/config.json")

	select {}
}
