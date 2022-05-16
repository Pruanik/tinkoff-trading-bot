package main

import (
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		panic("Can't read configuration file: " + err.Error())
	}

	fmt.Println(config.DatabaseConfig.User)
}
