package main

import (
	"webmalc/mb-redirector/common/config"
	"webmalc/mb-redirector/common/logger"
)

func main() {
	config.Setup()
	c := NewConfig()
	server := NewServer(c, NewAPI(c, logger.NewLogger()))
	server.Run()
}
