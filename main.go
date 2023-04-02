package main

import (
	"fmt"

	"github.com/madfelps/gopportunities/config"
	"github.com/madfelps/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	fmt.Println("ao")

	logger = *config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
