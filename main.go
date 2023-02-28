package main

import (
	"crpalert/clients"
	"crpalert/handlers"
)

func main() {
	bot := clients.Init()
	handlers.Init(bot)
}
