package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yanickxia/gas-meter/pkg/server"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	s := server.NewServer(":9876")
	if err := s.Run(); err != nil {
		return
	}
}
