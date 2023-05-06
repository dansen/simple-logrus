package main

import (
	log "github.com/dansen/simple-logrus/log"
)

func main() {
	println("Hello, world!")
	log.SetOutput(true, "logfiles/test.log")
	log.Infof("Hello, world!")
}
