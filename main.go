package main

import (
	"binance/listener/cmd/listener"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	listener.RunListener()
}
