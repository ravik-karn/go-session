package main

import (
	"flag"
	"fmt"
	"go-database/factory"
	"log"

	"go-database/config"
)

func main() {
	flag.Usage = usage
	confFile := flag.String("config", "", "Path to JSON configuration file if any")

	var conf *config.Config
	var errs []error
	if confFile == nil || *confFile == "" {
		conf, errs = config.LoadConfigFromEnv()
	} else {
		conf, errs = config.LoadConfigFromFile(*confFile)
	}
	if len(errs) > 0 {
		for _, err := range errs {
			log.Print(err)
		}
		log.Fatalln("Invalid configuration. Aborting!")
	}

	f := factory.New(conf)
	f.DB().Query("")
}

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
}

const (
	usagePrefix = `
Demo service for database connection

Usage:
	service-name -config [FILE] # reads configuration from [FILE]
	service-name                # reads configuration from environment variables

Options:
`
)
