package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"My-Docker/container"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
