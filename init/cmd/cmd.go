package cmd

import (
	"fmt"
	"golang_backend_study/config"
	"golang_backend_study/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}
	fmt.Println(c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return &c
}
