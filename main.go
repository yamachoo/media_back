package main

import (
	"github.com/yamachoo/media_back/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}
