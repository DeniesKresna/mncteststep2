package main

import (
	"fmt"

	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/mncteststep2/config"
	"github.com/DeniesKresna/mncteststep2/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	cfg := config.CreateNewConfig()

	err = cfg.SetConfigApplication()
	if err != nil {
		utlog.Errorf(err.Error())
		return
	}

	err = cfg.SetConfigDatabase()
	if err != nil {
		utlog.Errorf(err.Error())
		return
	}

	err = service.Start(cfg)
	if err != nil {
		utlog.Errorf(err.Error())
		return
	}
}
