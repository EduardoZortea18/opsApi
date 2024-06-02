package main

import (
	"github.com/EduardoZortea18/opsApi/config"
	"github.com/EduardoZortea18/opsApi/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initial config error: %v", err)
		return
	}

	router.Initialize()
}
