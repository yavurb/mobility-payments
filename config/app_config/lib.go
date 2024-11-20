package appconfig

import (
	"context"
	"fmt"
)

var Environments = []string{"development", "production"}

func LoadConfig() *Config {
	var (
		appConfig *Config
		err       error
	)

	for _, env := range Environments {
		fmt.Printf("Trying to load config from %s ⚙️\n", env)
		confPath := fmt.Sprintf("config/%s-config.pkl", env)

		appConfig, err = LoadFromPath(context.Background(), confPath)
		if err != nil {
			fmt.Printf("Failed to load config from %s\n", confPath)
			fmt.Println(err)
		} else if appConfig != nil {
			fmt.Printf("Loaded config from %s ✔︎\n", env)
			break
		}
	}

	if appConfig == nil {
		panic(err)
	}

	return appConfig
}
