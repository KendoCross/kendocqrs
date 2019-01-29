package main

import (
	"github.com/KendoCross/kendocqrs/presentation"
)

func main() {
	router := presentation.InitRouter()
	router.Run(":8123")

}
