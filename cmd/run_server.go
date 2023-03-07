package main

import (
	"GolangCurrencyMS/config"
	act_handler "GolangCurrencyMS/internal/handler/actions"
	"GolangCurrencyMS/internal/server"
	"fmt"
)

func main() {
	appConfiguration, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Oh no!")
		panic(err)
	}

	server := server.NewServer(
		appConfiguration,
		act_handler.NewGetCurrencyController(),
	)

	if err = server.Run(); err != nil {
		panic(err)
	}
}
