package main

import (
	"github.com/mutual-fund-tracker/src/container"
	"github.com/mutual-fund-tracker/src/server"
)

func main() {
	readyContainer, err := container.BuildContainer()
	if err != nil {
		panic(err)
	}

	if err = readyContainer.Invoke(server.StartServer); err != nil {
		panic(err)
	}
	return
}
