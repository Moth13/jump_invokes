package main

import (
	"flag"

	app "invokes/internal/api"
	utils "invokes/internal/utils"

	log "github.com/sirupsen/logrus"
)

// @title Invokes Rest API
// @version 0.0.0
// @BasePath /docs/
// @description Invokes Rest API provides user datas, and POST endpoint to handle invoices and transactions
// @termsOfService 'http://swagger.io/terms/'
// @license.name MothyV
func main() {

	var confPath = flag.String("conf", "config.yml", "conf file path")

	flag.Parse()

	config, err := utils.LoadConfiguration(confPath)
	if err != nil {
		log.Fatalf("Failed to load configuration from %s, error:%s", *confPath, err.Error())
		return
	}

	if err = utils.LoadLogger(&config); err != nil {
		log.Fatalf("Can't configure logger %s", err.Error())
		return
	}

	a := app.App{}

	err = a.Initialize(&config)
	if err != nil {
		log.Fatal(err)
		return
	}

	a.Run()
}
