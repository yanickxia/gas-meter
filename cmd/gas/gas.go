package main

import "github.com/yanickxia/gas-meter/pkg/server"

func main() {
	s := server.NewServer(":9876")
	if err := s.Run(); err != nil {
		return
	}
}
