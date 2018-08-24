package main

import (
	"log"
	"os"

	"github.com/titpetric/go-web-crontab/crontab"
)

func main() {
	config := flags("crontab", crontab.Flags)

	// log to stdout not stderr
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go NewMonitor(config.monitorInterval)

	if err := crontab.Init(); err != nil {
		log.Fatalf("Error initializing crm: %+v", err)
	}
	if err := crontab.Start(); err != nil {
		log.Fatalf("Error starting/running crm: %+v", err)
	}
}
